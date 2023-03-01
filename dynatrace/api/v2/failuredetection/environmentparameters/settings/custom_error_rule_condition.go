package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CustomErrorRuleCondition struct {
	CompareOperationType string `json:"compareOperationType"`
	TextValue            string `json:"textValue"`
	CaseSensitive        bool   `json:"caseSensitive"`
}

func (me CustomErrorRuleCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"compare_operation_type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Type of comparison operation",
		},
		"text_value": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Value to compare to",
		},
		"case_sensitive": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Should the value comparison be case-sensitive",
		},
	}
}

func (me *CustomErrorRuleCondition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"compare_operation_type": &me.CompareOperationType,
		"text_value":             &me.TextValue,
		"case_sensitive":         &me.CaseSensitive,
	})
}

func (me *CustomErrorRuleCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"compare_operation_type": &me.CompareOperationType,
		"text_value":             &me.TextValue,
		"case_sensitive":         &me.CaseSensitive,
	})
}
