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

	"github.com/fragment-dev/golang/internal/apijson"
	"github.com/fragment-dev/golang/internal/requestconfig"
	"github.com/fragment-dev/golang/option"
	"github.com/fragment-dev/golang/packages/param"
	"github.com/fragment-dev/golang/packages/respjson"
)

// Invoice management operations
//
// InvoiceService contains methods and other services that help with interacting
// with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r InvoiceService) {
	r = InvoiceService{}
	r.options = opts
	return
}

// Creates an invoice.
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *InvoiceNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an invoice.
func (r *InvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *InvoiceGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoices/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an invoice.
func (r *InvoiceService) Update(ctx context.Context, id string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *InvoiceUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoices/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all invoices.
func (r *InvoiceService) List(ctx context.Context, opts ...option.RequestOption) (res *InvoiceListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the version history of an invoice.
func (r *InvoiceService) ListHistory(ctx context.Context, id string, opts ...option.RequestOption) (res *InvoiceListHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoices/%s/history", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Searches invoices.
func (r *InvoiceService) Search(ctx context.Context, body InvoiceSearchParams, opts ...option.RequestOption) (res *InvoiceSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "invoices/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Invoice object.
type Invoice struct {
	// Unique invoice ID.
	ID string `json:"id" api:"required"`
	// Timestamp when the invoice was created. Uses ISO 8601 format.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Status of the invoice. Deprecated.
	//
	// Any of "active".
	//
	// Deprecated: deprecated
	Status InvoiceStatus `json:"status" api:"required"`
	// Tags for the invoice.
	Tags []InvoiceTag `json:"tags" api:"required"`
	// Current version of the invoice.
	Version float64 `json:"version" api:"required"`
	// Workspace the invoice belongs to.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Line items for the invoice.
	LineItems []InvoiceLineItem `json:"line_items"`
	// Timestamp when the invoice was last modified. Uses ISO 8601 format.
	Modified time.Time `json:"modified" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Version     respjson.Field
		WorkspaceID respjson.Field
		LineItems   respjson.Field
		Modified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invoice) RawJSON() string { return r.JSON.raw }
func (r *Invoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the invoice. Deprecated.
type InvoiceStatus string

const (
	InvoiceStatusActive InvoiceStatus = "active"
)

// A key-value tag pair.
type InvoiceTag struct {
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
func (r InvoiceTag) RawJSON() string { return r.JSON.raw }
func (r *InvoiceTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invoice line item.
type InvoiceLineItem struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	// Deprecated, use price.amount instead.
	//
	// Deprecated: deprecated
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
	CurrencyCode string `json:"currency_code" api:"required"`
	// Description of the line item.
	Description string `json:"description" api:"required"`
	// Price breakdown.
	Price InvoiceLineItemPrice `json:"price" api:"required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id" api:"required"`
	// Tags for the line item.
	Tags []InvoiceLineItemTag `json:"tags" api:"required"`
	// Type of the line item.
	//
	// Any of "payin", "payout".
	Type string `json:"type" api:"required"`
	// User-provided unique external ID.
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		CurrencyCode respjson.Field
		Description  respjson.Field
		Price        respjson.Field
		ProductID    respjson.Field
		Tags         respjson.Field
		Type         respjson.Field
		UserID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItem) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price breakdown.
type InvoiceLineItemPrice struct {
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	Amount string `json:"amount" api:"required"`
	// Number of units.
	Quantity int64 `json:"quantity" api:"required"`
	// Unit price as a string in the smallest currency unit, such as cents for USD.
	UnitPrice string `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Quantity    respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceLineItemPrice) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair.
type InvoiceLineItemTag struct {
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
func (r InvoiceLineItemTag) RawJSON() string { return r.JSON.raw }
func (r *InvoiceLineItemTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceNewResponse struct {
	// Invoice object.
	Data Invoice `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceGetResponse struct {
	// Invoice with balance details.
	Data InvoiceGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invoice with balance details.
type InvoiceGetResponseData struct {
	// Invoice-level balances by currency.
	Balances []InvoiceGetResponseDataBalance `json:"balances" api:"required"`
	// Payments allocated to the invoice.
	Payments []InvoiceGetResponseDataPayment `json:"payments" api:"required"`
	// Users involved in the invoice.
	Users []InvoiceGetResponseDataUser `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balances    respjson.Field
		Payments    respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Invoice
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceGetResponseDataBalance struct {
	// ISO 4217 or crypto currency code.
	Currency string `json:"currency" api:"required"`
	// Net balance breakdown.
	Net InvoiceGetResponseDataBalanceNet `json:"net" api:"required"`
	// Payins balance breakdown.
	Payins InvoiceGetResponseDataBalancePayins `json:"payins" api:"required"`
	// Payouts balance breakdown.
	Payouts InvoiceGetResponseDataBalancePayouts `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Net         respjson.Field
		Payins      respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataBalance) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Net balance breakdown.
type InvoiceGetResponseDataBalanceNet struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataBalanceNet) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataBalanceNet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payins balance breakdown.
type InvoiceGetResponseDataBalancePayins struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataBalancePayins) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataBalancePayins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payouts balance breakdown.
type InvoiceGetResponseDataBalancePayouts struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataBalancePayouts) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataBalancePayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A payment allocated to the invoice.
type InvoiceGetResponseDataPayment struct {
	// Amount allocated as a string in the smallest currency unit, such as cents for
	// USD.
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
	Currency string `json:"currency" api:"required"`
	// Timestamp when the parent transaction was posted. Uses ISO 8601 format.
	Posted time.Time `json:"posted" api:"required" format:"date-time"`
	// Transaction the payment is applied to.
	Transaction InvoiceGetResponseDataPaymentTransaction `json:"transaction" api:"required"`
	// Type of the payment.
	//
	// Any of "payin", "payout".
	Type string `json:"type" api:"required"`
	// User associated with the payment.
	User InvoiceGetResponseDataPaymentUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		Posted      respjson.Field
		Transaction respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataPayment) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transaction the payment is applied to.
type InvoiceGetResponseDataPaymentTransaction struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Tags from the parent transaction.
	Tags []InvoiceGetResponseDataPaymentTransactionTag `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataPaymentTransaction) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataPaymentTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair.
type InvoiceGetResponseDataPaymentTransactionTag struct {
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
func (r InvoiceGetResponseDataPaymentTransactionTag) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataPaymentTransactionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User associated with the payment.
type InvoiceGetResponseDataPaymentUser struct {
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
func (r InvoiceGetResponseDataPaymentUser) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataPaymentUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceGetResponseDataUser struct {
	// User-provided unique external ID.
	ID string `json:"id" api:"required"`
	// Per-currency balance breakdown for the user.
	Balances []InvoiceGetResponseDataUserBalance `json:"balances" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Balances    respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataUser) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceGetResponseDataUserBalance struct {
	// ISO 4217 or crypto currency code.
	Currency string `json:"currency" api:"required"`
	// Net balance breakdown.
	Net InvoiceGetResponseDataUserBalanceNet `json:"net" api:"required"`
	// Payins balance breakdown.
	Payins InvoiceGetResponseDataUserBalancePayins `json:"payins" api:"required"`
	// Payouts balance breakdown.
	Payouts InvoiceGetResponseDataUserBalancePayouts `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Net         respjson.Field
		Payins      respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataUserBalance) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataUserBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Net balance breakdown.
type InvoiceGetResponseDataUserBalanceNet struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataUserBalanceNet) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataUserBalanceNet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payins balance breakdown.
type InvoiceGetResponseDataUserBalancePayins struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataUserBalancePayins) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataUserBalancePayins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payouts balance breakdown.
type InvoiceGetResponseDataUserBalancePayouts struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseDataUserBalancePayouts) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseDataUserBalancePayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceUpdateResponse struct {
	// Invoice object.
	Data Invoice `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListResponse struct {
	// List of invoices.
	Data []Invoice `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListHistoryResponse struct {
	// Version history of the invoice.
	Data []Invoice `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceListHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results for invoices.
type InvoiceSearchResponse struct {
	// Search results for invoices.
	Data InvoiceSearchResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results for invoices.
type InvoiceSearchResponseData struct {
	// Invoices matching the search criteria.
	Invoices []InvoiceSearchResponseDataInvoice `json:"invoices" api:"required"`
	// Pagination cursors.
	PageInfo InvoiceSearchResponseDataPageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoices    respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invoice with balance details.
type InvoiceSearchResponseDataInvoice struct {
	// Invoice-level balances by currency.
	Balances []InvoiceSearchResponseDataInvoiceBalance `json:"balances" api:"required"`
	// Payments allocated to the invoice.
	Payments []InvoiceSearchResponseDataInvoicePayment `json:"payments" api:"required"`
	// Users involved in the invoice.
	Users []InvoiceSearchResponseDataInvoiceUser `json:"users" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balances    respjson.Field
		Payments    respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Invoice
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoice) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceSearchResponseDataInvoiceBalance struct {
	// ISO 4217 or crypto currency code.
	Currency string `json:"currency" api:"required"`
	// Net balance breakdown.
	Net InvoiceSearchResponseDataInvoiceBalanceNet `json:"net" api:"required"`
	// Payins balance breakdown.
	Payins InvoiceSearchResponseDataInvoiceBalancePayins `json:"payins" api:"required"`
	// Payouts balance breakdown.
	Payouts InvoiceSearchResponseDataInvoiceBalancePayouts `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Net         respjson.Field
		Payins      respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceBalance) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Net balance breakdown.
type InvoiceSearchResponseDataInvoiceBalanceNet struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceBalanceNet) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceBalanceNet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payins balance breakdown.
type InvoiceSearchResponseDataInvoiceBalancePayins struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceBalancePayins) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceBalancePayins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payouts balance breakdown.
type InvoiceSearchResponseDataInvoiceBalancePayouts struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceBalancePayouts) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceBalancePayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A payment allocated to the invoice.
type InvoiceSearchResponseDataInvoicePayment struct {
	// Amount allocated as a string in the smallest currency unit, such as cents for
	// USD.
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
	Currency string `json:"currency" api:"required"`
	// Timestamp when the parent transaction was posted. Uses ISO 8601 format.
	Posted time.Time `json:"posted" api:"required" format:"date-time"`
	// Transaction the payment is applied to.
	Transaction InvoiceSearchResponseDataInvoicePaymentTransaction `json:"transaction" api:"required"`
	// Type of the payment.
	//
	// Any of "payin", "payout".
	Type string `json:"type" api:"required"`
	// User associated with the payment.
	User InvoiceSearchResponseDataInvoicePaymentUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		Posted      respjson.Field
		Transaction respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoicePayment) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoicePayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transaction the payment is applied to.
type InvoiceSearchResponseDataInvoicePaymentTransaction struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Tags from the parent transaction.
	Tags []InvoiceSearchResponseDataInvoicePaymentTransactionTag `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoicePaymentTransaction) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoicePaymentTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair.
type InvoiceSearchResponseDataInvoicePaymentTransactionTag struct {
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
func (r InvoiceSearchResponseDataInvoicePaymentTransactionTag) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoicePaymentTransactionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User associated with the payment.
type InvoiceSearchResponseDataInvoicePaymentUser struct {
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
func (r InvoiceSearchResponseDataInvoicePaymentUser) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoicePaymentUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceSearchResponseDataInvoiceUser struct {
	// User-provided unique external ID.
	ID string `json:"id" api:"required"`
	// Per-currency balance breakdown for the user.
	Balances []InvoiceSearchResponseDataInvoiceUserBalance `json:"balances" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Balances    respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceUser) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceSearchResponseDataInvoiceUserBalance struct {
	// ISO 4217 or crypto currency code.
	Currency string `json:"currency" api:"required"`
	// Net balance breakdown.
	Net InvoiceSearchResponseDataInvoiceUserBalanceNet `json:"net" api:"required"`
	// Payins balance breakdown.
	Payins InvoiceSearchResponseDataInvoiceUserBalancePayins `json:"payins" api:"required"`
	// Payouts balance breakdown.
	Payouts InvoiceSearchResponseDataInvoiceUserBalancePayouts `json:"payouts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Net         respjson.Field
		Payins      respjson.Field
		Payouts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceUserBalance) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceUserBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Net balance breakdown.
type InvoiceSearchResponseDataInvoiceUserBalanceNet struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceUserBalanceNet) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceUserBalanceNet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payins balance breakdown.
type InvoiceSearchResponseDataInvoiceUserBalancePayins struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceUserBalancePayins) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceUserBalancePayins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payouts balance breakdown.
type InvoiceSearchResponseDataInvoiceUserBalancePayouts struct {
	// Actual amount as a string in the smallest currency unit, such as cents for USD.
	Actual string `json:"actual" api:"required"`
	// Expected amount as a string in the smallest currency unit, such as cents for
	// USD.
	Expected string `json:"expected" api:"required"`
	// Remaining amount as a string in the smallest currency unit, such as cents for
	// USD.
	Remaining string `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actual      respjson.Field
		Expected    respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceSearchResponseDataInvoiceUserBalancePayouts) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataInvoiceUserBalancePayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination cursors.
type InvoiceSearchResponseDataPageInfo struct {
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
func (r InvoiceSearchResponseDataPageInfo) RawJSON() string { return r.JSON.raw }
func (r *InvoiceSearchResponseDataPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceNewParams struct {
	// Unique ID for the invoice.
	InvoiceID string `json:"invoice_id" api:"required"`
	// Line items to create with the invoice.
	LineItems []InvoiceNewParamsLineItem `json:"line_items,omitzero" api:"required"`
	// Tags for the invoice.
	Tags []InvoiceNewParamsTag `json:"tags,omitzero"`
	paramObj
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data to create a line item.
//
// The properties Description, ProductID, Type, User are required.
type InvoiceNewParamsLineItem struct {
	// Description of the line item.
	Description string `json:"description" api:"required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id" api:"required"`
	// Type of the line item.
	//
	// Any of "payin", "payout".
	Type string `json:"type,omitzero" api:"required"`
	// Identifies a user by `id` or `external_id`.
	User InvoiceNewParamsLineItemUserUnion `json:"user,omitzero" api:"required"`
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	// Deprecated, use price instead.
	//
	// Deprecated: deprecated
	Amount param.Opt[string] `json:"amount,omitzero"`
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
	CurrencyCode string `json:"currency_code,omitzero"`
	// Price breakdown. Provide amount, or unit_price and quantity, or all three.
	Price InvoiceNewParamsLineItemPrice `json:"price,omitzero"`
	// Tags for the line item.
	Tags []InvoiceNewParamsLineItemTag `json:"tags,omitzero"`
	paramObj
}

func (r InvoiceNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsLineItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InvoiceNewParamsLineItem](
		"type", "payin", "payout",
	)
	apijson.RegisterFieldValidator[InvoiceNewParamsLineItem](
		"currency_code", "ADA", "BTC", "DAI", "ETH", "SOL", "USDC", "USDT", "USDG", "EURC", "CADC", "CADT", "XLM", "UNI", "BCH", "LTC", "AAVE", "LINK", "MATIC", "PTS", "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN", "BWP", "BYR", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GGP", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "IMP", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS", "SPL", "SRD", "SVC", "SYP", "STN", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TVD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VEF", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW", "LOGICAL", "CUSTOM",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvoiceNewParamsLineItemUserUnion struct {
	OfInvoiceNewsLineItemUserID         *InvoiceNewParamsLineItemUserID         `json:",omitzero,inline"`
	OfInvoiceNewsLineItemUserExternalID *InvoiceNewParamsLineItemUserExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u InvoiceNewParamsLineItemUserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInvoiceNewsLineItemUserID, u.OfInvoiceNewsLineItemUserExternalID)
}
func (u *InvoiceNewParamsLineItemUserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type InvoiceNewParamsLineItemUserID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r InvoiceNewParamsLineItemUserID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsLineItemUserID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsLineItemUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type InvoiceNewParamsLineItemUserExternalID struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r InvoiceNewParamsLineItemUserExternalID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsLineItemUserExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsLineItemUserExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price breakdown. Provide amount, or unit_price and quantity, or all three.
type InvoiceNewParamsLineItemPrice struct {
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	// Required if unit_price and quantity are not provided.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Number of units for the line item.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Price per unit as a string in the smallest currency unit, such as cents for USD.
	UnitPrice param.Opt[string] `json:"unit_price,omitzero"`
	paramObj
}

func (r InvoiceNewParamsLineItemPrice) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsLineItemPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsLineItemPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceNewParamsLineItemTag struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceNewParamsLineItemTag) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsLineItemTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsLineItemTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceNewParamsTag struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceNewParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceUpdateParams struct {
	// Current version of the invoice. Must match the stored version.
	CurrentInvoiceVersion float64 `json:"current_invoice_version" api:"required"`
	// Line item updates.
	LineItems InvoiceUpdateParamsLineItems `json:"line_items,omitzero"`
	// Tag updates.
	Tags InvoiceUpdateParamsTags `json:"tags,omitzero"`
	paramObj
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Line item updates.
type InvoiceUpdateParamsLineItems struct {
	// Line items to add to the invoice.
	Create []InvoiceUpdateParamsLineItemsCreate `json:"create,omitzero"`
	// Line items to remove from the invoice.
	Delete []InvoiceUpdateParamsLineItemsDelete `json:"delete,omitzero"`
	// Existing line items to update.
	Update []InvoiceUpdateParamsLineItemsUpdate `json:"update,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsLineItems) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItems
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItems) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data to create a line item.
//
// The properties Description, ProductID, Type, User are required.
type InvoiceUpdateParamsLineItemsCreate struct {
	// Description of the line item.
	Description string `json:"description" api:"required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id" api:"required"`
	// Type of the line item.
	//
	// Any of "payin", "payout".
	Type string `json:"type,omitzero" api:"required"`
	// Identifies a user by `id` or `external_id`.
	User InvoiceUpdateParamsLineItemsCreateUserUnion `json:"user,omitzero" api:"required"`
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	// Deprecated, use price instead.
	//
	// Deprecated: deprecated
	Amount param.Opt[string] `json:"amount,omitzero"`
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
	CurrencyCode string `json:"currency_code,omitzero"`
	// Price breakdown. Provide amount, or unit_price and quantity, or all three.
	Price InvoiceUpdateParamsLineItemsCreatePrice `json:"price,omitzero"`
	// Tags for the line item.
	Tags []InvoiceUpdateParamsLineItemsCreateTag `json:"tags,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsCreate) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InvoiceUpdateParamsLineItemsCreate](
		"type", "payin", "payout",
	)
	apijson.RegisterFieldValidator[InvoiceUpdateParamsLineItemsCreate](
		"currency_code", "ADA", "BTC", "DAI", "ETH", "SOL", "USDC", "USDT", "USDG", "EURC", "CADC", "CADT", "XLM", "UNI", "BCH", "LTC", "AAVE", "LINK", "MATIC", "PTS", "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN", "BWP", "BYR", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GGP", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "IMP", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS", "SPL", "SRD", "SVC", "SYP", "STN", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TVD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VEF", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW", "LOGICAL", "CUSTOM",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvoiceUpdateParamsLineItemsCreateUserUnion struct {
	OfInvoiceUpdatesLineItemsCreateUserID         *InvoiceUpdateParamsLineItemsCreateUserID         `json:",omitzero,inline"`
	OfInvoiceUpdatesLineItemsCreateUserExternalID *InvoiceUpdateParamsLineItemsCreateUserExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u InvoiceUpdateParamsLineItemsCreateUserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInvoiceUpdatesLineItemsCreateUserID, u.OfInvoiceUpdatesLineItemsCreateUserExternalID)
}
func (u *InvoiceUpdateParamsLineItemsCreateUserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type InvoiceUpdateParamsLineItemsCreateUserID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsCreateUserID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsCreateUserID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsCreateUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type InvoiceUpdateParamsLineItemsCreateUserExternalID struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsCreateUserExternalID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsCreateUserExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsCreateUserExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price breakdown. Provide amount, or unit_price and quantity, or all three.
type InvoiceUpdateParamsLineItemsCreatePrice struct {
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	// Required if unit_price and quantity are not provided.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Number of units for the line item.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Price per unit as a string in the smallest currency unit, such as cents for USD.
	UnitPrice param.Opt[string] `json:"unit_price,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsCreatePrice) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsCreatePrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsCreatePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsLineItemsCreateTag struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsCreateTag) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsCreateTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsCreateTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type InvoiceUpdateParamsLineItemsDelete struct {
	// Unique identifier for the line item to delete.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsDelete) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsDelete
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data for updating a line item.
//
// The property ID is required.
type InvoiceUpdateParamsLineItemsUpdate struct {
	// Unique identifier for the line item to update.
	ID          string                                  `json:"id" api:"required"`
	Description param.Opt[string]                       `json:"description,omitzero"`
	Price       InvoiceUpdateParamsLineItemsUpdatePrice `json:"price,omitzero"`
	// Tag updates.
	Tags InvoiceUpdateParamsLineItemsUpdateTags `json:"tags,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Quantity, UnitPrice are required.
type InvoiceUpdateParamsLineItemsUpdatePrice struct {
	// Number of units for the line item.
	Quantity int64 `json:"quantity" api:"required"`
	// Price per unit as a string in the smallest currency unit, such as cents for USD.
	UnitPrice string `json:"unit_price" api:"required"`
	// Total amount as a string in the smallest currency unit, such as cents for USD.
	Amount param.Opt[string] `json:"amount,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdatePrice) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdatePrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdatePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tag updates.
type InvoiceUpdateParamsLineItemsUpdateTags struct {
	// Tags to create. The tag key must not already exist.
	Create []InvoiceUpdateParamsLineItemsUpdateTagsCreate `json:"create,omitzero"`
	// Tags to remove.
	Delete []InvoiceUpdateParamsLineItemsUpdateTagsDelete `json:"delete,omitzero"`
	// Tags to set. Creates a new tag or updates an existing tag.
	Set []InvoiceUpdateParamsLineItemsUpdateTagsSet `json:"set,omitzero"`
	// Tags to update. The tag key must already exist.
	Update []InvoiceUpdateParamsLineItemsUpdateTagsUpdate `json:"update,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdateTags) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdateTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdateTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsLineItemsUpdateTagsCreate struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdateTagsCreate) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdateTagsCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdateTagsCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Key is required.
type InvoiceUpdateParamsLineItemsUpdateTagsDelete struct {
	// Tag key to delete.
	Key string `json:"key" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdateTagsDelete) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdateTagsDelete
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdateTagsDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsLineItemsUpdateTagsSet struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdateTagsSet) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdateTagsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdateTagsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsLineItemsUpdateTagsUpdate struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsLineItemsUpdateTagsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsLineItemsUpdateTagsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsLineItemsUpdateTagsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tag updates.
type InvoiceUpdateParamsTags struct {
	// Tags to create. The tag key must not already exist.
	Create []InvoiceUpdateParamsTagsCreate `json:"create,omitzero"`
	// Tags to remove.
	Delete []InvoiceUpdateParamsTagsDelete `json:"delete,omitzero"`
	// Tags to set. Creates a new tag or updates an existing tag.
	Set []InvoiceUpdateParamsTagsSet `json:"set,omitzero"`
	// Tags to update. The tag key must already exist.
	Update []InvoiceUpdateParamsTagsUpdate `json:"update,omitzero"`
	paramObj
}

func (r InvoiceUpdateParamsTags) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsTagsCreate struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsTagsCreate) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsTagsCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsTagsCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Key is required.
type InvoiceUpdateParamsTagsDelete struct {
	// Tag key to delete.
	Key string `json:"key" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsTagsDelete) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsTagsDelete
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsTagsDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsTagsSet struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsTagsSet) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsTagsSet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsTagsSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type InvoiceUpdateParamsTagsUpdate struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceUpdateParamsTagsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceUpdateParamsTagsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceUpdateParamsTagsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceSearchParams struct {
	// Filter criteria for the search.
	Filter InvoiceSearchParamsFilter `json:"filter,omitzero" api:"required"`
	// Pagination parameters.
	PageInfo InvoiceSearchParamsPageInfo `json:"page_info,omitzero"`
	paramObj
}

func (r InvoiceSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter criteria for the search.
type InvoiceSearchParamsFilter struct {
	// Filter by invoice status. `open` returns invoices with non-zero clearing account
	// balances.
	//
	// Any of "open".
	Status string `json:"status,omitzero"`
	// Tag-based filter criteria. When both `any` and `all` are provided, results must
	// match every entry in `all` AND at least one entry in `any`.
	Tags InvoiceSearchParamsFilterTags `json:"tags,omitzero"`
	// Filter invoices by tags on transactions allocated to them. Returns invoices that
	// have at least one allocated transaction matching the specified tags.
	TransactionTags InvoiceSearchParamsFilterTransactionTags `json:"transaction_tags,omitzero"`
	// Line item user filter criteria. When both `any` and `all` are provided, results
	// must match every entry in `all` AND at least one entry in `any`.
	Users InvoiceSearchParamsFilterUsers `json:"users,omitzero"`
	paramObj
}

func (r InvoiceSearchParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InvoiceSearchParamsFilter](
		"status", "open",
	)
}

// Tag-based filter criteria. When both `any` and `all` are provided, results must
// match every entry in `all` AND at least one entry in `any`.
type InvoiceSearchParamsFilterTags struct {
	// Returns invoices matching every specified tag, using AND logic.
	All []InvoiceSearchParamsFilterTagsAll `json:"all,omitzero"`
	// Returns invoices matching at least one of the specified tags, using OR logic.
	Any []InvoiceSearchParamsFilterTagsAny `json:"any,omitzero"`
	paramObj
}

func (r InvoiceSearchParamsFilterTags) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag filter.
//
// The properties Key, Value are required.
type InvoiceSearchParamsFilterTagsAll struct {
	// Tag key to filter on. Must be an exact match; wildcards are not supported.
	Key string `json:"key" api:"required"`
	// Tag value pattern to filter on. Supports wildcards: `*` matches any characters,
	// `?` matches a single character. Use `\*` or `\?` to match literal asterisks or
	// question marks. Use `*` to match any value for the given key.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterTagsAll) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterTagsAll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterTagsAll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag filter.
//
// The properties Key, Value are required.
type InvoiceSearchParamsFilterTagsAny struct {
	// Tag key to filter on. Must be an exact match; wildcards are not supported.
	Key string `json:"key" api:"required"`
	// Tag value pattern to filter on. Supports wildcards: `*` matches any characters,
	// `?` matches a single character. Use `\*` or `\?` to match literal asterisks or
	// question marks. Use `*` to match any value for the given key.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterTagsAny) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterTagsAny
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterTagsAny) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter invoices by tags on transactions allocated to them. Returns invoices that
// have at least one allocated transaction matching the specified tags.
type InvoiceSearchParamsFilterTransactionTags struct {
	// Returns transactions matching every specified tag, using AND logic.
	All []InvoiceSearchParamsFilterTransactionTagsAll `json:"all,omitzero"`
	// Returns transactions matching at least one of the specified tags, using OR
	// logic.
	Any []InvoiceSearchParamsFilterTransactionTagsAny `json:"any,omitzero"`
	paramObj
}

func (r InvoiceSearchParamsFilterTransactionTags) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterTransactionTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterTransactionTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag filter.
//
// The properties Key, Value are required.
type InvoiceSearchParamsFilterTransactionTagsAll struct {
	// Tag key to filter on. Must be an exact match; wildcards are not supported.
	Key string `json:"key" api:"required"`
	// Tag value pattern to filter on. Supports wildcards: `*` matches any characters,
	// `?` matches a single character. Use `\*` or `\?` to match literal asterisks or
	// question marks. Use `*` to match any value for the given key.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterTransactionTagsAll) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterTransactionTagsAll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterTransactionTagsAll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag filter.
//
// The properties Key, Value are required.
type InvoiceSearchParamsFilterTransactionTagsAny struct {
	// Tag key to filter on. Must be an exact match; wildcards are not supported.
	Key string `json:"key" api:"required"`
	// Tag value pattern to filter on. Supports wildcards: `*` matches any characters,
	// `?` matches a single character. Use `\*` or `\?` to match literal asterisks or
	// question marks. Use `*` to match any value for the given key.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterTransactionTagsAny) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterTransactionTagsAny
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterTransactionTagsAny) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Line item user filter criteria. When both `any` and `all` are provided, results
// must match every entry in `all` AND at least one entry in `any`.
type InvoiceSearchParamsFilterUsers struct {
	// Returns invoices matching every specified line item user, using AND logic.
	All []InvoiceSearchParamsFilterUsersAllUnion `json:"all,omitzero"`
	// Returns invoices matching at least one of the specified line item users, using
	// OR logic.
	Any []InvoiceSearchParamsFilterUsersAnyUnion `json:"any,omitzero"`
	paramObj
}

func (r InvoiceSearchParamsFilterUsers) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterUsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterUsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvoiceSearchParamsFilterUsersAllUnion struct {
	OfInvoiceSearchsFilterUsersAllID         *InvoiceSearchParamsFilterUsersAllID         `json:",omitzero,inline"`
	OfInvoiceSearchsFilterUsersAllExternalID *InvoiceSearchParamsFilterUsersAllExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u InvoiceSearchParamsFilterUsersAllUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInvoiceSearchsFilterUsersAllID, u.OfInvoiceSearchsFilterUsersAllExternalID)
}
func (u *InvoiceSearchParamsFilterUsersAllUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type InvoiceSearchParamsFilterUsersAllID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterUsersAllID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterUsersAllID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterUsersAllID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type InvoiceSearchParamsFilterUsersAllExternalID struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterUsersAllExternalID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterUsersAllExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterUsersAllExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvoiceSearchParamsFilterUsersAnyUnion struct {
	OfInvoiceSearchsFilterUsersAnyID         *InvoiceSearchParamsFilterUsersAnyID         `json:",omitzero,inline"`
	OfInvoiceSearchsFilterUsersAnyExternalID *InvoiceSearchParamsFilterUsersAnyExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u InvoiceSearchParamsFilterUsersAnyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInvoiceSearchsFilterUsersAnyID, u.OfInvoiceSearchsFilterUsersAnyExternalID)
}
func (u *InvoiceSearchParamsFilterUsersAnyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type InvoiceSearchParamsFilterUsersAnyID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterUsersAnyID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterUsersAnyID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterUsersAnyID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type InvoiceSearchParamsFilterUsersAnyExternalID struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r InvoiceSearchParamsFilterUsersAnyExternalID) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsFilterUsersAnyExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsFilterUsersAnyExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination parameters.
type InvoiceSearchParamsPageInfo struct {
	// Cursor for fetching the next page of results.
	After param.Opt[string] `json:"after,omitzero"`
	// Number of results to return. Defaults to 20.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	paramObj
}

func (r InvoiceSearchParamsPageInfo) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceSearchParamsPageInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceSearchParamsPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
