// "serverSideErrors": "500-599",
// "failOnMissingResponseCodeServerSide": true,
// "clientSideErrors": "400-403, 405-498",
// "failOnMissingResponseCodeClientSide": true

package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type HTTPResponseCodes struct {
	ServerSideErrors                   string `json:"serverSideErrors"`
	FailOnMissingReponseCodeServerSide bool   `json:"failOnMissingResponseCodeServerSide"`
	ClientSideErrors                   string `json:"clientSideErrors"`
	FailOnMissingReponseCodeClientSide bool   `json:"failOnMissingResponseCodeClientSide"`
}

func (me *HTTPResponseCodes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server_side_errors": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "HTTP response codes which indicate an error on the server side",
		},
		"fail_on_missing_response_code_server_side": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Treat missing HTTP response code as server error",
		},
		"client_side_errors": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "HTTP response codes which indicate an error on the client side",
		},
		"fail_on_missing_response_code_client_side": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Treat missing HTTP response code as client error",
		},
	}
}

func (me *HTTPResponseCodes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]interface{}{
		"server_side_errors": me.ServerSideErrors,
		"broken_links":       me.BrokenLinkDomains,
	})
}

func (me *HTTPResponseCodes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"http_not_found_failures": me.HTTPNotFoundFailures,
		"broken_links":            me.BrokenLinkDomains,
	})
}
