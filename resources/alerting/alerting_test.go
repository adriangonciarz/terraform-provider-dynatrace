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

package alerting_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	alertingapi "github.com/dtcookie/dynatrace/api/config/alerting"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/testbase"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const ResourceName = "dynatrace_alerting_profile"
const TestDataFolder = "../../test_data/alerting_profiles"
const RequestPath = "%s/alertingProfiles/%s"
const TestCaseID = "alerting_profiles"

type TestStruct struct {
	resourceKey string
}

func (test *TestStruct) Anonymize(m map[string]interface{}) {
	delete(m, "id")
	delete(m, "displayName")
	delete(m, "metadata")
}

func (test *TestStruct) ResourceKey() string {
	return test.resourceKey
}

func (test *TestStruct) CreateTestCase(file string, localJSONFile string, t *testing.T) (*resource.TestCase, error) {
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(file); err != nil {
		return nil, err
	}
	config := string(content)
	name := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	resourceName := test.ResourceKey() + "." + name
	config = strings.ReplaceAll(config, "#name#", name)
	return &resource.TestCase{
		PreCheck:          func() { testbase.TestAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testbase.TestAccProviderFactories,
		CheckDestroy:      test.CheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					CheckExists(resourceName, t),
					testbase.CompareLocalRemote(test, resourceName, localJSONFile, t),
				),
			},
		},
	}, nil
}

func TestAlerting(t *testing.T) {
	if disabled, ok := testbase.DisabledTests[TestCaseID]; ok && disabled {
		t.Skip()
	}
	test := &TestStruct{resourceKey: ResourceName}
	var err error
	var testCase *resource.TestCase
	if testCase, err = test.CreateTestCase(
		TestDataFolder+"/example_a.tf",
		TestDataFolder+"/example_a.json",
		t,
	); err != nil {
		t.Fatal(err)
		return
	}
	resource.Test(t, *testCase)
}

func TestAlerting2(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting3(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting4(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting5(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting6(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting7(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting8(t *testing.T) {
	TestAlerting(t)
}

func TestAlerting9(t *testing.T) {
	TestAlerting(t)
}

func (test *TestStruct) URL(id string) string {
	envURL := testbase.TestAccProvider.Meta().(*config.ProviderConfiguration).DTenvURL
	reqPath := RequestPath
	return fmt.Sprintf(reqPath, envURL, id)
}

func (test *TestStruct) CheckDestroy(s *terraform.State) error {
	providerConf := testbase.TestAccProvider.Meta().(*config.ProviderConfiguration)
	restClient := alertingapi.NewService(providerConf.DTenvURL, providerConf.APIToken)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ResourceName {
			continue
		}

		id := rs.Primary.ID

		if _, err := restClient.Get(id); err != nil {
			// HTTP Response "404 Not Found" signals a success
			if strings.Contains(err.Error(), `"code": 404`) {
				return nil
			}
			// any other error should fail the test
			return err
		}
		return fmt.Errorf("Configuration still exists: %s", rs.Primary.ID)
	}

	return nil
}

func CheckExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testbase.TestAccProvider.Meta().(*config.ProviderConfiguration)
		restClient := alertingapi.NewService(providerConf.DTenvURL, providerConf.APIToken)

		if rs, ok := s.RootModule().Resources[n]; ok {
			if _, err := restClient.Get(rs.Primary.ID); err != nil {
				return err
			}
			return nil
		}

		return fmt.Errorf("Not found: %s", n)
	}
}
