package providers

import (
	"context"
	"os"

	"log"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// TokenSource provides the `AccessToken` for OAuth2 requests
// at Digital Ocean.
type TokenSource struct {
	AccessToken string
}

// DigitalOceanClient returns a Digital Ocean API client.
func DigitalOceanClient() (*DigitalOcean, error) {
	pat := os.Getenv("DO_TOKEN")

	digitalocean := &DigitalOcean{}

	context, client, err := NewContext(pat)

	if err != nil {
		log.Fatal(err)
	}

	digitalocean.context = context
	digitalocean.client = client

	return digitalocean, err
}

// Token returns the `TokenSource` for OAuth2 authentication.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// NewContext returns the `Context` for Digital Ocean requests.
func NewContext(token string) (context.Context, *godo.Client, error) {
	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	return ctx, client, nil
}
