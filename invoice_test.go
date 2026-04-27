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

func TestInvoiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.New(context.TODO(), fragment.InvoiceNewParams{
		InvoiceID: "invoice_2024_001",
		LineItems: []fragment.InvoiceNewParamsLineItem{{
			Description: "Professional services for January 2026",
			ProductID:   "prod_1234567890",
			Type:        "payout",
			User: fragment.InvoiceNewParamsLineItemUserUnion{
				OfInvoiceNewsLineItemUserID: &fragment.InvoiceNewParamsLineItemUserID{
					ID: "user_abc123",
				},
			},
			Amount:       fragment.String("1000"),
			CurrencyCode: "USD",
			Price: fragment.InvoiceNewParamsLineItemPrice{
				Amount:    fragment.String("1000"),
				Quantity:  fragment.Int(2),
				UnitPrice: fragment.String("500"),
			},
			Tags: []fragment.InvoiceNewParamsLineItemTag{{
				Key:   "department",
				Value: "engineering",
			}},
		}},
		Tags: []fragment.InvoiceNewParamsTag{{
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

func TestInvoiceGet(t *testing.T) {
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
	_, err := client.Invoices.Get(context.TODO(), "inv_1234567890")
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.Update(
		context.TODO(),
		"inv_1234567890",
		fragment.InvoiceUpdateParams{
			CurrentInvoiceVersion: 3,
			LineItems: fragment.InvoiceUpdateParamsLineItems{
				Create: []fragment.InvoiceUpdateParamsLineItemsCreate{{
					Description: "Professional services for January 2026",
					ProductID:   "prod_1234567890",
					Type:        "payout",
					User: fragment.InvoiceUpdateParamsLineItemsCreateUserUnion{
						OfInvoiceUpdatesLineItemsCreateUserID: &fragment.InvoiceUpdateParamsLineItemsCreateUserID{
							ID: "user_abc123",
						},
					},
					Amount:       fragment.String("1000"),
					CurrencyCode: "USD",
					Price: fragment.InvoiceUpdateParamsLineItemsCreatePrice{
						Amount:    fragment.String("1000"),
						Quantity:  fragment.Int(2),
						UnitPrice: fragment.String("500"),
					},
					Tags: []fragment.InvoiceUpdateParamsLineItemsCreateTag{{
						Key:   "department",
						Value: "engineering",
					}},
				}},
				Delete: []fragment.InvoiceUpdateParamsLineItemsDelete{{
					ID: "id",
				}},
				Update: []fragment.InvoiceUpdateParamsLineItemsUpdate{{
					ID:          "li_1234567890",
					Description: fragment.String("description"),
					Price: fragment.InvoiceUpdateParamsLineItemsUpdatePrice{
						Quantity:  2,
						UnitPrice: "500",
						Amount:    fragment.String("2000"),
					},
					Tags: fragment.InvoiceUpdateParamsLineItemsUpdateTags{
						Create: []fragment.InvoiceUpdateParamsLineItemsUpdateTagsCreate{{
							Key:   "department",
							Value: "engineering",
						}},
						Delete: []fragment.InvoiceUpdateParamsLineItemsUpdateTagsDelete{{
							Key: "key",
						}},
						Set: []fragment.InvoiceUpdateParamsLineItemsUpdateTagsSet{{
							Key:   "department",
							Value: "engineering",
						}},
						Update: []fragment.InvoiceUpdateParamsLineItemsUpdateTagsUpdate{{
							Key:   "department",
							Value: "engineering",
						}},
					},
				}},
			},
			Tags: fragment.InvoiceUpdateParamsTags{
				Create: []fragment.InvoiceUpdateParamsTagsCreate{{
					Key:   "department",
					Value: "engineering",
				}},
				Delete: []fragment.InvoiceUpdateParamsTagsDelete{{
					Key: "key",
				}},
				Set: []fragment.InvoiceUpdateParamsTagsSet{{
					Key:   "department",
					Value: "engineering",
				}},
				Update: []fragment.InvoiceUpdateParamsTagsUpdate{{
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

func TestInvoiceList(t *testing.T) {
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
	_, err := client.Invoices.List(context.TODO())
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceListHistory(t *testing.T) {
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
	_, err := client.Invoices.ListHistory(context.TODO(), "inv_1234567890")
	if err != nil {
		var apierr *fragment.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.Search(context.TODO(), fragment.InvoiceSearchParams{
		Filter: fragment.InvoiceSearchParamsFilter{
			Status: "open",
			Tags: fragment.InvoiceSearchParamsFilterTags{
				All: []fragment.InvoiceSearchParamsFilterTagsAll{{
					Key:   "department",
					Value: "engineering",
				}},
				Any: []fragment.InvoiceSearchParamsFilterTagsAny{{
					Key:   "department",
					Value: "eng*",
				}},
			},
			TransactionTags: fragment.InvoiceSearchParamsFilterTransactionTags{
				All: []fragment.InvoiceSearchParamsFilterTransactionTagsAll{{
					Key:   "department",
					Value: "engineering",
				}},
				Any: []fragment.InvoiceSearchParamsFilterTransactionTagsAny{{
					Key:   "department",
					Value: "eng*",
				}},
			},
			Users: fragment.InvoiceSearchParamsFilterUsers{
				All: []fragment.InvoiceSearchParamsFilterUsersAllUnion{{
					OfInvoiceSearchsFilterUsersAllID: &fragment.InvoiceSearchParamsFilterUsersAllID{
						ID: "user_abc123",
					},
				}},
				Any: []fragment.InvoiceSearchParamsFilterUsersAnyUnion{{
					OfInvoiceSearchsFilterUsersAnyID: &fragment.InvoiceSearchParamsFilterUsersAnyID{
						ID: "user_abc123",
					},
				}},
			},
		},
		PageInfo: fragment.InvoiceSearchParamsPageInfo{
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
