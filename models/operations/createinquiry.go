// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/MostafaGamal/sdk/models/components"
)

type CreateInquiryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateInquiryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
