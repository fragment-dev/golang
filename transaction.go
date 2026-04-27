// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/fragment-dev/fragment-billing-go/internal/apijson"
	"github.com/fragment-dev/fragment-billing-go/internal/apiquery"
	"github.com/fragment-dev/fragment-billing-go/internal/requestconfig"
	"github.com/fragment-dev/fragment-billing-go/option"
	"github.com/fragment-dev/fragment-billing-go/packages/param"
	"github.com/fragment-dev/fragment-billing-go/packages/respjson"
)

// Transaction sync operations
//
// TransactionService contains methods and other services that help with
// interacting with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.options = opts
	return
}

// Creates a transaction.
func (r *TransactionService) New(ctx context.Context, body TransactionNewParams, opts ...option.RequestOption) (res *TransactionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a transaction by ID or external ID.
func (r *TransactionService) Get(ctx context.Context, transactionRef string, opts ...option.RequestOption) (res *TransactionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if transactionRef == "" {
		err = errors.New("missing required transaction_ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("transactions/%s", url.PathEscape(transactionRef))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a transaction.
func (r *TransactionService) Update(ctx context.Context, transactionRef string, body TransactionUpdateParams, opts ...option.RequestOption) (res *TransactionUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if transactionRef == "" {
		err = errors.New("missing required transaction_ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("transactions/%s", url.PathEscape(transactionRef))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all transactions.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves the version history of a transaction.
func (r *TransactionService) ListHistory(ctx context.Context, transactionRef string, opts ...option.RequestOption) (res *TransactionListHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if transactionRef == "" {
		err = errors.New("missing required transaction_ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("transactions/%s/history", url.PathEscape(transactionRef))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Searches transactions.
func (r *TransactionService) Search(ctx context.Context, body TransactionSearchParams, opts ...option.RequestOption) (res *TransactionSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transactions/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Searches transaction allocations.
func (r *TransactionService) SearchAllocations(ctx context.Context, body TransactionSearchAllocationsParams, opts ...option.RequestOption) (res *TransactionSearchAllocationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transactions/allocations/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transaction object.
type Transaction struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// External account for the transaction.
	Account TransactionAccount `json:"account" api:"required"`
	// Allocations applied to the transaction.
	Allocations []TransactionAllocation `json:"allocations" api:"required"`
	// Transaction amount, as a string in the smallest currency unit, such as cents for
	// USD. Can be positive or negative.
	Amount string `json:"amount" api:"required"`
	// Timestamp when the transaction was created. Uses ISO 8601 format.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// ISO 4217 or crypto currency code.
	//
	// Any of "ADA", "BTC", "DAI", "ETH", "SOL", "USDC", "USDT", "USDG", "EURC",
	// "CADC", "CADT", "XLM", "UNI", "BCH", "LTC", "AAVE", "LINK", "MATIC", "PTS",
	// "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM",
	// "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN",
	// "BWP", "BYR", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUC",
	// "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR",
	// "FJD", "FKP", "GBP", "GEL", "GGP", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD",
	// "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "IMP", "INR", "IQD", "IRR",
	// "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD",
	// "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA",
	// "MKD", "MMK", "MNT", "MOP", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD",
	// "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR",
	// "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG",
	// "SEK", "SGD", "SHP", "SLL", "SOS", "SPL", "SRD", "SVC", "SYP", "STN", "SZL",
	// "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TVD", "TWD", "TZS", "UAH",
	// "UGX", "USD", "UYU", "UZS", "VEF", "VND", "VUV", "WST", "XAF", "XCD", "XOF",
	// "XPF", "YER", "ZAR", "ZMW", "LOGICAL", "CUSTOM".
	Currency TransactionCurrency `json:"currency" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Timestamp when the transaction was posted. Uses ISO 8601 format.
	Posted time.Time `json:"posted" api:"required" format:"date-time"`
	// Tags for the transaction.
	Tags []TransactionTag `json:"tags" api:"required"`
	// Amount not yet allocated, as a string.
	UnallocatedAmount string `json:"unallocated_amount" api:"required"`
	// Current version of the transaction.
	Version int64 `json:"version" api:"required"`
	// Timestamp when the transaction was last modified. Uses ISO 8601 format.
	Modified time.Time `json:"modified" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Account           respjson.Field
		Allocations       respjson.Field
		Amount            respjson.Field
		Created           respjson.Field
		Currency          respjson.Field
		ExternalID        respjson.Field
		Posted            respjson.Field
		Tags              respjson.Field
		UnallocatedAmount respjson.Field
		Version           respjson.Field
		Modified          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transaction) RawJSON() string { return r.JSON.raw }
func (r *Transaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External account for the transaction.
type TransactionAccount struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionAccount) RawJSON() string { return r.JSON.raw }
func (r *TransactionAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An allocation linking a transaction to an invoice.
type TransactionAllocation struct {
	// Allocated amount, as a positive string in the smallest currency unit, such as
	// cents for USD.
	Amount string `json:"amount" api:"required"`
	// Invoice the allocation is applied against.
	InvoiceID string `json:"invoice_id" api:"required"`
	// Type of allocation.
	//
	// Any of "invoice_payin", "invoice_payout".
	Type string `json:"type" api:"required"`
	// User associated with the allocation.
	User TransactionAllocationUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		InvoiceID   respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionAllocation) RawJSON() string { return r.JSON.raw }
func (r *TransactionAllocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User associated with the allocation.
type TransactionAllocationUser struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionAllocationUser) RawJSON() string { return r.JSON.raw }
func (r *TransactionAllocationUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ISO 4217 or crypto currency code.
type TransactionCurrency string

const (
	TransactionCurrencyAda     TransactionCurrency = "ADA"
	TransactionCurrencyBtc     TransactionCurrency = "BTC"
	TransactionCurrencyDai     TransactionCurrency = "DAI"
	TransactionCurrencyEth     TransactionCurrency = "ETH"
	TransactionCurrencySol     TransactionCurrency = "SOL"
	TransactionCurrencyUsdc    TransactionCurrency = "USDC"
	TransactionCurrencyUsdt    TransactionCurrency = "USDT"
	TransactionCurrencyUsdg    TransactionCurrency = "USDG"
	TransactionCurrencyEurc    TransactionCurrency = "EURC"
	TransactionCurrencyCadc    TransactionCurrency = "CADC"
	TransactionCurrencyCadt    TransactionCurrency = "CADT"
	TransactionCurrencyXlm     TransactionCurrency = "XLM"
	TransactionCurrencyUni     TransactionCurrency = "UNI"
	TransactionCurrencyBch     TransactionCurrency = "BCH"
	TransactionCurrencyLtc     TransactionCurrency = "LTC"
	TransactionCurrencyAave    TransactionCurrency = "AAVE"
	TransactionCurrencyLink    TransactionCurrency = "LINK"
	TransactionCurrencyMatic   TransactionCurrency = "MATIC"
	TransactionCurrencyPts     TransactionCurrency = "PTS"
	TransactionCurrencyAed     TransactionCurrency = "AED"
	TransactionCurrencyAfn     TransactionCurrency = "AFN"
	TransactionCurrencyAll     TransactionCurrency = "ALL"
	TransactionCurrencyAmd     TransactionCurrency = "AMD"
	TransactionCurrencyAng     TransactionCurrency = "ANG"
	TransactionCurrencyAoa     TransactionCurrency = "AOA"
	TransactionCurrencyArs     TransactionCurrency = "ARS"
	TransactionCurrencyAud     TransactionCurrency = "AUD"
	TransactionCurrencyAwg     TransactionCurrency = "AWG"
	TransactionCurrencyAzn     TransactionCurrency = "AZN"
	TransactionCurrencyBam     TransactionCurrency = "BAM"
	TransactionCurrencyBbd     TransactionCurrency = "BBD"
	TransactionCurrencyBdt     TransactionCurrency = "BDT"
	TransactionCurrencyBgn     TransactionCurrency = "BGN"
	TransactionCurrencyBhd     TransactionCurrency = "BHD"
	TransactionCurrencyBif     TransactionCurrency = "BIF"
	TransactionCurrencyBmd     TransactionCurrency = "BMD"
	TransactionCurrencyBnd     TransactionCurrency = "BND"
	TransactionCurrencyBob     TransactionCurrency = "BOB"
	TransactionCurrencyBrl     TransactionCurrency = "BRL"
	TransactionCurrencyBsd     TransactionCurrency = "BSD"
	TransactionCurrencyBtn     TransactionCurrency = "BTN"
	TransactionCurrencyBwp     TransactionCurrency = "BWP"
	TransactionCurrencyByr     TransactionCurrency = "BYR"
	TransactionCurrencyBzd     TransactionCurrency = "BZD"
	TransactionCurrencyCad     TransactionCurrency = "CAD"
	TransactionCurrencyCdf     TransactionCurrency = "CDF"
	TransactionCurrencyChf     TransactionCurrency = "CHF"
	TransactionCurrencyClp     TransactionCurrency = "CLP"
	TransactionCurrencyCny     TransactionCurrency = "CNY"
	TransactionCurrencyCop     TransactionCurrency = "COP"
	TransactionCurrencyCrc     TransactionCurrency = "CRC"
	TransactionCurrencyCuc     TransactionCurrency = "CUC"
	TransactionCurrencyCup     TransactionCurrency = "CUP"
	TransactionCurrencyCve     TransactionCurrency = "CVE"
	TransactionCurrencyCzk     TransactionCurrency = "CZK"
	TransactionCurrencyDjf     TransactionCurrency = "DJF"
	TransactionCurrencyDkk     TransactionCurrency = "DKK"
	TransactionCurrencyDop     TransactionCurrency = "DOP"
	TransactionCurrencyDzd     TransactionCurrency = "DZD"
	TransactionCurrencyEgp     TransactionCurrency = "EGP"
	TransactionCurrencyErn     TransactionCurrency = "ERN"
	TransactionCurrencyEtb     TransactionCurrency = "ETB"
	TransactionCurrencyEur     TransactionCurrency = "EUR"
	TransactionCurrencyFjd     TransactionCurrency = "FJD"
	TransactionCurrencyFkp     TransactionCurrency = "FKP"
	TransactionCurrencyGbp     TransactionCurrency = "GBP"
	TransactionCurrencyGel     TransactionCurrency = "GEL"
	TransactionCurrencyGgp     TransactionCurrency = "GGP"
	TransactionCurrencyGhs     TransactionCurrency = "GHS"
	TransactionCurrencyGip     TransactionCurrency = "GIP"
	TransactionCurrencyGmd     TransactionCurrency = "GMD"
	TransactionCurrencyGnf     TransactionCurrency = "GNF"
	TransactionCurrencyGtq     TransactionCurrency = "GTQ"
	TransactionCurrencyGyd     TransactionCurrency = "GYD"
	TransactionCurrencyHkd     TransactionCurrency = "HKD"
	TransactionCurrencyHnl     TransactionCurrency = "HNL"
	TransactionCurrencyHrk     TransactionCurrency = "HRK"
	TransactionCurrencyHtg     TransactionCurrency = "HTG"
	TransactionCurrencyHuf     TransactionCurrency = "HUF"
	TransactionCurrencyIdr     TransactionCurrency = "IDR"
	TransactionCurrencyIls     TransactionCurrency = "ILS"
	TransactionCurrencyImp     TransactionCurrency = "IMP"
	TransactionCurrencyInr     TransactionCurrency = "INR"
	TransactionCurrencyIqd     TransactionCurrency = "IQD"
	TransactionCurrencyIrr     TransactionCurrency = "IRR"
	TransactionCurrencyIsk     TransactionCurrency = "ISK"
	TransactionCurrencyJmd     TransactionCurrency = "JMD"
	TransactionCurrencyJod     TransactionCurrency = "JOD"
	TransactionCurrencyJpy     TransactionCurrency = "JPY"
	TransactionCurrencyKes     TransactionCurrency = "KES"
	TransactionCurrencyKgs     TransactionCurrency = "KGS"
	TransactionCurrencyKhr     TransactionCurrency = "KHR"
	TransactionCurrencyKmf     TransactionCurrency = "KMF"
	TransactionCurrencyKpw     TransactionCurrency = "KPW"
	TransactionCurrencyKrw     TransactionCurrency = "KRW"
	TransactionCurrencyKwd     TransactionCurrency = "KWD"
	TransactionCurrencyKyd     TransactionCurrency = "KYD"
	TransactionCurrencyKzt     TransactionCurrency = "KZT"
	TransactionCurrencyLak     TransactionCurrency = "LAK"
	TransactionCurrencyLbp     TransactionCurrency = "LBP"
	TransactionCurrencyLkr     TransactionCurrency = "LKR"
	TransactionCurrencyLrd     TransactionCurrency = "LRD"
	TransactionCurrencyLsl     TransactionCurrency = "LSL"
	TransactionCurrencyLyd     TransactionCurrency = "LYD"
	TransactionCurrencyMad     TransactionCurrency = "MAD"
	TransactionCurrencyMdl     TransactionCurrency = "MDL"
	TransactionCurrencyMga     TransactionCurrency = "MGA"
	TransactionCurrencyMkd     TransactionCurrency = "MKD"
	TransactionCurrencyMmk     TransactionCurrency = "MMK"
	TransactionCurrencyMnt     TransactionCurrency = "MNT"
	TransactionCurrencyMop     TransactionCurrency = "MOP"
	TransactionCurrencyMur     TransactionCurrency = "MUR"
	TransactionCurrencyMvr     TransactionCurrency = "MVR"
	TransactionCurrencyMwk     TransactionCurrency = "MWK"
	TransactionCurrencyMxn     TransactionCurrency = "MXN"
	TransactionCurrencyMyr     TransactionCurrency = "MYR"
	TransactionCurrencyMzn     TransactionCurrency = "MZN"
	TransactionCurrencyNad     TransactionCurrency = "NAD"
	TransactionCurrencyNgn     TransactionCurrency = "NGN"
	TransactionCurrencyNio     TransactionCurrency = "NIO"
	TransactionCurrencyNok     TransactionCurrency = "NOK"
	TransactionCurrencyNpr     TransactionCurrency = "NPR"
	TransactionCurrencyNzd     TransactionCurrency = "NZD"
	TransactionCurrencyOmr     TransactionCurrency = "OMR"
	TransactionCurrencyPab     TransactionCurrency = "PAB"
	TransactionCurrencyPen     TransactionCurrency = "PEN"
	TransactionCurrencyPgk     TransactionCurrency = "PGK"
	TransactionCurrencyPhp     TransactionCurrency = "PHP"
	TransactionCurrencyPkr     TransactionCurrency = "PKR"
	TransactionCurrencyPln     TransactionCurrency = "PLN"
	TransactionCurrencyPyg     TransactionCurrency = "PYG"
	TransactionCurrencyQar     TransactionCurrency = "QAR"
	TransactionCurrencyRon     TransactionCurrency = "RON"
	TransactionCurrencyRsd     TransactionCurrency = "RSD"
	TransactionCurrencyRub     TransactionCurrency = "RUB"
	TransactionCurrencyRwf     TransactionCurrency = "RWF"
	TransactionCurrencySar     TransactionCurrency = "SAR"
	TransactionCurrencySbd     TransactionCurrency = "SBD"
	TransactionCurrencyScr     TransactionCurrency = "SCR"
	TransactionCurrencySdg     TransactionCurrency = "SDG"
	TransactionCurrencySek     TransactionCurrency = "SEK"
	TransactionCurrencySgd     TransactionCurrency = "SGD"
	TransactionCurrencyShp     TransactionCurrency = "SHP"
	TransactionCurrencySll     TransactionCurrency = "SLL"
	TransactionCurrencySos     TransactionCurrency = "SOS"
	TransactionCurrencySpl     TransactionCurrency = "SPL"
	TransactionCurrencySrd     TransactionCurrency = "SRD"
	TransactionCurrencySvc     TransactionCurrency = "SVC"
	TransactionCurrencySyp     TransactionCurrency = "SYP"
	TransactionCurrencyStn     TransactionCurrency = "STN"
	TransactionCurrencySzl     TransactionCurrency = "SZL"
	TransactionCurrencyThb     TransactionCurrency = "THB"
	TransactionCurrencyTjs     TransactionCurrency = "TJS"
	TransactionCurrencyTmt     TransactionCurrency = "TMT"
	TransactionCurrencyTnd     TransactionCurrency = "TND"
	TransactionCurrencyTop     TransactionCurrency = "TOP"
	TransactionCurrencyTry     TransactionCurrency = "TRY"
	TransactionCurrencyTtd     TransactionCurrency = "TTD"
	TransactionCurrencyTvd     TransactionCurrency = "TVD"
	TransactionCurrencyTwd     TransactionCurrency = "TWD"
	TransactionCurrencyTzs     TransactionCurrency = "TZS"
	TransactionCurrencyUah     TransactionCurrency = "UAH"
	TransactionCurrencyUgx     TransactionCurrency = "UGX"
	TransactionCurrencyUsd     TransactionCurrency = "USD"
	TransactionCurrencyUyu     TransactionCurrency = "UYU"
	TransactionCurrencyUzs     TransactionCurrency = "UZS"
	TransactionCurrencyVef     TransactionCurrency = "VEF"
	TransactionCurrencyVnd     TransactionCurrency = "VND"
	TransactionCurrencyVuv     TransactionCurrency = "VUV"
	TransactionCurrencyWst     TransactionCurrency = "WST"
	TransactionCurrencyXaf     TransactionCurrency = "XAF"
	TransactionCurrencyXcd     TransactionCurrency = "XCD"
	TransactionCurrencyXof     TransactionCurrency = "XOF"
	TransactionCurrencyXpf     TransactionCurrency = "XPF"
	TransactionCurrencyYer     TransactionCurrency = "YER"
	TransactionCurrencyZar     TransactionCurrency = "ZAR"
	TransactionCurrencyZmw     TransactionCurrency = "ZMW"
	TransactionCurrencyLogical TransactionCurrency = "LOGICAL"
	TransactionCurrencyCustom  TransactionCurrency = "CUSTOM"
)

// A key-value tag pair.
type TransactionTag struct {
	// Tag key.
	Key string `json:"key" api:"required"`
	// Tag value.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionTag) RawJSON() string { return r.JSON.raw }
func (r *TransactionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionNewResponse struct {
	// Transaction object.
	Data Transaction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionGetResponse struct {
	// Transaction object.
	Data Transaction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionUpdateResponse struct {
	// Transaction object.
	Data Transaction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListResponse struct {
	// List of transaction objects matching the filter criteria.
	Data []Transaction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListHistoryResponse struct {
	// List of transaction versions over time, ordered by version, oldest first.
	Data []Transaction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionSearchResponse struct {
	// Deprecated. Use `data_v2.transactions` instead. Returns the full unpaginated
	// list of matching transactions.
	//
	// Deprecated: deprecated
	Data   []Transaction                   `json:"data" api:"required"`
	DataV2 TransactionSearchResponseDataV2 `json:"data_v2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		DataV2      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionSearchResponseDataV2 struct {
	// Pagination cursors.
	PageInfo TransactionSearchResponseDataV2PageInfo `json:"page_info" api:"required"`
	// Transactions matching the search criteria.
	Transactions []Transaction `json:"transactions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageInfo     respjson.Field
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchResponseDataV2) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchResponseDataV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination cursors.
type TransactionSearchResponseDataV2PageInfo struct {
	// Cursor to fetch the next page of results.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchResponseDataV2PageInfo) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchResponseDataV2PageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionSearchAllocationsResponse struct {
	// List of allocation search results.
	Data []TransactionSearchAllocationsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchAllocationsResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchAllocationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An allocation with a reference to its parent transaction.
type TransactionSearchAllocationsResponseData struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Allocated amount, as a positive string in the smallest currency unit, such as
	// cents for USD.
	Amount string `json:"amount" api:"required"`
	// Invoice the allocation is applied against.
	InvoiceID string `json:"invoice_id" api:"required"`
	// Timestamp when the parent transaction was posted. Uses ISO 8601 format.
	Posted time.Time `json:"posted" api:"required" format:"date-time"`
	// Transaction the allocation is applied to.
	Transaction TransactionSearchAllocationsResponseDataTransaction `json:"transaction" api:"required"`
	// Type of allocation.
	//
	// Any of "invoice_payin", "invoice_payout".
	Type string `json:"type" api:"required"`
	// User associated with the allocation.
	User TransactionSearchAllocationsResponseDataUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		InvoiceID   respjson.Field
		Posted      respjson.Field
		Transaction respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchAllocationsResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchAllocationsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transaction the allocation is applied to.
type TransactionSearchAllocationsResponseDataTransaction struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchAllocationsResponseDataTransaction) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchAllocationsResponseDataTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User associated with the allocation.
type TransactionSearchAllocationsResponseDataUser struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSearchAllocationsResponseDataUser) RawJSON() string { return r.JSON.raw }
func (r *TransactionSearchAllocationsResponseDataUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionNewParams struct {
	// External account for the transaction. Identify it by `id`, `external_id`, or
	// both.
	Account TransactionNewParamsAccount `json:"account,omitzero" api:"required"`
	// Allocations for the transaction. An empty array indicates unreconciled funds.
	Allocations []TransactionNewParamsAllocation `json:"allocations,omitzero" api:"required"`
	// Transaction amount, as a string in the smallest currency unit, such as cents for
	// USD. Can be positive or negative.
	Amount string `json:"amount" api:"required"`
	// ISO 4217 or crypto currency code.
	//
	// Any of "ADA", "BTC", "DAI", "ETH", "SOL", "USDC", "USDT", "USDG", "EURC",
	// "CADC", "CADT", "XLM", "UNI", "BCH", "LTC", "AAVE", "LINK", "MATIC", "PTS",
	// "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM",
	// "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN",
	// "BWP", "BYR", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUC",
	// "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR",
	// "FJD", "FKP", "GBP", "GEL", "GGP", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD",
	// "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "IMP", "INR", "IQD", "IRR",
	// "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD",
	// "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA",
	// "MKD", "MMK", "MNT", "MOP", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD",
	// "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR",
	// "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG",
	// "SEK", "SGD", "SHP", "SLL", "SOS", "SPL", "SRD", "SVC", "SYP", "STN", "SZL",
	// "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TVD", "TWD", "TZS", "UAH",
	// "UGX", "USD", "UYU", "UZS", "VEF", "VND", "VUV", "WST", "XAF", "XCD", "XOF",
	// "XPF", "YER", "ZAR", "ZMW", "LOGICAL", "CUSTOM".
	Currency TransactionNewParamsCurrency `json:"currency,omitzero" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Timestamp when the transaction was posted. Uses ISO 8601 format.
	Posted time.Time `json:"posted" api:"required" format:"date-time"`
	// Tags for the transaction.
	Tags []TransactionNewParamsTag `json:"tags,omitzero"`
	paramObj
}

func (r TransactionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External account for the transaction. Identify it by `id`, `external_id`, or
// both.
type TransactionNewParamsAccount struct {
	// FRAGMENT generated unique ID.
	ID param.Opt[string] `json:"id,omitzero"`
	// User-provided unique ID.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	paramObj
}

func (r TransactionNewParamsAccount) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParamsAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParamsAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An allocation linking a transaction to an invoice.
//
// The properties Amount, InvoiceID, Type, User are required.
type TransactionNewParamsAllocation struct {
	// Allocation amount, as a positive string in the smallest currency unit, such as
	// cents for USD.
	Amount string `json:"amount" api:"required"`
	// Invoice to allocate against.
	InvoiceID string `json:"invoice_id" api:"required"`
	// Type of allocation.
	//
	// Any of "invoice_payin", "invoice_payout".
	Type string `json:"type,omitzero" api:"required"`
	// Identifies a user by `id` or `external_id`.
	User TransactionNewParamsAllocationUserUnion `json:"user,omitzero" api:"required"`
	paramObj
}

func (r TransactionNewParamsAllocation) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParamsAllocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParamsAllocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TransactionNewParamsAllocation](
		"type", "invoice_payin", "invoice_payout",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransactionNewParamsAllocationUserUnion struct {
	OfTransactionNewsAllocationUserID         *TransactionNewParamsAllocationUserID         `json:",omitzero,inline"`
	OfTransactionNewsAllocationUserExternalID *TransactionNewParamsAllocationUserExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u TransactionNewParamsAllocationUserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransactionNewsAllocationUserID, u.OfTransactionNewsAllocationUserExternalID)
}
func (u *TransactionNewParamsAllocationUserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type TransactionNewParamsAllocationUserID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r TransactionNewParamsAllocationUserID) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParamsAllocationUserID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParamsAllocationUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type TransactionNewParamsAllocationUserExternalID struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r TransactionNewParamsAllocationUserExternalID) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParamsAllocationUserExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParamsAllocationUserExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ISO 4217 or crypto currency code.
type TransactionNewParamsCurrency string

const (
	TransactionNewParamsCurrencyAda     TransactionNewParamsCurrency = "ADA"
	TransactionNewParamsCurrencyBtc     TransactionNewParamsCurrency = "BTC"
	TransactionNewParamsCurrencyDai     TransactionNewParamsCurrency = "DAI"
	TransactionNewParamsCurrencyEth     TransactionNewParamsCurrency = "ETH"
	TransactionNewParamsCurrencySol     TransactionNewParamsCurrency = "SOL"
	TransactionNewParamsCurrencyUsdc    TransactionNewParamsCurrency = "USDC"
	TransactionNewParamsCurrencyUsdt    TransactionNewParamsCurrency = "USDT"
	TransactionNewParamsCurrencyUsdg    TransactionNewParamsCurrency = "USDG"
	TransactionNewParamsCurrencyEurc    TransactionNewParamsCurrency = "EURC"
	TransactionNewParamsCurrencyCadc    TransactionNewParamsCurrency = "CADC"
	TransactionNewParamsCurrencyCadt    TransactionNewParamsCurrency = "CADT"
	TransactionNewParamsCurrencyXlm     TransactionNewParamsCurrency = "XLM"
	TransactionNewParamsCurrencyUni     TransactionNewParamsCurrency = "UNI"
	TransactionNewParamsCurrencyBch     TransactionNewParamsCurrency = "BCH"
	TransactionNewParamsCurrencyLtc     TransactionNewParamsCurrency = "LTC"
	TransactionNewParamsCurrencyAave    TransactionNewParamsCurrency = "AAVE"
	TransactionNewParamsCurrencyLink    TransactionNewParamsCurrency = "LINK"
	TransactionNewParamsCurrencyMatic   TransactionNewParamsCurrency = "MATIC"
	TransactionNewParamsCurrencyPts     TransactionNewParamsCurrency = "PTS"
	TransactionNewParamsCurrencyAed     TransactionNewParamsCurrency = "AED"
	TransactionNewParamsCurrencyAfn     TransactionNewParamsCurrency = "AFN"
	TransactionNewParamsCurrencyAll     TransactionNewParamsCurrency = "ALL"
	TransactionNewParamsCurrencyAmd     TransactionNewParamsCurrency = "AMD"
	TransactionNewParamsCurrencyAng     TransactionNewParamsCurrency = "ANG"
	TransactionNewParamsCurrencyAoa     TransactionNewParamsCurrency = "AOA"
	TransactionNewParamsCurrencyArs     TransactionNewParamsCurrency = "ARS"
	TransactionNewParamsCurrencyAud     TransactionNewParamsCurrency = "AUD"
	TransactionNewParamsCurrencyAwg     TransactionNewParamsCurrency = "AWG"
	TransactionNewParamsCurrencyAzn     TransactionNewParamsCurrency = "AZN"
	TransactionNewParamsCurrencyBam     TransactionNewParamsCurrency = "BAM"
	TransactionNewParamsCurrencyBbd     TransactionNewParamsCurrency = "BBD"
	TransactionNewParamsCurrencyBdt     TransactionNewParamsCurrency = "BDT"
	TransactionNewParamsCurrencyBgn     TransactionNewParamsCurrency = "BGN"
	TransactionNewParamsCurrencyBhd     TransactionNewParamsCurrency = "BHD"
	TransactionNewParamsCurrencyBif     TransactionNewParamsCurrency = "BIF"
	TransactionNewParamsCurrencyBmd     TransactionNewParamsCurrency = "BMD"
	TransactionNewParamsCurrencyBnd     TransactionNewParamsCurrency = "BND"
	TransactionNewParamsCurrencyBob     TransactionNewParamsCurrency = "BOB"
	TransactionNewParamsCurrencyBrl     TransactionNewParamsCurrency = "BRL"
	TransactionNewParamsCurrencyBsd     TransactionNewParamsCurrency = "BSD"
	TransactionNewParamsCurrencyBtn     TransactionNewParamsCurrency = "BTN"
	TransactionNewParamsCurrencyBwp     TransactionNewParamsCurrency = "BWP"
	TransactionNewParamsCurrencyByr     TransactionNewParamsCurrency = "BYR"
	TransactionNewParamsCurrencyBzd     TransactionNewParamsCurrency = "BZD"
	TransactionNewParamsCurrencyCad     TransactionNewParamsCurrency = "CAD"
	TransactionNewParamsCurrencyCdf     TransactionNewParamsCurrency = "CDF"
	TransactionNewParamsCurrencyChf     TransactionNewParamsCurrency = "CHF"
	TransactionNewParamsCurrencyClp     TransactionNewParamsCurrency = "CLP"
	TransactionNewParamsCurrencyCny     TransactionNewParamsCurrency = "CNY"
	TransactionNewParamsCurrencyCop     TransactionNewParamsCurrency = "COP"
	TransactionNewParamsCurrencyCrc     TransactionNewParamsCurrency = "CRC"
	TransactionNewParamsCurrencyCuc     TransactionNewParamsCurrency = "CUC"
	TransactionNewParamsCurrencyCup     TransactionNewParamsCurrency = "CUP"
	TransactionNewParamsCurrencyCve     TransactionNewParamsCurrency = "CVE"
	TransactionNewParamsCurrencyCzk     TransactionNewParamsCurrency = "CZK"
	TransactionNewParamsCurrencyDjf     TransactionNewParamsCurrency = "DJF"
	TransactionNewParamsCurrencyDkk     TransactionNewParamsCurrency = "DKK"
	TransactionNewParamsCurrencyDop     TransactionNewParamsCurrency = "DOP"
	TransactionNewParamsCurrencyDzd     TransactionNewParamsCurrency = "DZD"
	TransactionNewParamsCurrencyEgp     TransactionNewParamsCurrency = "EGP"
	TransactionNewParamsCurrencyErn     TransactionNewParamsCurrency = "ERN"
	TransactionNewParamsCurrencyEtb     TransactionNewParamsCurrency = "ETB"
	TransactionNewParamsCurrencyEur     TransactionNewParamsCurrency = "EUR"
	TransactionNewParamsCurrencyFjd     TransactionNewParamsCurrency = "FJD"
	TransactionNewParamsCurrencyFkp     TransactionNewParamsCurrency = "FKP"
	TransactionNewParamsCurrencyGbp     TransactionNewParamsCurrency = "GBP"
	TransactionNewParamsCurrencyGel     TransactionNewParamsCurrency = "GEL"
	TransactionNewParamsCurrencyGgp     TransactionNewParamsCurrency = "GGP"
	TransactionNewParamsCurrencyGhs     TransactionNewParamsCurrency = "GHS"
	TransactionNewParamsCurrencyGip     TransactionNewParamsCurrency = "GIP"
	TransactionNewParamsCurrencyGmd     TransactionNewParamsCurrency = "GMD"
	TransactionNewParamsCurrencyGnf     TransactionNewParamsCurrency = "GNF"
	TransactionNewParamsCurrencyGtq     TransactionNewParamsCurrency = "GTQ"
	TransactionNewParamsCurrencyGyd     TransactionNewParamsCurrency = "GYD"
	TransactionNewParamsCurrencyHkd     TransactionNewParamsCurrency = "HKD"
	TransactionNewParamsCurrencyHnl     TransactionNewParamsCurrency = "HNL"
	TransactionNewParamsCurrencyHrk     TransactionNewParamsCurrency = "HRK"
	TransactionNewParamsCurrencyHtg     TransactionNewParamsCurrency = "HTG"
	TransactionNewParamsCurrencyHuf     TransactionNewParamsCurrency = "HUF"
	TransactionNewParamsCurrencyIdr     TransactionNewParamsCurrency = "IDR"
	TransactionNewParamsCurrencyIls     TransactionNewParamsCurrency = "ILS"
	TransactionNewParamsCurrencyImp     TransactionNewParamsCurrency = "IMP"
	TransactionNewParamsCurrencyInr     TransactionNewParamsCurrency = "INR"
	TransactionNewParamsCurrencyIqd     TransactionNewParamsCurrency = "IQD"
	TransactionNewParamsCurrencyIrr     TransactionNewParamsCurrency = "IRR"
	TransactionNewParamsCurrencyIsk     TransactionNewParamsCurrency = "ISK"
	TransactionNewParamsCurrencyJmd     TransactionNewParamsCurrency = "JMD"
	TransactionNewParamsCurrencyJod     TransactionNewParamsCurrency = "JOD"
	TransactionNewParamsCurrencyJpy     TransactionNewParamsCurrency = "JPY"
	TransactionNewParamsCurrencyKes     TransactionNewParamsCurrency = "KES"
	TransactionNewParamsCurrencyKgs     TransactionNewParamsCurrency = "KGS"
	TransactionNewParamsCurrencyKhr     TransactionNewParamsCurrency = "KHR"
	TransactionNewParamsCurrencyKmf     TransactionNewParamsCurrency = "KMF"
	TransactionNewParamsCurrencyKpw     TransactionNewParamsCurrency = "KPW"
	TransactionNewParamsCurrencyKrw     TransactionNewParamsCurrency = "KRW"
	TransactionNewParamsCurrencyKwd     TransactionNewParamsCurrency = "KWD"
	TransactionNewParamsCurrencyKyd     TransactionNewParamsCurrency = "KYD"
	TransactionNewParamsCurrencyKzt     TransactionNewParamsCurrency = "KZT"
	TransactionNewParamsCurrencyLak     TransactionNewParamsCurrency = "LAK"
	TransactionNewParamsCurrencyLbp     TransactionNewParamsCurrency = "LBP"
	TransactionNewParamsCurrencyLkr     TransactionNewParamsCurrency = "LKR"
	TransactionNewParamsCurrencyLrd     TransactionNewParamsCurrency = "LRD"
	TransactionNewParamsCurrencyLsl     TransactionNewParamsCurrency = "LSL"
	TransactionNewParamsCurrencyLyd     TransactionNewParamsCurrency = "LYD"
	TransactionNewParamsCurrencyMad     TransactionNewParamsCurrency = "MAD"
	TransactionNewParamsCurrencyMdl     TransactionNewParamsCurrency = "MDL"
	TransactionNewParamsCurrencyMga     TransactionNewParamsCurrency = "MGA"
	TransactionNewParamsCurrencyMkd     TransactionNewParamsCurrency = "MKD"
	TransactionNewParamsCurrencyMmk     TransactionNewParamsCurrency = "MMK"
	TransactionNewParamsCurrencyMnt     TransactionNewParamsCurrency = "MNT"
	TransactionNewParamsCurrencyMop     TransactionNewParamsCurrency = "MOP"
	TransactionNewParamsCurrencyMur     TransactionNewParamsCurrency = "MUR"
	TransactionNewParamsCurrencyMvr     TransactionNewParamsCurrency = "MVR"
	TransactionNewParamsCurrencyMwk     TransactionNewParamsCurrency = "MWK"
	TransactionNewParamsCurrencyMxn     TransactionNewParamsCurrency = "MXN"
	TransactionNewParamsCurrencyMyr     TransactionNewParamsCurrency = "MYR"
	TransactionNewParamsCurrencyMzn     TransactionNewParamsCurrency = "MZN"
	TransactionNewParamsCurrencyNad     TransactionNewParamsCurrency = "NAD"
	TransactionNewParamsCurrencyNgn     TransactionNewParamsCurrency = "NGN"
	TransactionNewParamsCurrencyNio     TransactionNewParamsCurrency = "NIO"
	TransactionNewParamsCurrencyNok     TransactionNewParamsCurrency = "NOK"
	TransactionNewParamsCurrencyNpr     TransactionNewParamsCurrency = "NPR"
	TransactionNewParamsCurrencyNzd     TransactionNewParamsCurrency = "NZD"
	TransactionNewParamsCurrencyOmr     TransactionNewParamsCurrency = "OMR"
	TransactionNewParamsCurrencyPab     TransactionNewParamsCurrency = "PAB"
	TransactionNewParamsCurrencyPen     TransactionNewParamsCurrency = "PEN"
	TransactionNewParamsCurrencyPgk     TransactionNewParamsCurrency = "PGK"
	TransactionNewParamsCurrencyPhp     TransactionNewParamsCurrency = "PHP"
	TransactionNewParamsCurrencyPkr     TransactionNewParamsCurrency = "PKR"
	TransactionNewParamsCurrencyPln     TransactionNewParamsCurrency = "PLN"
	TransactionNewParamsCurrencyPyg     TransactionNewParamsCurrency = "PYG"
	TransactionNewParamsCurrencyQar     TransactionNewParamsCurrency = "QAR"
	TransactionNewParamsCurrencyRon     TransactionNewParamsCurrency = "RON"
	TransactionNewParamsCurrencyRsd     TransactionNewParamsCurrency = "RSD"
	TransactionNewParamsCurrencyRub     TransactionNewParamsCurrency = "RUB"
	TransactionNewParamsCurrencyRwf     TransactionNewParamsCurrency = "RWF"
	TransactionNewParamsCurrencySar     TransactionNewParamsCurrency = "SAR"
	TransactionNewParamsCurrencySbd     TransactionNewParamsCurrency = "SBD"
	TransactionNewParamsCurrencyScr     TransactionNewParamsCurrency = "SCR"
	TransactionNewParamsCurrencySdg     TransactionNewParamsCurrency = "SDG"
	TransactionNewParamsCurrencySek     TransactionNewParamsCurrency = "SEK"
	TransactionNewParamsCurrencySgd     TransactionNewParamsCurrency = "SGD"
	TransactionNewParamsCurrencyShp     TransactionNewParamsCurrency = "SHP"
	TransactionNewParamsCurrencySll     TransactionNewParamsCurrency = "SLL"
	TransactionNewParamsCurrencySos     TransactionNewParamsCurrency = "SOS"
	TransactionNewParamsCurrencySpl     TransactionNewParamsCurrency = "SPL"
	TransactionNewParamsCurrencySrd     TransactionNewParamsCurrency = "SRD"
	TransactionNewParamsCurrencySvc     TransactionNewParamsCurrency = "SVC"
	TransactionNewParamsCurrencySyp     TransactionNewParamsCurrency = "SYP"
	TransactionNewParamsCurrencyStn     TransactionNewParamsCurrency = "STN"
	TransactionNewParamsCurrencySzl     TransactionNewParamsCurrency = "SZL"
	TransactionNewParamsCurrencyThb     TransactionNewParamsCurrency = "THB"
	TransactionNewParamsCurrencyTjs     TransactionNewParamsCurrency = "TJS"
	TransactionNewParamsCurrencyTmt     TransactionNewParamsCurrency = "TMT"
	TransactionNewParamsCurrencyTnd     TransactionNewParamsCurrency = "TND"
	TransactionNewParamsCurrencyTop     TransactionNewParamsCurrency = "TOP"
	TransactionNewParamsCurrencyTry     TransactionNewParamsCurrency = "TRY"
	TransactionNewParamsCurrencyTtd     TransactionNewParamsCurrency = "TTD"
	TransactionNewParamsCurrencyTvd     TransactionNewParamsCurrency = "TVD"
	TransactionNewParamsCurrencyTwd     TransactionNewParamsCurrency = "TWD"
	TransactionNewParamsCurrencyTzs     TransactionNewParamsCurrency = "TZS"
	TransactionNewParamsCurrencyUah     TransactionNewParamsCurrency = "UAH"
	TransactionNewParamsCurrencyUgx     TransactionNewParamsCurrency = "UGX"
	TransactionNewParamsCurrencyUsd     TransactionNewParamsCurrency = "USD"
	TransactionNewParamsCurrencyUyu     TransactionNewParamsCurrency = "UYU"
	TransactionNewParamsCurrencyUzs     TransactionNewParamsCurrency = "UZS"
	TransactionNewParamsCurrencyVef     TransactionNewParamsCurrency = "VEF"
	TransactionNewParamsCurrencyVnd     TransactionNewParamsCurrency = "VND"
	TransactionNewParamsCurrencyVuv     TransactionNewParamsCurrency = "VUV"
	TransactionNewParamsCurrencyWst     TransactionNewParamsCurrency = "WST"
	TransactionNewParamsCurrencyXaf     TransactionNewParamsCurrency = "XAF"
	TransactionNewParamsCurrencyXcd     TransactionNewParamsCurrency = "XCD"
	TransactionNewParamsCurrencyXof     TransactionNewParamsCurrency = "XOF"
	TransactionNewParamsCurrencyXpf     TransactionNewParamsCurrency = "XPF"
	TransactionNewParamsCurrencyYer     TransactionNewParamsCurrency = "YER"
	TransactionNewParamsCurrencyZar     TransactionNewParamsCurrency = "ZAR"
	TransactionNewParamsCurrencyZmw     TransactionNewParamsCurrency = "ZMW"
	TransactionNewParamsCurrencyLogical TransactionNewParamsCurrency = "LOGICAL"
	TransactionNewParamsCurrencyCustom  TransactionNewParamsCurrency = "CUSTOM"
)

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type TransactionNewParamsTag struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TransactionNewParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionUpdateParams struct {
	// Current version of the transaction. Must match the stored version.
	CurrentTransactionVersion int64 `json:"current_transaction_version" api:"required"`
	// Allocation updates.
	Allocations TransactionUpdateParamsAllocations `json:"allocations,omitzero"`
	// Tag updates.
	Tags TransactionUpdateParamsTags `json:"tags,omitzero"`
	paramObj
}

func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allocation updates.
type TransactionUpdateParamsAllocations struct {
	// Allocations to create.
	Create []TransactionUpdateParamsAllocationsCreate `json:"create,omitzero"`
	// Allocations to update.
	Update []TransactionUpdateParamsAllocationsUpdate `json:"update,omitzero"`
	paramObj
}

func (r TransactionUpdateParamsAllocations) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsAllocations
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsAllocations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An allocation linking a transaction to an invoice.
//
// The properties Amount, InvoiceID, Type, User are required.
type TransactionUpdateParamsAllocationsCreate struct {
	// Allocation amount, as a positive string in the smallest currency unit, such as
	// cents for USD.
	Amount string `json:"amount" api:"required"`
	// Invoice to allocate against.
	InvoiceID string `json:"invoice_id" api:"required"`
	// Type of allocation.
	//
	// Any of "invoice_payin", "invoice_payout".
	Type string `json:"type,omitzero" api:"required"`
	// Identifies a user by `id` or `external_id`.
	User TransactionUpdateParamsAllocationsCreateUserUnion `json:"user,omitzero" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsAllocationsCreate) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsAllocationsCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsAllocationsCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TransactionUpdateParamsAllocationsCreate](
		"type", "invoice_payin", "invoice_payout",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransactionUpdateParamsAllocationsCreateUserUnion struct {
	OfTransactionUpdatesAllocationsCreateUserID         *TransactionUpdateParamsAllocationsCreateUserID         `json:",omitzero,inline"`
	OfTransactionUpdatesAllocationsCreateUserExternalID *TransactionUpdateParamsAllocationsCreateUserExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u TransactionUpdateParamsAllocationsCreateUserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransactionUpdatesAllocationsCreateUserID, u.OfTransactionUpdatesAllocationsCreateUserExternalID)
}
func (u *TransactionUpdateParamsAllocationsCreateUserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type TransactionUpdateParamsAllocationsCreateUserID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsAllocationsCreateUserID) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsAllocationsCreateUserID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsAllocationsCreateUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type TransactionUpdateParamsAllocationsCreateUserExternalID struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsAllocationsCreateUserExternalID) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsAllocationsCreateUserExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsAllocationsCreateUserExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Amount are required.
type TransactionUpdateParamsAllocationsUpdate struct {
	// Allocation to update.
	ID string `json:"id" api:"required"`
	// Updated allocation amount, as a positive string in the smallest currency unit,
	// such as cents for USD.
	Amount string `json:"amount" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsAllocationsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsAllocationsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsAllocationsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tag updates.
type TransactionUpdateParamsTags struct {
	// Tags to create. The tag key must not already exist.
	Create []TransactionUpdateParamsTagsCreate `json:"create,omitzero"`
	// Tags to remove.
	Delete []TransactionUpdateParamsTagsDelete `json:"delete,omitzero"`
	// Tags to set. Creates a new tag or updates an existing tag.
	Set []TransactionUpdateParamsTagsSet `json:"set,omitzero"`
	// Tags to update. The tag key must already exist.
	Update []TransactionUpdateParamsTagsUpdate `json:"update,omitzero"`
	paramObj
}

func (r TransactionUpdateParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type TransactionUpdateParamsTagsCreate struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsTagsCreate) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsTagsCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsTagsCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Key is required.
type TransactionUpdateParamsTagsDelete struct {
	// Tag key to delete.
	Key string `json:"key" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsTagsDelete) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsTagsDelete
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsTagsDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type TransactionUpdateParamsTagsSet struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsTagsSet) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsTagsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsTagsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type TransactionUpdateParamsTagsUpdate struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TransactionUpdateParamsTagsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsTagsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsTagsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListParams struct {
	// Filter by account `id` or `external_id`. If the account does not exist, returns
	// an empty list.
	Account param.Opt[string] `query:"account,omitzero" json:"-"`
	// Filter by reconciliation status. `reconciled` returns transactions where
	// unallocated_amount is 0. `unreconciled` returns transactions where
	// unallocated_amount is not 0. Omit for all transactions.
	//
	// Any of "reconciled", "unreconciled".
	ReconciliationStatus TransactionListParamsReconciliationStatus `query:"reconciliation_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by reconciliation status. `reconciled` returns transactions where
// unallocated_amount is 0. `unreconciled` returns transactions where
// unallocated_amount is not 0. Omit for all transactions.
type TransactionListParamsReconciliationStatus string

const (
	TransactionListParamsReconciliationStatusReconciled   TransactionListParamsReconciliationStatus = "reconciled"
	TransactionListParamsReconciliationStatusUnreconciled TransactionListParamsReconciliationStatus = "unreconciled"
)

type TransactionSearchParams struct {
	// Filter for searching transactions.
	Filter TransactionSearchParamsFilter `json:"filter,omitzero" api:"required"`
	// Pagination parameters.
	PageInfo TransactionSearchParamsPageInfo `json:"page_info,omitzero"`
	paramObj
}

func (r TransactionSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter for searching transactions.
type TransactionSearchParamsFilter struct {
	// Account filter.
	Account TransactionSearchParamsFilterAccount `json:"account,omitzero"`
	// Tag-based filter criteria. When both `any` and `all` are provided, results must
	// match every entry in `all` AND at least one entry in `any`.
	Tags TransactionSearchParamsFilterTags `json:"tags,omitzero"`
	paramObj
}

func (r TransactionSearchParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account filter.
//
// The property Any is required.
type TransactionSearchParamsFilterAccount struct {
	// Match transactions belonging to any of these accounts, using OR logic.
	Any []TransactionSearchParamsFilterAccountAny `json:"any,omitzero" api:"required"`
	paramObj
}

func (r TransactionSearchParamsFilterAccount) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsFilterAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsFilterAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External account for the transaction. Identify it by `id`, `external_id`, or
// both.
type TransactionSearchParamsFilterAccountAny struct {
	// FRAGMENT generated unique ID.
	ID param.Opt[string] `json:"id,omitzero"`
	// User-provided unique ID.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	paramObj
}

func (r TransactionSearchParamsFilterAccountAny) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsFilterAccountAny
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsFilterAccountAny) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tag-based filter criteria. When both `any` and `all` are provided, results must
// match every entry in `all` AND at least one entry in `any`.
type TransactionSearchParamsFilterTags struct {
	// Returns transactions matching every specified tag, using AND logic.
	All []TransactionSearchParamsFilterTagsAll `json:"all,omitzero"`
	// Returns transactions matching at least one of the specified tags, using OR
	// logic.
	Any []TransactionSearchParamsFilterTagsAny `json:"any,omitzero"`
	paramObj
}

func (r TransactionSearchParamsFilterTags) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsFilterTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsFilterTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag filter.
//
// The properties Key, Value are required.
type TransactionSearchParamsFilterTagsAll struct {
	// Tag key to filter on. Must be an exact match; wildcards are not supported.
	Key string `json:"key" api:"required"`
	// Tag value pattern to filter on. Supports wildcards: `*` matches any characters,
	// `?` matches a single character. Use `\*` or `\?` to match literal asterisks or
	// question marks. Use `*` to match any value for the given key.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TransactionSearchParamsFilterTagsAll) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsFilterTagsAll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsFilterTagsAll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag filter.
//
// The properties Key, Value are required.
type TransactionSearchParamsFilterTagsAny struct {
	// Tag key to filter on. Must be an exact match; wildcards are not supported.
	Key string `json:"key" api:"required"`
	// Tag value pattern to filter on. Supports wildcards: `*` matches any characters,
	// `?` matches a single character. Use `\*` or `\?` to match literal asterisks or
	// question marks. Use `*` to match any value for the given key.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TransactionSearchParamsFilterTagsAny) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsFilterTagsAny
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsFilterTagsAny) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination parameters.
type TransactionSearchParamsPageInfo struct {
	// Cursor for fetching the next page of results.
	After param.Opt[string] `json:"after,omitzero"`
	// Number of results to return. Defaults to 20.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	paramObj
}

func (r TransactionSearchParamsPageInfo) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchParamsPageInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchParamsPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionSearchAllocationsParams struct {
	// Filter for searching transaction allocations.
	Filter TransactionSearchAllocationsParamsFilter `json:"filter,omitzero" api:"required"`
	paramObj
}

func (r TransactionSearchAllocationsParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchAllocationsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchAllocationsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter for searching transaction allocations.
//
// The property InvoiceID is required.
type TransactionSearchAllocationsParamsFilter struct {
	// Invoice ID filter.
	InvoiceID TransactionSearchAllocationsParamsFilterInvoiceID `json:"invoice_id,omitzero" api:"required"`
	paramObj
}

func (r TransactionSearchAllocationsParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchAllocationsParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchAllocationsParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invoice ID filter.
//
// The property Any is required.
type TransactionSearchAllocationsParamsFilterInvoiceID struct {
	// Match allocations where invoice_id is any of these values, using OR logic.
	Any []string `json:"any,omitzero" api:"required"`
	paramObj
}

func (r TransactionSearchAllocationsParamsFilterInvoiceID) MarshalJSON() (data []byte, err error) {
	type shadow TransactionSearchAllocationsParamsFilterInvoiceID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionSearchAllocationsParamsFilterInvoiceID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
