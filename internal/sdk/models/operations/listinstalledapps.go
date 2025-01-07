// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/models/shared"
	"net/http"
)

type ListInstalledAppsRequest struct {
	// Filter apps by specific component type
	ComponentType *shared.ComponentType `queryParam:"style=form,explode=true,name=componentType"`
	// Page number for pagination
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of items per page
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=pageSize"`
}

func (l ListInstalledAppsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListInstalledAppsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListInstalledAppsRequest) GetComponentType() *shared.ComponentType {
	if o == nil {
		return nil
	}
	return o.ComponentType
}

func (o *ListInstalledAppsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListInstalledAppsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type Pagination struct {
	Page     *int64 `json:"page,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	Total    *int64 `json:"total,omitempty"`
}

func (o *Pagination) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *Pagination) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *Pagination) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

// ListInstalledAppsResponseBody - Successful response
type ListInstalledAppsResponseBody struct {
	Apps       []shared.App `json:"apps,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

func (o *ListInstalledAppsResponseBody) GetApps() []shared.App {
	if o == nil {
		return nil
	}
	return o.Apps
}

func (o *ListInstalledAppsResponseBody) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type ListInstalledAppsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	Object *ListInstalledAppsResponseBody
}

func (o *ListInstalledAppsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListInstalledAppsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListInstalledAppsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListInstalledAppsResponse) GetObject() *ListInstalledAppsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}