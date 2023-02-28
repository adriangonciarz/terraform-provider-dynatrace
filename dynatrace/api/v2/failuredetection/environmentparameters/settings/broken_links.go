package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BrokenLinks struct {
	HTTPNotFoundFailures bool      `json:"http404NotFoundFailures"`
	BrokenLinkDomains    []*string `json:"brokenLinkDomains,omitempty"`
}

func (me *BrokenLinks) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_not_found_failures": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Consider 404 HTTP response code from service as failures",
		},
		"broken_links": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Dynatrace will consider 404s thrown by hosts at these domains to be service failures related to your application",
		},
	}
}

func (me *BrokenLinks) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]interface{}{
		"http_not_found_failures": me.HTTPNotFoundFailures,
		"broken_links":            me.BrokenLinkDomains,
	})
}

func (me *BrokenLinks) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"http_not_found_failures": me.HTTPNotFoundFailures,
		"broken_links":            me.BrokenLinkDomains,
	})
}
