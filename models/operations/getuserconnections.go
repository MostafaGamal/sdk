// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/MostafaGamal/sdk/internal/utils"
	"github.com/MostafaGamal/sdk/models/components"
)

type GetUserConnectionsRequest struct {
	Limit  *int64  `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// Key to Sort by followed by order (- for descending)
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	// Filter by resources having fields starting with this parameter value
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (g GetUserConnectionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetUserConnectionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetUserConnectionsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUserConnectionsRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUserConnectionsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetUserConnectionsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetUserConnectionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of the user's created connections (accounts)
	AisConnectionsResponse *components.AisConnectionsResponse
}

func (o *GetUserConnectionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserConnectionsResponse) GetAisConnectionsResponse() *components.AisConnectionsResponse {
	if o == nil {
		return nil
	}
	return o.AisConnectionsResponse
}
