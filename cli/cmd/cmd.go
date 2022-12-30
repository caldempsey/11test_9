package cmd

import (
	"github.com/unweave/unweave-v2/cli/config"
	"github.com/unweave/unweave-v2/client"
)

func InitUnweaveClient() *client.Client {
	// Get token. Priority: CLI flag > Project Token > User Token
	// TODO: Implement ProjectToken parsing

	token := config.UnweaveConfig.User.Token
	if config.AuthToken != "" {
		token = config.AuthToken
	}

	return client.NewClient(
		client.Config{
			ApiURL: config.UnweaveConfig.ApiURL,
			Token:  token,
		})
}
