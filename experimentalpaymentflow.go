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

// Payment flow operations
//
// ExperimentalPaymentFlowService contains methods and other services that help
// with interacting with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalPaymentFlowService] method instead.
type ExperimentalPaymentFlowService struct {
	options []option.RequestOption
}

// NewExperimentalPaymentFlowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentalPaymentFlowService(opts ...option.RequestOption) (r ExperimentalPaymentFlowService) {
	r = ExperimentalPaymentFlowService{}
	r.options = opts
	return
}

// Creates a new payment flow.
func (r *ExperimentalPaymentFlowService) New(ctx context.Context, body ExperimentalPaymentFlowNewParams, opts ...option.RequestOption) (res *ExperimentalPaymentFlowNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "payment-flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets a payment flow by ID or external ID.
func (r *ExperimentalPaymentFlowService) Get(ctx context.Context, paymentFlowRef string, opts ...option.RequestOption) (res *ExperimentalPaymentFlowGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if paymentFlowRef == "" {
		err = errors.New("missing required payment_flow_ref parameter")
		return nil, err
	}
	path := fmt.Sprintf("payment-flows/%s", url.PathEscape(paymentFlowRef))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Searches payment flows.
func (r *ExperimentalPaymentFlowService) Search(ctx context.Context, body ExperimentalPaymentFlowSearchParams, opts ...option.RequestOption) (res *ExperimentalPaymentFlowSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "payment-flows/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Payment flow object.
type PaymentFlow struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Timestamp when the payment flow was created.
	Created string `json:"created" api:"required"`
	// User-provided unique external ID.
	ExternalID string `json:"external_id" api:"required"`
	// Invoice being settled.
	Invoice PaymentFlowInvoice `json:"invoice" api:"required"`
	// Timestamp when the payment flow was last modified.
	Modified string `json:"modified" api:"required"`
	// Payment plan for UI rendering.
	PaymentPlan PaymentFlowPaymentPlan `json:"payment_plan" api:"required"`
	// Status of the payment flow.
	Status string `json:"status" api:"required"`
	// Type of payment flow.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		ExternalID  respjson.Field
		Invoice     respjson.Field
		Modified    respjson.Field
		PaymentPlan respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentFlow) RawJSON() string { return r.JSON.raw }
func (r *PaymentFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invoice being settled.
type PaymentFlowInvoice struct {
	// Invoice identifier.
	ID string `json:"id" api:"required"`
	// Invoice external ID.
	ExternalID string `json:"external_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentFlowInvoice) RawJSON() string { return r.JSON.raw }
func (r *PaymentFlowInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment plan for UI rendering.
type PaymentFlowPaymentPlan struct {
	// Payment batches.
	Batches []PaymentFlowPaymentPlanBatch `json:"batches" api:"required"`
	// When the plan was generated.
	GeneratedAt string `json:"generated_at" api:"required"`
	// Invoice identifier.
	InvoiceID string `json:"invoice_id" api:"required"`
	// Plan version.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Batches     respjson.Field
		GeneratedAt respjson.Field
		InvoiceID   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentFlowPaymentPlan) RawJSON() string { return r.JSON.raw }
func (r *PaymentFlowPaymentPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentFlowPaymentPlanBatch struct {
	// Batch identifier.
	BatchID string `json:"batch_id" api:"required"`
	// Batches this one depends on.
	DependsOn []string `json:"depends_on" api:"required"`
	// Human-readable batch label.
	Label string `json:"label" api:"required"`
	// Payments in this batch.
	Payments []PaymentFlowPaymentPlanBatchPayment `json:"payments" api:"required"`
	// Batch status.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchID     respjson.Field
		DependsOn   respjson.Field
		Label       respjson.Field
		Payments    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentFlowPaymentPlanBatch) RawJSON() string { return r.JSON.raw }
func (r *PaymentFlowPaymentPlanBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentFlowPaymentPlanBatchPayment struct {
	// Amount in smallest currency unit.
	Amount string `json:"amount" api:"required"`
	// Currency code.
	Currency string `json:"currency" api:"required"`
	// Direction of the payment.
	Direction string `json:"direction" api:"required"`
	// FRAGMENT generated unique ID.
	PaymentID string `json:"payment_id" api:"required"`
	// Status of the payment.
	Status string `json:"status" api:"required"`
	// User associated with the payment.
	User PaymentFlowPaymentPlanBatchPaymentUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		Direction   respjson.Field
		PaymentID   respjson.Field
		Status      respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentFlowPaymentPlanBatchPayment) RawJSON() string { return r.JSON.raw }
func (r *PaymentFlowPaymentPlanBatchPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User associated with the payment.
type PaymentFlowPaymentPlanBatchPaymentUser struct {
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
func (r PaymentFlowPaymentPlanBatchPaymentUser) RawJSON() string { return r.JSON.raw }
func (r *PaymentFlowPaymentPlanBatchPaymentUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentalPaymentFlowNewResponse struct {
	// Payment flow object.
	Data PaymentFlow `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentalPaymentFlowNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentalPaymentFlowNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentalPaymentFlowGetResponse struct {
	// Payment flow object.
	Data PaymentFlow `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentalPaymentFlowGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentalPaymentFlowGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of payment flows.
type ExperimentalPaymentFlowSearchResponse struct {
	Data []PaymentFlow `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentalPaymentFlowSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentalPaymentFlowSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentalPaymentFlowNewParams struct {
	// User-provided unique external ID.
	ExternalID string `json:"external_id" api:"required"`
	// Invoice to settle.
	Invoice ExperimentalPaymentFlowNewParamsInvoiceUnion `json:"invoice,omitzero" api:"required"`
	// Type of payment flow.
	//
	// Any of "single_invoice_settlement".
	Type ExperimentalPaymentFlowNewParamsType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ExperimentalPaymentFlowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentFlowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentFlowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExperimentalPaymentFlowNewParamsInvoiceUnion struct {
	OfExperimentalPaymentFlowNewsInvoiceID         *ExperimentalPaymentFlowNewParamsInvoiceID         `json:",omitzero,inline"`
	OfExperimentalPaymentFlowNewsInvoiceExternalID *ExperimentalPaymentFlowNewParamsInvoiceExternalID `json:",omitzero,inline"`
	paramUnion
}

func (u ExperimentalPaymentFlowNewParamsInvoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExperimentalPaymentFlowNewsInvoiceID, u.OfExperimentalPaymentFlowNewsInvoiceExternalID)
}
func (u *ExperimentalPaymentFlowNewParamsInvoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type ExperimentalPaymentFlowNewParamsInvoiceID struct {
	// Fragment invoice ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r ExperimentalPaymentFlowNewParamsInvoiceID) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentFlowNewParamsInvoiceID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentFlowNewParamsInvoiceID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExternalID is required.
type ExperimentalPaymentFlowNewParamsInvoiceExternalID struct {
	// Invoice external ID.
	ExternalID string `json:"external_id" api:"required"`
	paramObj
}

func (r ExperimentalPaymentFlowNewParamsInvoiceExternalID) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentFlowNewParamsInvoiceExternalID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentFlowNewParamsInvoiceExternalID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of payment flow.
type ExperimentalPaymentFlowNewParamsType string

const (
	ExperimentalPaymentFlowNewParamsTypeSingleInvoiceSettlement ExperimentalPaymentFlowNewParamsType = "single_invoice_settlement"
)

type ExperimentalPaymentFlowSearchParams struct {
	// Filter by invoice ID.
	InvoiceID param.Opt[string] `json:"invoice_id,omitzero"`
	// Pagination parameters.
	PageInfo ExperimentalPaymentFlowSearchParamsPageInfo `json:"page_info,omitzero"`
	paramObj
}

func (r ExperimentalPaymentFlowSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentFlowSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentFlowSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination parameters.
type ExperimentalPaymentFlowSearchParamsPageInfo struct {
	// Pagination cursor.
	After param.Opt[string] `json:"after,omitzero"`
	// Maximum number of results.
	Limit param.Opt[float64] `json:"limit,omitzero"`
	paramObj
}

func (r ExperimentalPaymentFlowSearchParamsPageInfo) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentalPaymentFlowSearchParamsPageInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentalPaymentFlowSearchParamsPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
