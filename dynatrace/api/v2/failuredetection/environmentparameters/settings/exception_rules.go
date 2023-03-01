package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExceptionRules struct {
	IgnoreAllExceptions        bool              `json:"ignoreAllExceptions"`
	CustomErrorRules           []CustomErrorRule `json:"customErrorRules"`
	IgnoreSpanFailureDetection bool              `json:"ignoreSpanFailureDetection"`
}

func (me ExceptionRules) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (me *ExceptionRules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{})
}

func (me *ExceptionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{})
}
