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

func TestProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.New(context.TODO(), fragment.ProductNewParams{
		Code:        "PROD_001",
		Description: fragment.String("Premium subscription service"),
		PaidByRoles: []fragment.ProductNewParamsPaidByRoleUnion{{
			OfProductNewsPaidByRoleName: &fragment.ProductNewParamsPaidByRoleName{
				Name: "buyer",
			},
		}},
		PaidToRoles: []fragment.ProductNewParamsPaidToRoleUnion{{
			OfProductNewsPaidToRoleName: &fragment.ProductNewParamsPaidToRoleName{
				Name: "seller",
			},
		}},
	})
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductGet(t *testing.T) {
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
	_, err := client.Products.Get(context.TODO(), "PROD_001")
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductList(t *testing.T) {
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
	_, err := client.Products.List(context.TODO())
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
