package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CustomErrorRule struct {
	RequestAttribute string                   `json:"requestAttribute"`
	Condition        CustomErrorRuleCondition `json:"condition"`
}

func (me CustomErrorRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"request_attribute": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ID of request attribute to apply custom error rule",
		},
		"condition": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "Definition of the rule condition",
		},
	}
}

func (me *CustomErrorRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"request_attributes": &me.RequestAttribute,
		"condition":          &me.Condition,
	})
}

func (me *CustomErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"request_attributes": &me.RequestAttribute,
		"condition":          &me.Condition,
	})
}