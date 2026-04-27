// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fragment

import (
	"github.com/fragment-dev/fragment-billing-go/option"
)

// ExperimentalService contains methods and other services that help with
// interacting with the fragment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalService] method instead.
type ExperimentalService struct {
	options []option.RequestOption
	// Payment flow operations
	PaymentFlows ExperimentalPaymentFlowService
	// Payment operations
	Payments ExperimentalPaymentService
}

// NewExperimentalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExperimentalService(opts ...option.RequestOption) (r ExperimentalService) {
	r = ExperimentalService{}
	r.options = opts
	r.PaymentFlows = NewExperimentalPaymentFlowService(opts...)
	r.Payments = NewExperimentalPaymentService(opts...)
	return
}
