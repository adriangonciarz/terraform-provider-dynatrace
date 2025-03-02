/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package vault

import (
	"context"
	"reflect"
	"time"

	"github.com/dtcookie/dynatrace/api/config/credentials/vault"
	monitorsApi "github.com/dtcookie/dynatrace/api/config/synthetic/monitors"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/terraform"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/hcl2sdk"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/logging"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/resources/synthetic/monitors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource produces terraform resource definition for Management Zones
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        hcl2sdk.Convert(new(vault.Credentials).Schema()),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func Resolve(d *schema.ResourceData) (*vault.Credentials, error) {
	resolver, err := terraform.NewResolver(d)
	if err != nil {
		return nil, err
	}
	untypedConfig, err := resolver.Resolve(reflect.TypeOf(vault.Credentials{}))
	if err != nil {
		return nil, err
	}
	typedConfig := untypedConfig.(vault.Credentials)
	return &typedConfig, nil
}

func NewService(m interface{}) *vault.ServiceClient {
	conf := m.(*config.ProviderConfiguration)
	apiService := vault.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials := new(vault.Credentials)
	if err := credentials.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	credentials.ID = nil
	objStub, err := NewService(m).Create(credentials)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(objStub.ID)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials := new(vault.Credentials)
	if err := credentials.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	credentials.ID = opt.NewString(d.Id())
	if err := NewService(m).Update(credentials); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	credentials, err := NewService(m).Get(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	marshalled, err := credentials.MarshalHCL()
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range marshalled {
		d.Set(k, v)
	}
	return diag.Diagnostics{}
}

// Delete the configuration
func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var err error
	credService := NewService(m)
	var config *vault.Credentials
	if config, err = credService.Get(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	if config.ExternalVault != nil && len(config.CredentialUsageSummary) == 1 && config.CredentialUsageSummary[0].MonitorType == vault.MonitorTypes.HTTPMonitor && config.CredentialUsageSummary[0].Count == 1 {
		synService := monitors.NewService(m)
		var monitors *monitorsApi.Monitors
		if monitors, err = synService.ListHTTP(); err != nil {
			return diag.FromErr(err)
		}

		var compare string
		if config.ExternalVault.ClientSecret != nil || config.ExternalVault.ClientID != nil || config.ExternalVault.TenantID != nil {
			compare = "Monitor synchronizing credentials with Azure Key Vault (" + d.Id() + ")"
		} else if config.ExternalVault.RoleID != nil || config.ExternalVault.Certificate != nil {
			compare = "Monitor synchronizing credentials with HashiCorp Vault (" + d.Id() + ")"
		}

		for _, monitor := range monitors.Monitors {
			if monitor.Name == compare {
				// log.Println("Deleting: ", monitor.Name)
				synService.Delete(monitor.EntityID)
				for i := 0; i < 40; i++ {
					if err := NewService(m).Delete(d.Id()); err == nil {
						return diag.Diagnostics{}
					}
					time.Sleep(time.Second * 2)
				}
				return diag.FromErr(err)
			}
		}

	}

	if err := NewService(m).Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}
