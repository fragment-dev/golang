// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/fragment-go"
	"github.com/stainless-sdks/fragment-go/internal/testutil"
	"github.com/stainless-sdks/fragment-go/option"
)

func TestTransactionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.New(context.TODO(), fragment.TransactionNewParams{
		Account: fragment.TransactionNewParamsAccount{
			ID:         fragment.String("ext_account_YWJjMTIz"),
			ExternalID: fragment.String("acct_external_123"),
		},
		Allocations: []fragment.TransactionNewParamsAllocation{{
			Amount:    "1000",
			InvoiceID: "inv_abc123",
			Type:      "invoice_payin",
			User: fragment.TransactionNewParamsAllocationUserUnion{
				OfTransactionNewsAllocationUserID: &fragment.TransactionNewParamsAllocationUserID{
					ID: "user_abc123",
				},
			},
		}},
		Amount:     "-1000",
		Currency:   fragment.TransactionNewParamsCurrencyUsd,
		ExternalID: "bank_txn_123",
		Posted:     time.Now(),
		Tags: []fragment.TransactionNewParamsTag{{
			Key:   "department",
			Value: "engineering",
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

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "txn_abc123")
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Update(
		context.TODO(),
		"txn_abc123",
		fragment.TransactionUpdateParams{
			CurrentTransactionVersion: 0,
			Allocations: fragment.TransactionUpdateParamsAllocations{
				Create: []fragment.TransactionUpdateParamsAllocationsCreate{{
					Amount:    "1000",
					InvoiceID: "inv_abc123",
					Type:      "invoice_payin",
					User: fragment.TransactionUpdateParamsAllocationsCreateUserUnion{
						OfTransactionUpdatesAllocationsCreateUserID: &fragment.TransactionUpdateParamsAllocationsCreateUserID{
							ID: "user_abc123",
						},
					},
				}},
				Update: []fragment.TransactionUpdateParamsAllocationsUpdate{{
					ID:     "alloc_abc123",
					Amount: "2000",
				}},
			},
			Tags: fragment.TransactionUpdateParamsTags{
				Create: []fragment.TransactionUpdateParamsTagsCreate{{
					Key:   "department",
					Value: "engineering",
				}},
				Delete: []fragment.TransactionUpdateParamsTagsDelete{{
					Key: "key",
				}},
				Set: []fragment.TransactionUpdateParamsTagsSet{{
					Key:   "department",
					Value: "engineering",
				}},
				Update: []fragment.TransactionUpdateParamsTagsUpdate{{
					Key:   "department",
					Value: "engineering",
				}},
			},
		},
	)
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), fragment.TransactionListParams{
		Account:              fragment.String("ext_account_YWJjMTIz"),
		ReconciliationStatus: fragment.TransactionListParamsReconciliationStatusReconciled,
	})
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListHistory(t *testing.T) {
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
	_, err := client.Transactions.ListHistory(context.TODO(), "txn_abc123")
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Search(context.TODO(), fragment.TransactionSearchParams{
		Filter: fragment.TransactionSearchParamsFilter{
			Account: fragment.TransactionSearchParamsFilterAccount{
				Any: []fragment.TransactionSearchParamsFilterAccountAny{{
					ID:         fragment.String("ext_account_YWJjMTIz"),
					ExternalID: fragment.String("acct_external_123"),
				}},
			},
			Tags: fragment.TransactionSearchParamsFilterTags{
				All: []fragment.TransactionSearchParamsFilterTagsAll{{
					Key:   "department",
					Value: "engineering",
				}},
				Any: []fragment.TransactionSearchParamsFilterTagsAny{{
					Key:   "department",
					Value: "eng*",
				}},
			},
		},
		PageInfo: fragment.TransactionSearchParamsPageInfo{
			After: fragment.String("after"),
			Limit: fragment.Int(20),
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

func TestTransactionSearchAllocations(t *testing.T) {
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
	_, err := client.Transactions.SearchAllocations(context.TODO(), fragment.TransactionSearchAllocationsParams{
		Filter: fragment.TransactionSearchAllocationsParamsFilter{
			InvoiceID: fragment.TransactionSearchAllocationsParamsFilterInvoiceID{
				Any: []string{"inv_abc123"},
			},
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
