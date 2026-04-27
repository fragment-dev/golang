// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/fragment-dev/fragment-billing-go"
	"github.com/fragment-dev/fragment-billing-go/internal/testutil"
	"github.com/fragment-dev/fragment-billing-go/option"
)

func TestExperimentalPaymentFlowNew(t *testing.T) {
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
	_, err := client.Experimental.PaymentFlows.New(context.TODO(), fragment.ExperimentalPaymentFlowNewParams{
		ExternalID: "pf_123",
		Invoice: fragment.ExperimentalPaymentFlowNewParamsInvoiceUnion{
			OfExperimentalPaymentFlowNewsInvoiceID: &fragment.ExperimentalPaymentFlowNewParamsInvoiceID{
				ID: "inv_abc123",
			},
		},
		Type: fragment.ExperimentalPaymentFlowNewParamsTypeSingleInvoiceSettlement,
	})
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentalPaymentFlowGet(t *testing.T) {
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
	_, err := client.Experimental.PaymentFlows.Get(context.TODO(), "pf_abc123")
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentalPaymentFlowSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Experimental.PaymentFlows.Search(context.TODO(), fragment.ExperimentalPaymentFlowSearchParams{
		InvoiceID: fragment.String("invoice_id"),
		PageInfo: fragment.ExperimentalPaymentFlowSearchParamsPageInfo{
			After: fragment.String("after"),
			Limit: fragment.Float(0),
		},
	})
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
