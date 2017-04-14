package do

import (
	"context"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

var pat = "3719c85afcc0f16d380d0b10431f7f6fc6988dbb30de047d7aeeb7aa4e6ecfc4"

// TokenSource provides the `AccessToken` for OAuth2 requests
// at Digital Ocean.
type TokenSource struct {
	AccessToken string
}

// Token returns the `TokenSource` for OAuth2 authentication.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// NewContext returns the `Context` for Digital Ocean requests.
func NewContext() (context.Context, *godo.Client, error) {
	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	return ctx, client, nil
}
