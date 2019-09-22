package repository

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/google/go-github/github"
	"go.5kpbs.io/operator/pkg/operator"
)

type GithubRepository struct {
	Owner        string
	Repository   string
	GithubClient *github.Client
}

func (r *GithubRepository) Tags() ([]string, error) {
	tagList, _, err := r.GithubClient.Repositories.ListTags(context.Background(), r.Owner, r.Repository, &github.ListOptions{
		PerPage: 100,
	})

	if err != nil {
		return nil, err
	}

	tags := make([]string, 0, len(tagList))

	for _, tag := range tagList {
		tags = append(tags, tag.GetName())
	}
	return tags, nil
}

// Vers return list vers, which is construct by semver (\d+.\d+.\d+), exclude tags from unexpected hotfix
func (r *GithubRepository) Vers() ([]*operator.Semver, error) {

	tags, err := r.Tags()

	if err != nil {
		return nil, err
	}

	vers := make([]*operator.Semver, 0)

	rx, err := regexp.Compile("\\d+.\\d+.\\d+")

	if err != nil {
		return nil, err
	}

	verExists := make(map[string]struct{}, 0)
	ok := struct{}{}

	for _, tag := range tags {
		if !rx.MatchString(tag) { // Skip if version is not suitable with pattern
			continue
		}
		ver := operator.Tag2Semver(tag)
		// Skip if version is exists
		if _, ok := verExists[ver.String()]; ok {
			continue
		}
		vers = append(vers, ver)
		verExists[ver.String()] = ok
	}

	// Sort version from high to low
	sort.Slice(vers, func(i, j int) bool {
		return !vers[i].Less(vers[j])
	})

	return vers, nil
}

func (r *GithubRepository) Changelog(head string, base string) (*operator.Changelog, error) {

	if base == "" {
		vers, err := r.Vers()
		if err != nil {
			return nil, err
		}
		if len(vers) > 0 {
			base = vers[0].String()
		}
	}

	changelog := &operator.Changelog{
		Base: base,
		Head: head,
	}

	diff, _, err := r.GithubClient.Repositories.CompareCommits(context.Background(), r.Owner, r.Repository, base, head)

	if err != nil {
		return nil, err
	}

	for _, commit := range diff.Commits {
		message := commit.GetCommit().GetMessage()
		changelogMessage := fmt.Sprintf("%s %s", commit.GetSHA(), strings.Split(message, "\n")[0])
		switch {
		case strings.HasPrefix(message, "Merge "): // Ignore merge branch, merge pull request commit
			continue
		case strings.Contains(message, "#fixed"):
			changelog.Fixed = append(changelog.Fixed, changelogMessage)
		case strings.Contains(message, "#security"):
			changelog.Security = append(changelog.Security, changelogMessage)
		case strings.Contains(message, "#added"):
			changelog.Added = append(changelog.Added, changelogMessage)
		case strings.Contains(message, "#deprecated"):
			changelog.Deprecated = append(changelog.Deprecated, changelogMessage)
		case strings.Contains(message, "#removed"):
			changelog.Removed = append(changelog.Removed, changelogMessage)
		case strings.Contains(message, "#changed"):
			changelog.Changed = append(changelog.Changed, changelogMessage)
		default:
			changelog.MissingDefination = append(changelog.MissingDefination, changelogMessage)
		}
	}
	return changelog, nil
}

func (r *GithubRepository) DraftRelease(head string, base string) (*operator.Release, error) {
	vers, err := r.Vers()

	if err != nil {
		return nil, err
	}

	if len(vers) > 0 {
		changelog, err := r.Changelog(base, head)
		if err != nil {
			return nil, err
		}
		semver := vers[0]
		semver.Upgrade(changelog)
		return &operator.Release{Semver: semver, Changelog: changelog}, nil
	}
	return &operator.Release{Semver: operator.Tag2Semver("v1.0.0"), Changelog: &operator.Changelog{Head: head}}, nil
}

func (r *GithubRepository) Release(head string, base string) error {
	rl, err := r.DraftRelease(base, head)
	if err != nil {
		return err
	}

	body := rl.Changelog.String()

	name := rl.Semver.Name()
	tag := rl.Semver.String()

	_, _, err = r.GithubClient.Repositories.CreateRelease(context.Background(), r.Owner, r.Repository, &github.RepositoryRelease{
		Name:            &name,
		TagName:         &tag,
		TargetCommitish: &head,
		Body:            &body,
	})

	if err != nil {
		return err
	}

	refObject, _, err := r.GithubClient.Git.GetRef(context.Background(), r.Owner, r.Repository, fmt.Sprintf("refs/heads/%s", head))

	if err != nil {
		return err
	}

	releaseRef := fmt.Sprintf("refs/heads/releases/%s", tag)
	_, _, err = r.GithubClient.Git.CreateRef(context.Background(), r.Owner, r.Repository, &github.Reference{
		Ref:    &releaseRef,
		Object: refObject.GetObject(),
	})

	return err
}
