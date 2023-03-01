package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FailureDetectionEnvironmentParameter struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	HTTPResponseCodes HTTPResponseCodes `json:"httpResponseCodes"`
	BrokenLinks       BrokenLinks       `json:"brokenLinks"`
	ExceptionRules    ExceptionRules    `json:"exceptionRules"`
}

func (me *FailureDetectionEnvironmentParameter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    false,
			Description: "Name of failure detection parameter",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of failure detection parameter",
		},
		"http_response_codes": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "Defines which HTTP response codes should be treated as server/client side errors",
			Elem: &schema.Resource{
				Schema: new(HTTPResponseCodes).Schema(),
			},
		},
		"broken_links": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "Makes requests resulting in 404 HTTP Response Code treated as server-side service failures",
			Elem: &schema.Resource{
				Schema: new(BrokenLinks).Schema(),
			},
		},
		"exception_rules": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "Customize failure detection for specific exceptions and errors",
			Elem: &schema.Resource{
				Schema: new(ExceptionRules).Schema(),
			},
		},
	}
}

func (me *FailureDetectionEnvironmentParameter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"name":                &me.Name,
		"description":         &me.Description,
		"http_response_codes": &me.HTTPResponseCodes,
		"broken_links":        &me.BrokenLinks,
		"exception_rules":     &me.ExceptionRules,
	})
}

func (me *FailureDetectionEnvironmentParameter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":                &me.Name,
		"description":         &me.Description,
		"http_response_codes": &me.HTTPResponseCodes,
		"broken_links":        &me.BrokenLinks,
		"exception_rules":     &me.ExceptionRules,
	})
}
