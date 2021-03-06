package deep

import (
	"context"
	"database/sql"

	"github.com/src-d/ghsync/models"

	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-kallax.v1"
	log "gopkg.in/src-d/go-log.v1"
	"gopkg.in/src-d/go-queue.v1"
)

type IssueSyncer struct {
	s *models.IssueStore
	c *github.Client
}

func NewIssueSyncer(db *sql.DB, c *github.Client) *IssueSyncer {
	return &IssueSyncer{
		s: models.NewIssueStore(db),
		c: c,
	}
}

func (s *IssueSyncer) QueueRepository(q queue.Queue, owner, repo string) error {
	opts := &github.IssueListByRepoOptions{}
	opts.ListOptions.PerPage = listOptionsPerPage
	opts.State = "all"

	logger := log.New(log.Fields{"type": IssueSyncTask, "owner": owner, "repo": repo})
	logger.Infof("starting to publish queue jobs")

	for {
		issues, r, err := s.c.Issues.ListByRepo(context.TODO(), owner, repo, opts)
		if err != nil {
			return err
		}

		for _, i := range issues {
			if i.PullRequestLinks != nil {
				continue
			}

			j, err := NewIssueSyncJob(owner, repo, i.GetNumber())
			if err != nil {
				return err
			}

			l := logger.With(log.Fields{"issue": i.GetNumber()})
			l.Debugf("queue request")
			if err := q.Publish(j); err != nil {
				l.Errorf(err, "publishing job")
				return nil
			}
		}

		if r.NextPage == 0 {
			break
		}

		opts.Page = r.NextPage
	}

	logger.Infof("finished to publish queue jobs")

	return nil
}

func (s *IssueSyncer) Sync(owner string, repo string, number int) error {
	issue, _, err := s.c.Issues.Get(context.TODO(), owner, repo, number)
	if err != nil {
		return err
	}

	record, err := s.s.FindOne(models.NewIssueQuery().
		Where(kallax.And(
			kallax.Eq(models.Schema.Issue.RepositoryOwner, owner),
			kallax.Eq(models.Schema.Issue.RepositoryName, repo),
			kallax.Eq(models.Schema.Issue.Number, number),
		)),
	)

	if record == nil {
		record = models.NewIssue()
		record.Issue = *issue

		return s.s.Insert(record)
	}

	record.Issue = *issue
	_, err = s.s.Update(record)
	return err

}
