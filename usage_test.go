// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment_test

import (
	"context"
	"os"
	"testing"

	"github.com/fragment-dev/fragment-billing-go"
	"github.com/fragment-dev/fragment-billing-go/internal/testutil"
	"github.com/fragment-dev/fragment-billing-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := fragment.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	externalAccount, err := client.ExternalAccounts.New(context.TODO(), fragment.ExternalAccountNewParams{
		ExternalID: "ext_acc_123",
		Name:       "Checking Account",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", externalAccount.Data)
}
