package client

import (
	"golang.org/x/oauth2"

	"github.com/dghubble/go-twitter/twitter"
	log "github.com/sirupsen/logrus"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ApiKey       string
	ApiSecretKey string
	BearerToken  string
}

type TokenSource struct {
	t *oauth2.Token
}

func (ts *TokenSource) Token() (*oauth2.Token, error) {
	return ts.t, nil
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func GetClient(creds *Credentials) (*twitter.Client, error) {
	src := &TokenSource{
		t: &oauth2.Token{
			AccessToken: creds.BearerToken,
			TokenType:   "Bearer",
		},
	}

	httpClient := oauth2.NewClient(oauth2.NoContext, src)

	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}
