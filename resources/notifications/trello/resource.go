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

package trello

import (
	"context"

	"github.com/dtcookie/dynatrace/api/config/v2/notifications"
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
		Schema:        hcl2sdk.Convert(new(notifications.Trello).Schema()),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func NewService(m interface{}) *notifications.ServiceClient {
	conf := m.(*config.ProviderConfiguration)
	apiService := notifications.NewService(conf.DTApiV2URL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(notifications.Trello)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	id, err := NewService(m).Create(&notifications.Notification{
		Enabled:   config.Enabled,
		Name:      config.Name,
		ProfileID: config.ProfileID,
		Type:      notifications.Types.Trello,
		Trello:    config,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(notifications.Trello)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	if err := NewService(m).Update(d.Id(), &notifications.Notification{
		Enabled:   config.Enabled,
		Name:      config.Name,
		ProfileID: config.ProfileID,
		Type:      notifications.Types.Trello,
		Trello:    config,
	}); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := notifications.Notification{}
	if err := NewService(m).Get(d.Id(), &config); err != nil {
		return diag.FromErr(err)
	}
	config.Trello.Enabled = config.Enabled
	config.Trello.Name = config.Name
	config.Trello.ProfileID = config.ProfileID
	marshalled, err := config.Trello.MarshalHCL()
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
	if err := NewService(m).Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
