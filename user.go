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

// User management operations
//
// UserService contains methods and other services that help with interacting with
// the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

// Creates a user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *UserNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists all users.
func (r *UserService) List(ctx context.Context, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// User object.
type User struct {
	// FRAGMENT generated unique ID.
	ID string `json:"id" api:"required"`
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Name of the user's role. Deprecated, use tags instead.
	//
	// Deprecated: deprecated
	Role string `json:"role" api:"required"`
	// Tags for the user.
	Tags []UserTag `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		Role        respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair.
type UserTag struct {
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
func (r UserTag) RawJSON() string { return r.JSON.raw }
func (r *UserTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewResponse struct {
	// User object.
	Data User `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of users.
type UserListResponse struct {
	// List of users.
	Data []User `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	// User-provided unique ID.
	ExternalID string `json:"external_id" api:"required"`
	// Name of the role to assign. Deprecated, use tags instead.
	Role param.Opt[string] `json:"role,omitzero"`
	// Tags for the user.
	Tags []UserNewParamsTag `json:"tags,omitzero"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key-value tag pair for metadata.
//
// The properties Key, Value are required.
type UserNewParamsTag struct {
	// Tag key. Must not contain #, /, or :. Max 50 characters.
	Key string `json:"key" api:"required"`
	// Tag value. Must not contain #, /, or :. Max 200 characters.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r UserNewParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
