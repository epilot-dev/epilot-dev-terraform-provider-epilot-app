// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ComponentType - Type of app component
type ComponentType string

const (
	ComponentTypeCustomJourneyBlock ComponentType = "CUSTOM_JOURNEY_BLOCK"
	ComponentTypePortalExtension    ComponentType = "PORTAL_EXTENSION"
)

func (e ComponentType) ToPointer() *ComponentType {
	return &e
}
func (e *ComponentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CUSTOM_JOURNEY_BLOCK":
		fallthrough
	case "PORTAL_EXTENSION":
		*e = ComponentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ComponentType: %v", v)
	}
}
