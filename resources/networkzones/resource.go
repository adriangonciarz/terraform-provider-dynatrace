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

package networkzones

import (
	"context"

	"github.com/dtcookie/dynatrace/api/config/networkzones"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/hcl"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/hcl2sdk"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/logging"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource produces terraform resource definition for Management Zones
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        hcl2sdk.Convert(new(networkzones.NetworkZone).Schema()),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func NewService(m interface{}) *networkzones.ServiceClient {
	conf := m.(*config.ProviderConfiguration)
	apiService := networkzones.NewService(conf.DTApiV2URL, conf.APIToken)
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if config.HTTPVerbose {
		rest.Verbose = true
	}
	config := new(networkzones.NetworkZone)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	if config.Description != nil && *config.Description == "The default network zone. This is the network zone for OneAgents or ActiveGates that do not have any network zone configured." {
		d.SetId("default")
		return Read(ctx, d, m)
	}
	abc := NewService(m)
	nz, err := abc.Create(config)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*nz.ID)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if config.HTTPVerbose {
		rest.Verbose = true
	}
	config := new(networkzones.NetworkZone)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	if err := NewService(m).Update(d.Id(), config); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, err := NewService(m).Get(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	marshalled, err := config.MarshalHCL()
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
	if d.Id() != "default" {
		if err := NewService(m).Delete(d.Id()); err != nil {
			return diag.FromErr(err)
		}
	}
	return diag.Diagnostics{}
}
