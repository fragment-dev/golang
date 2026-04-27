// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/fragment-dev/fragment-billing-go/internal/apijson"
	"github.com/fragment-dev/fragment-billing-go/internal/requestconfig"
	"github.com/fragment-dev/fragment-billing-go/option"
	"github.com/fragment-dev/fragment-billing-go/packages/param"
	"github.com/fragment-dev/fragment-billing-go/packages/respjson"
)

// Payment operations
//
// ExperimentalPaymentService contains methods and other services that help with
// interacting with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalPaymentService] method instead.
type ExperimentalPaymentService struct {
	options []option.RequestOption
}

// NewExperimentalPaymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentalPaymentService(opts ...option.RequestOption) (r ExperimentalPaymentService) {
	r = ExperimentalPaymentService{}
	r.options = opts
	return
}

// Gets a payment by ID or external ID.
func (r *ExperimentalPaymentService) Get(ctx context.Context, paymentRef string, opts ...option.RequestOption) (res *ExperimentalPaymentGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paymentRef == "" {
		err = errors.New("missing required payment_ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("payments/%s", url.PathEscape(paymentRef))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Searches payments.
func (r *ExperimentalPaymentService) Search(ctx context.Context, body ExperimentalPaymentSearchParams, opts ...option.RequestOption) (res *ExperimentalPaymentSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "payments/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Payment object.
type Payment struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Amount in smallest currency unit.
	Amount string `json:"amount" api:"required"`
	// Timestamp when the payment was created.
	Created string `json:"created" api:"required"`
	// Currency code.
	Currency string `json:"currency" api:"required"`
	// Direction of the payment.
	Direction string `json:"direction" api:"required"`
	// Timestamp when the payment was last modified.
	Modified string `json:"modified" api:"required"`
	// Payment account ID.
	PaymentAccountID string `json:"payment_account_id" api:"required"`
	// Payment flow ID.
	PaymentFlowID string `json:"payment_flow_id" api:"required"`
	// Status of the payment.
	Status string `json:"status" api:"required"`
	// Associated transaction IDs.
	TransactionIDs []string `json:"transaction_ids" api:"required"`
	// User-provided unique ID when the payment was created with one.
	ExternalID string `json:"external_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Created          respjson.Field
		Currency         respjson.Field
		Direction        respjson.Field
		Modified         respjson.Field
		PaymentAccountID respjson.Field
		PaymentFlowID    respjson.Field
		Status           respjson.Field
		TransactionIDs   respjson.Field
		ExternalID       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Payment) RawJSON() string { return r.JSON.raw }
func (r *Payment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentalPaymentGetResponse struct {
	// Payment object.
	Data Payment `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentalPaymentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentalPaymentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of payments.
type ExperimentalPaymentSearchResponse struct {
	Data []Payment `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentalPaymentSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentalPaymentSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentalPaymentSearchParams struct {
	// Filter by payment flow ID.
	PaymentFlowID param.Opt[string] `json:"payment_flow_id,omitzero"`
	// Pagination parameters.
	PageInfo ExperimentalPaymentSearchParamsPageInfo `json:"page_info,omitzero"`
	paramObj
}

func (r ExperimentalPaymentSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination parameters.
type ExperimentalPaymentSearchParamsPageInfo struct {
	// Pagination cursor.
	After param.Opt[string] `json:"after,omitzero"`
	// Maximum number of results.
	Limit param.Opt[float64] `json:"limit,omitzero"`
	paramObj
}

func (r ExperimentalPaymentSearchParamsPageInfo) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentSearchParamsPageInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentSearchParamsPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
