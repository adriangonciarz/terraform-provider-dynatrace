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
	return map[string]*schema.Schema{}
}

func (me *CustomErrorRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{})
}

func (me *CustomErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{})
}
