package s1_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stellaraf/go-s1"
	"github.com/stellaraf/go-utils/environment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Environment struct {
	APIToken      string `env:"API_TOKEN"`
	ManagementURL string `env:"MANAGEMENT_URL"`
}

var Env Environment

func init() {
	isVercel := os.Getenv("VERCEL") == "1" // if running in Vercel, value will be "1"
	isCI := os.Getenv("CI") == "true"      // if running in GitHub Actions, value will be "true"
	opts := &environment.EnvironmentOptions{
		DotEnv: !isVercel && !isCI,
	}
	err := environment.Load(&Env, opts)
	if err != nil {
		panic(err)
	}
}

func Test_S1(t *testing.T) {
	client, err := s1.New(Env.ManagementURL, Env.APIToken)
	require.NoError(t, err)
	ctx := context.Background()
	t.Run("accounts", func(t *testing.T) {
		t.Parallel()
		res, err := client.AccountsGet(ctx, nil)
		require.NoError(t, err)
		accounts, err := s1.ParseAccountsGetRes(res)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, accounts.StatusCode())
		assert.True(t, len(*accounts.JSON200.Data) > 0)
	})
}
