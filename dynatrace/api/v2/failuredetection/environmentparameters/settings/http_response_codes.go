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
			Description: "HTTP response codes which indicate an error on the server side, as number, range or multiple values. Eg. '500', '500-502', '500, 503-504'",
		},
		"fail_on_missing_response_code_server_side": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Treat missing HTTP response code as server error",
		},
		"client_side_errors": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "HTTP response codes which indicate an error on the client side, as number, range or multiple values. Eg. '400', '400-402', '400, 403-404'",
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
		"server_side_errors":                        me.ServerSideErrors,
		"fail_on_missing_response_code_server_side": me.FailOnMissingReponseCodeServerSide,
		"client_side_errors":                        me.ClientSideErrors,
		"fail_on_missing_response_code_client_side": me.FailOnMissingReponseCodeClientSide,
	})
}

func (me *HTTPResponseCodes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"server_side_errors":                        me.ServerSideErrors,
		"fail_on_missing_response_code_server_side": me.FailOnMissingReponseCodeServerSide,
		"client_side_errors":                        me.ClientSideErrors,
		"fail_on_missing_response_code_client_side": me.FailOnMissingReponseCodeClientSide,
	})
}
