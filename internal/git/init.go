package git

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/es3649/proj/pkg/log"
	"github.com/google/go-github/github"
)

// Config contains settings for a new github repository
type Config struct {
	IsPrivate     bool
	IgnoreFlavor  string
	IncludeReadme bool
	OmitReadme    bool
}

// InitGit initializes a new remote GitHub repository
func InitGit(cfg Config) error {

	client, err := newGithubClient()
	if err != nil {
		return err
	}

	// create the repository
	repo := &github.Repository{
		// Name: github.String()
		Private:  github.Bool(cfg.IsPrivate),
		AutoInit: github.Bool(!cfg.IncludeReadme),
		// TODO add .gitignore flavor based on language
	}

	_, _, err = client.Repositories.Create(context.Background(), "", repo)
	if err != nil {
		log.Log.Error("Failed to create repository")
		return err
	}

	return nil
}

// newGithubClient creates a new authenticated client for making
// github API calls
func newGithubClient() (*github.Client, error) {
	// get a context
	ctx := context.Background()

	// read in the auth token from the place we store it
	authTok, err := getAuthToken()
	if err != nil {
		log.Log.Error("Unable to load GitHub auth token")
		return nil, err
	}

	// tbh I don't know what tc and ts do, but go-github says
	// I need them
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authTok},
	)

	tc := oauth2.NewClient(ctx, ts)

	// create and return the client
	return github.NewClient(tc), nil
}

// getAuthToken pulls the auth token from the location
// specified by configs
func getAuthToken() (string, error) {
	return "", nil
}
