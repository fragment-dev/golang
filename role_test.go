// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/fragment-go"
	"github.com/stainless-sdks/fragment-go/internal/testutil"
	"github.com/stainless-sdks/fragment-go/option"
)

func TestRoleNew(t *testing.T) {
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
	_, err := client.Roles.New(context.TODO(), fragment.RoleNewParams{
		Role: "admin",
	})
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoleList(t *testing.T) {
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
	_, err := client.Roles.List(context.TODO())
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
