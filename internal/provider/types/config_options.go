// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ConfigOptions struct {
	ComponentType types.String             `tfsdk:"component_type"`
	Configuration []ComponentConfiguration `tfsdk:"configuration"`
}
