// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-app/internal/sdk/models/shared"
	"net/http"
)

type GetInstalledAppRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetInstalledAppRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type GetInstalledAppResponse struct {
	// Details about installed app.
	App *shared.App
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetInstalledAppResponse) GetApp() *shared.App {
	if o == nil {
		return nil
	}
	return o.App
}

func (o *GetInstalledAppResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInstalledAppResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInstalledAppResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
