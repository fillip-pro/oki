package do

import (
	"github.com/digitalocean/godo"
)

// Account provides the `godo.Account` information from Digital Ocean.
func Account() (*godo.Account, error) {
	ctx, client, err := NewContext()

	account, _, err := client.Account.Get(ctx)

	return account, err
}

// ListKeys lists the keys registered with Digital Ocean
func ListKeys() ([]godo.Key, error) {
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	ctx, client, err := NewContext()

	keys, _, err := client.Keys.List(ctx, opt)

	return keys, err
}
