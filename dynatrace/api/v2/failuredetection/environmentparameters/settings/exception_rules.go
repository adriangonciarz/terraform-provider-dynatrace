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
	return map[string]*schema.Schema{
		"ignore_all_exceptions": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "",
		},
		"custom_error_rules": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "",
		},
		"ignore_span_failure_detection": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "",
		},
	}
}

func (me *ExceptionRules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"ignore_all_exceptions":         &me.IgnoreAllExceptions,
		"custom_error_rules":            &me.CustomErrorRules,
		"ignore_span_failure_detection": &me.IgnoreSpanFailureDetection,
	})
}

func (me *ExceptionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"ignore_all_exceptions":         &me.IgnoreAllExceptions,
		"custom_error_rules":            &me.CustomErrorRules,
		"ignore_span_failure_detection": &me.IgnoreSpanFailureDetection,
	})
}
