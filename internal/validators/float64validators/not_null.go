// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package float64validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Float64 = Float64NotNullValidator{}

// Float64NotNullValidator validates that an attribute is not null. Most
// attributes should set Required: true instead, however in certain scenarios,
// such as a computed nested attribute, all underlying attributes must also be
// computed for planning to not show unexpected differences.
type Float64NotNullValidator struct{}

// Description describes the validation in plain text formatting.
func (v Float64NotNullValidator) Description(_ context.Context) string {
	return "value must be configured"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v Float64NotNullValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v Float64NotNullValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	if !req.ConfigValue.IsNull() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing Attribute Value",
		req.Path.String()+": "+v.Description(ctx),
	)
}

// NotNull returns an validator which ensures that the attribute is
// configured. Most attributes should set Required: true instead, however in
// certain scenarios, such as a computed nested attribute, all underlying
// attributes must also be computed for planning to not show unexpected
// differences.
func NotNull() validator.Float64 {
	return Float64NotNullValidator{}
}