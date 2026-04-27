// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment

import (
	"context"
	"net/http"
	"slices"

	"github.com/fragment-dev/fragment-billing-go/internal/apijson"
	"github.com/fragment-dev/fragment-billing-go/internal/requestconfig"
	"github.com/fragment-dev/fragment-billing-go/option"
	"github.com/fragment-dev/fragment-billing-go/packages/param"
	"github.com/fragment-dev/fragment-billing-go/packages/respjson"
)

// Role management operations
//
// RoleService contains methods and other services that help with interacting with
// the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoleService] method instead.
type RoleService struct {
	options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r RoleService) {
	r = RoleService{}
	r.options = opts
	return
}

// Creates a role. Deprecated, use user tags instead.
//
// Deprecated: deprecated
func (r *RoleService) New(ctx context.Context, body RoleNewParams, opts ...option.RequestOption) (res *RoleNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists all roles. Deprecated, use user tags instead.
//
// Deprecated: deprecated
func (r *RoleService) List(ctx context.Context, opts ...option.RequestOption) (res *RoleListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Role object. Deprecated, use user tags instead.
//
// Deprecated: deprecated
type Role struct {
	// FRAGMENT generated unique ID. Deprecated.
	//
	// Deprecated: deprecated
	ID string `json:"id" api:"required"`
	// Name of the role. Deprecated, use user tags instead.
	//
	// Deprecated: deprecated
	Role string `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: deprecated
type RoleNewResponse struct {
	// Role object. Deprecated, use user tags instead.
	//
	// Deprecated: deprecated
	Data Role `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoleNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RoleNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of roles. Deprecated, use user tags instead.
//
// Deprecated: deprecated
type RoleListResponse struct {
	// List of roles. Deprecated, use user tags instead.
	//
	// Deprecated: deprecated
	Data []Role `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoleListResponse) RawJSON() string { return r.JSON.raw }
func (r *RoleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoleNewParams struct {
	// Name of the role. Deprecated, use user tags instead.
	Role string `json:"role" api:"required"`
	paramObj
}

func (r RoleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RoleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
