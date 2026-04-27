// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/fragment-go/internal/apijson"
	"github.com/stainless-sdks/fragment-go/internal/requestconfig"
	"github.com/stainless-sdks/fragment-go/option"
	"github.com/stainless-sdks/fragment-go/packages/param"
	"github.com/stainless-sdks/fragment-go/packages/respjson"
)

// External account management operations
//
// ExternalAccountService contains methods and other services that help with
// interacting with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalAccountService] method instead.
type ExternalAccountService struct {
	options []option.RequestOption
}

// NewExternalAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExternalAccountService(opts ...option.RequestOption) (r ExternalAccountService) {
	r = ExternalAccountService{}
	r.options = opts
	return
}

// Creates an external account.
func (r *ExternalAccountService) New(ctx context.Context, body ExternalAccountNewParams, opts ...option.RequestOption) (res *ExternalAccountNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "external-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists all external accounts.
func (r *ExternalAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *ExternalAccountListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "external-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// External account object.
type ExternalAccount struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Name of the account.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalAccount) RawJSON() string { return r.JSON.raw }
func (r *ExternalAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalAccountNewResponse struct {
	// External account object.
	Data ExternalAccount `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of external accounts.
type ExternalAccountListResponse struct {
	Data []ExternalAccount `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalAccountNewParams struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Name of the account.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
