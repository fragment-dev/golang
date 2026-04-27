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
	"github.com/fragment-dev/fragment-billing-go/internal/requestconfig"
	"github.com/fragment-dev/fragment-billing-go/option"
	"github.com/fragment-dev/fragment-billing-go/packages/param"
	"github.com/fragment-dev/fragment-billing-go/packages/respjson"
)

// Product management operations
//
// ProductService contains methods and other services that help with interacting
// with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	options []option.RequestOption
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r ProductService) {
	r = ProductService{}
	r.options = opts
	return
}

// Creates a product.
func (r *ProductService) New(ctx context.Context, body ProductNewParams, opts ...option.RequestOption) (res *ProductNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a product by code.
func (r *ProductService) Get(ctx context.Context, code string, opts ...option.RequestOption) (res *ProductGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all products.
func (r *ProductService) List(ctx context.Context, opts ...option.RequestOption) (res *ProductListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Product object.
type Product struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Product code.
	Code string `json:"code" api:"required"`
	// Timestamp when the product was created. Uses ISO 8601 format.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// Roles that can pay for the product.
	PaidByRoles []ProductPaidByRole `json:"paid_by_roles" api:"required"`
	// Roles that can receive payment for the product.
	PaidToRoles []ProductPaidToRole `json:"paid_to_roles" api:"required"`
	// Current version of the product.
	UpdateVersion float64 `json:"update_version" api:"required"`
	// Workspace ID of the product.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Product description.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Code          respjson.Field
		Created       respjson.Field
		PaidByRoles   respjson.Field
		PaidToRoles   respjson.Field
		UpdateVersion respjson.Field
		WorkspaceID   respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Product) RawJSON() string { return r.JSON.raw }
func (r *Product) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role reference in product API responses.
type ProductPaidByRole struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Name of the role.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductPaidByRole) RawJSON() string { return r.JSON.raw }
func (r *ProductPaidByRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role reference in product API responses.
type ProductPaidToRole struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// Name of the role.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductPaidToRole) RawJSON() string { return r.JSON.raw }
func (r *ProductPaidToRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductNewResponse struct {
	// Product object.
	Data Product `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductGetResponse struct {
	// Product object.
	Data Product `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListResponse struct {
	// List of products.
	Data []Product `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductNewParams struct {
	// Unique product code.
	Code string `json:"code" api:"required"`
	// Product description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Roles that can pay for the product. Reference roles by `id` or `name`. At least
	// one of `paid_by_roles` or `paid_to_roles` must be provided.
	PaidByRoles []ProductNewParamsPaidByRoleUnion `json:"paid_by_roles,omitzero"`
	// Roles that can receive payment for the product. Reference roles by `id` or
	// `name`. At least one of `paid_by_roles` or `paid_to_roles` must be provided.
	PaidToRoles []ProductNewParamsPaidToRoleUnion `json:"paid_to_roles,omitzero"`
	paramObj
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProductNewParamsPaidByRoleUnion struct {
	OfProductNewsPaidByRoleID   *ProductNewParamsPaidByRoleID   `json:",omitzero,inline"`
	OfProductNewsPaidByRoleName *ProductNewParamsPaidByRoleName `json:",omitzero,inline"`
	paramUnion
}

func (u ProductNewParamsPaidByRoleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProductNewsPaidByRoleID, u.OfProductNewsPaidByRoleName)
}
func (u *ProductNewParamsPaidByRoleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type ProductNewParamsPaidByRoleID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r ProductNewParamsPaidByRoleID) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParamsPaidByRoleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParamsPaidByRoleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type ProductNewParamsPaidByRoleName struct {
	// Name of the role.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ProductNewParamsPaidByRoleName) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParamsPaidByRoleName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParamsPaidByRoleName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProductNewParamsPaidToRoleUnion struct {
	OfProductNewsPaidToRoleID   *ProductNewParamsPaidToRoleID   `json:",omitzero,inline"`
	OfProductNewsPaidToRoleName *ProductNewParamsPaidToRoleName `json:",omitzero,inline"`
	paramUnion
}

func (u ProductNewParamsPaidToRoleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProductNewsPaidToRoleID, u.OfProductNewsPaidToRoleName)
}
func (u *ProductNewParamsPaidToRoleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ID is required.
type ProductNewParamsPaidToRoleID struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r ProductNewParamsPaidToRoleID) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParamsPaidToRoleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParamsPaidToRoleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type ProductNewParamsPaidToRoleName struct {
	// Name of the role.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ProductNewParamsPaidToRoleName) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParamsPaidToRoleName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParamsPaidToRoleName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
