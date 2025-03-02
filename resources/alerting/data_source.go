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

package alerting

import (
	alertingapi "github.com/dtcookie/dynatrace/api/config/alerting"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceRead,
		Schema: map[string]*schema.Schema{
			"profiles": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func DataSourceRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(uuid.New().String())
	conf := m.(*config.ProviderConfiguration)
	apiService := alertingapi.NewService(conf.DTenvURL, conf.APIToken)
	stubList, err := apiService.List()
	if err != nil {
		return err
	}
	if len(stubList.Values) > 0 {
		profiles := map[string]interface{}{}
		for _, stub := range stubList.Values {
			profiles[stub.Name] = stub.ID
		}
		d.Set("profiles", profiles)
	}
	return nil
}
