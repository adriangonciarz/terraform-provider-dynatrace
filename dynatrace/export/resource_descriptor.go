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

package export

import (
	"reflect"
	"strings"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/alerting/connectivityalerts"
	database_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/databases"
	disk_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/infrastructure/disks"
	disk_specific_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/infrastructure/disks/perdiskoverride"
	host_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/infrastructure/hosts"
	custom_app_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/custom"
	custom_app_crash_rate "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/custom/crashrate"
	mobile_app_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/mobile"
	mobile_app_crash_rate "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/mobile/crashrate"
	web_app_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/web"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/declarativegrouping"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/deployment/management/updatewindows"
	hostmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/host/monitoring"
	hostprocessgroupmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/host/processgroups/monitoringstate"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/schemalesslogmetric"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/mainframe/txmonitoring"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/apache"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/dotnet"
	golang "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/go"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/iis"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/java"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/nginx"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/nodejs"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/opentracingnative"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/php"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/varnish"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/wsmb"
	processmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/process/monitoring"
	customprocessmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/process/monitoring/custom"
	processavailability "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processavailability"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/advanceddetectionrule"
	workloaddetection "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/cloudapplication/workloaddetection"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/detectionflags"
	processgroupmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/monitoring/state"
	processgroupsimpledetection "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/simpledetectionrule"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processvisibility"
	rumcustomenablement "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/custom/enablement"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/ipmappings"
	rummobileenablement "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/mobile/enablement"
	rumprocessgroup "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/processgroup"
	rumproviderbreakdown "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/providerbreakdown"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/userexperiencescore"
	rumwebenablement "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/enablement"
	webappresourcecleanup "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/resourcecleanuprules"
	sessionreplaywebprivacy "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/sessionreplay/web/privacypreferences"
	sessionreplayresourcecapture "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/sessionreplay/web/resourcecapturing"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/settings/mutedrequests"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/availability"
	browseroutagehandling "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/browser/outagehandling"
	browserperformancethresholds "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/browser/performancethresholds"
	httpcookies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/http/cookies"
	httpoutagehandling "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/http/outagehandling"
	httpperformancethresholds "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/http/performancethresholds"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/usability/analytics"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/groups"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/users"
	alertingv1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/alerting"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/dashboards"
	maintenancev1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/maintenance"
	managementzonesv1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/managementzones"
	notificationsv1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/notifications"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/requestnaming/order"
	locations "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/synthetic/locations/private"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/alerting"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/anomalies/frequentissues"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/anomalies/metricevents"
	service_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/anomalies/services"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/availability/processgroupalerting"
	ddupool "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ddupool"
	environmentparameters "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/failuredetection/environmentparameters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/filters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/imsbridges"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/queuemanagers"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/queuesharinggroups"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/keyrequests"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/networkzones"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/ansible"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/email"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/jira"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/opsgenie"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/pagerduty"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/servicenow"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/slack"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/trello"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/victorops"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/webhook"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/xmatters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/slo"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/attributes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/capture"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/ctxprop"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/entrypoints"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/resattr"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings/services/cache"

	v2managementzones "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/managementzones"

	application_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/applications"
	database_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/databaseservices"
	disk_event_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/diskevents"
	host_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/hosts"
	custom_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/metricevents"
	pg_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/processgroups"
	service_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/services"

	host_naming "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/naming/hosts"
	processgroup_naming "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/naming/processgroups"
	service_naming "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/naming/services"
	networkzone "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/networkzones"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/mobile"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web/dataprivacy"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web/detection"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web/errors"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/autotags"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/aws"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/azure"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/cloudfoundry"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/kubernetes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/vault"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/customservices"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/dashboards/sharing"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/jsondashboards"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/requestattributes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/requestnaming"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/synthetic/monitors/browser"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/synthetic/monitors/http"

	calculated_service_metrics "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/metrics/calculated/service"
	v2maintenance "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/maintenance"
)

func NewResourceDescriptor[T settings.Settings](fn func(credentials *settings.Credentials) settings.CRUDService[T], dependencies ...Dependency) ResourceDescriptor {
	return ResourceDescriptor{
		Service: func(credentials *settings.Credentials) settings.CRUDService[settings.Settings] {
			return &settings.GenericCRUDService[T]{Service: cache.CRUD(fn(credentials))}
		},
		protoType:    newSettings(fn),
		Dependencies: dependencies,
	}
}

func newSettings[T settings.Settings](sfn func(credentials *settings.Credentials) settings.CRUDService[T]) T {
	var proto T
	return reflect.New(reflect.TypeOf(proto).Elem()).Interface().(T)
}

type ResourceDescriptor struct {
	Dependencies []Dependency
	Service      func(credentials *settings.Credentials) settings.CRUDService[settings.Settings]
	protoType    settings.Settings
	except       func(id string, name string) bool
}

func (me ResourceDescriptor) Specify(t notifications.Type) ResourceDescriptor {
	if notification, ok := me.protoType.(*notifications.Notification); ok {
		notification.Type = t
	}
	return me
}

func (me ResourceDescriptor) Except(except func(id string, name string) bool) ResourceDescriptor {
	me.except = except
	return me
}

func (me ResourceDescriptor) NewSettings() settings.Settings {
	res := reflect.New(reflect.TypeOf(me.protoType).Elem()).Interface().(settings.Settings)
	if notification, ok := res.(*notifications.Notification); ok {
		notification.Type = me.protoType.(*notifications.Notification).Type
	}
	return res
}

var AllResources = map[ResourceType]ResourceDescriptor{
	ResourceTypes.Alerting: NewResourceDescriptor(
		alerting.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.AnsibleTowerNotification: NewResourceDescriptor(
		ansible.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.AnsibleTower),
	ResourceTypes.ApplicationAnomalies: NewResourceDescriptor(application_anomalies.Service),
	ResourceTypes.ApplicationDataPrivacy: NewResourceDescriptor(
		dataprivacy.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.ApplicationDetection: NewResourceDescriptor(
		detection.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.ApplicationErrorRules: NewResourceDescriptor(
		errors.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.AutoTag: NewResourceDescriptor(
		autotags.Service,
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.AWSCredentials: NewResourceDescriptor(
		aws.Service,
		Dependencies.ID(ResourceTypes.AWSCredentials),
	),
	ResourceTypes.AzureCredentials: NewResourceDescriptor(azure.Service),
	ResourceTypes.BrowserMonitor: NewResourceDescriptor(
		browser.Service,
		Dependencies.ID(ResourceTypes.SyntheticLocation),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.Credentials),
	).Except(func(id string, name string) bool {
		return strings.HasPrefix(name, "Monitor synchronizing credentials with")
	}),
	ResourceTypes.CalculatedServiceMetric: NewResourceDescriptor(
		calculated_service_metrics.Service,
		Dependencies.ManagementZone,
		Dependencies.RequestAttribute,
		Dependencies.ManagementZone,
		Dependencies.Service,
		Dependencies.Host,
		Dependencies.HostGroup,
		Dependencies.ProcessGroup,
		Dependencies.ProcessGroupInstance,
	),
	ResourceTypes.CloudFoundryCredentials: NewResourceDescriptor(cloudfoundry.Service),
	ResourceTypes.CustomAnomalies: NewResourceDescriptor(
		custom_anomalies.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	).Except(func(id string, name string) bool {
		return strings.HasPrefix(id, "builtin:") || strings.HasPrefix(id, "ruxit.") || strings.HasPrefix(id, "dynatrace.") || strings.HasPrefix(id, "custom.remote.python.") || strings.HasPrefix(id, "custom.python.")
	}),
	ResourceTypes.CustomAppAnomalies: NewResourceDescriptor(
		custom_app_anomalies.Service,
		Coalesce(Dependencies.DeviceApplicationMethod),
		Coalesce(Dependencies.CustomApplication),
	),
	ResourceTypes.CustomAppCrashRate: NewResourceDescriptor(
		custom_app_crash_rate.Service,
		Coalesce(Dependencies.CustomApplication),
	),
	ResourceTypes.MobileAppAnomalies: NewResourceDescriptor(
		mobile_app_anomalies.Service,
		Coalesce(Dependencies.DeviceApplicationMethod),
		Coalesce(Dependencies.MobileApplication),
	),
	ResourceTypes.MobileAppCrashRate: NewResourceDescriptor(
		mobile_app_crash_rate.Service,
		Coalesce(Dependencies.MobileApplication),
	),
	ResourceTypes.WebAppAnomalies: NewResourceDescriptor(
		web_app_anomalies.Service,
		Coalesce(Dependencies.Application),
		Coalesce(Dependencies.ApplicationMethod),
	),
	ResourceTypes.CustomService: NewResourceDescriptor(customservices.Service),
	ResourceTypes.Credentials: NewResourceDescriptor(
		vault.Service,
		Dependencies.ID(ResourceTypes.Credentials),
	),
	ResourceTypes.JSONDashboard: NewResourceDescriptor(
		jsondashboards.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
		Dependencies.ManagementZone,
		Dependencies.Service,
		Dependencies.ID(ResourceTypes.SLO),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.MobileApplication),
		Dependencies.ID(ResourceTypes.SyntheticLocation),
		Dependencies.ID(ResourceTypes.HTTPMonitor),
		Dependencies.ID(ResourceTypes.CalculatedServiceMetric),
		Dependencies.ID(ResourceTypes.BrowserMonitor),
	),
	ResourceTypes.DashboardSharing: NewResourceDescriptor(
		sharing.Service,
		Dependencies.ID(ResourceTypes.JSONDashboard),
	),
	ResourceTypes.DatabaseAnomalies:  NewResourceDescriptor(database_anomalies.Service),
	ResourceTypes.DiskEventAnomalies: NewResourceDescriptor(disk_event_anomalies.Service),
	ResourceTypes.DiskAnomaliesV2: NewResourceDescriptor(
		disk_anomalies_v2.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.DiskSpecificAnomaliesV2: NewResourceDescriptor(
		disk_specific_anomalies_v2.Service,
		Coalesce(Dependencies.Disk),
	),
	ResourceTypes.EmailNotification: NewResourceDescriptor(
		email.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Email),
	ResourceTypes.FrequentIssues: NewResourceDescriptor(frequentissues.Service),
	ResourceTypes.HostAnomalies:  NewResourceDescriptor(host_anomalies.Service),
	ResourceTypes.HostAnomaliesV2: NewResourceDescriptor(
		host_anomalies_v2.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.HTTPMonitor: NewResourceDescriptor(
		http.Service,
		Dependencies.ID(ResourceTypes.SyntheticLocation),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.Credentials),
	),
	ResourceTypes.HostNaming:   NewResourceDescriptor(host_naming.Service),
	ResourceTypes.IBMMQFilters: NewResourceDescriptor(filters.Service),
	ResourceTypes.IMSBridge:    NewResourceDescriptor(imsbridges.Service),
	ResourceTypes.JiraNotification: NewResourceDescriptor(
		jira.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Jira),
	ResourceTypes.KeyRequests: NewResourceDescriptor(
		keyrequests.Service,
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.KubernetesCredentials: NewResourceDescriptor(kubernetes.Service),
	ResourceTypes.Maintenance: NewResourceDescriptor(
		v2maintenance.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.ManagementZoneV2: NewResourceDescriptor(v2managementzones.Service),
	ResourceTypes.MetricEvents:     NewResourceDescriptor(metricevents.Service),
	ResourceTypes.MobileApplication: NewResourceDescriptor(
		mobile.Service,
		Dependencies.ID(ResourceTypes.RequestAttribute),
	),
	ResourceTypes.MutedRequests: NewResourceDescriptor(
		mutedrequests.Service,
		Coalesce(Dependencies.Service),
	),
	ResourceTypes.NetworkZone:  NewResourceDescriptor(networkzone.Service),
	ResourceTypes.NetworkZones: NewResourceDescriptor(networkzones.Service),
	ResourceTypes.OpsGenieNotification: NewResourceDescriptor(
		opsgenie.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.OpsGenie),
	ResourceTypes.PagerDutyNotification: NewResourceDescriptor(
		pagerduty.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.PagerDuty),
	ResourceTypes.ProcessGroupNaming: NewResourceDescriptor(processgroup_naming.Service),
	ResourceTypes.QueueManager:       NewResourceDescriptor(queuemanagers.Service),
	ResourceTypes.RequestAttribute: NewResourceDescriptor(
		requestattributes.Service,
		Dependencies.Host,
		Dependencies.HostGroup,
		Dependencies.ProcessGroup,
		Dependencies.ProcessGroupInstance,
		Dependencies.Service,
	),
	ResourceTypes.RequestNaming: NewResourceDescriptor(
		requestnaming.Service,
		Dependencies.RequestAttribute,
	),
	ResourceTypes.ResourceAttributes: NewResourceDescriptor(resattr.Service),
	ResourceTypes.ServiceAnomalies:   NewResourceDescriptor(service_anomalies.Service),
	ResourceTypes.ServiceAnomaliesV2: NewResourceDescriptor(
		service_anomalies_v2.Service,
		Coalesce(Dependencies.ServiceMethod),
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ServiceNaming: NewResourceDescriptor(service_naming.Service),
	ResourceTypes.ServiceNowNotification: NewResourceDescriptor(
		servicenow.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.ServiceNow),
	ResourceTypes.SlackNotification: NewResourceDescriptor(
		slack.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Slack),
	ResourceTypes.SLO: NewResourceDescriptor(
		slo.Service,
		Dependencies.ManagementZone,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
		Dependencies.ID(ResourceTypes.CalculatedServiceMetric),
	),
	ResourceTypes.SpanAttribute:          NewResourceDescriptor(attributes.Service),
	ResourceTypes.SpanCaptureRule:        NewResourceDescriptor(capture.Service),
	ResourceTypes.SpanContextPropagation: NewResourceDescriptor(ctxprop.Service),
	ResourceTypes.SpanEntryPoint:         NewResourceDescriptor(entrypoints.Service),
	ResourceTypes.SyntheticLocation:      NewResourceDescriptor(locations.Service),
	ResourceTypes.TrelloNotification: NewResourceDescriptor(
		trello.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Trello),
	ResourceTypes.VictorOpsNotification: NewResourceDescriptor(
		victorops.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.VictorOps),
	ResourceTypes.WebApplication: NewResourceDescriptor(
		web.Service,
		Dependencies.ID(ResourceTypes.RequestAttribute),
	),
	ResourceTypes.WebHookNotification: NewResourceDescriptor(
		webhook.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.WebHook),
	ResourceTypes.XMattersNotification: NewResourceDescriptor(
		xmatters.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.XMatters),

	ResourceTypes.MaintenanceWindow: NewResourceDescriptor(
		maintenancev1.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.ManagementZone: NewResourceDescriptor(managementzonesv1.Service),
	ResourceTypes.Dashboard: NewResourceDescriptor(
		dashboards.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
		Dependencies.ManagementZone,
		Dependencies.ID(ResourceTypes.SLO),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.SyntheticLocation),
	),
	ResourceTypes.Notification: NewResourceDescriptor(
		notificationsv1.Service,
		Dependencies.LegacyID(ResourceTypes.Alerting),
	),
	ResourceTypes.QueueSharingGroups: NewResourceDescriptor(queuesharinggroups.Service),
	ResourceTypes.AlertingProfile: NewResourceDescriptor(
		alertingv1.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.RequestNamings: NewResourceDescriptor(
		order.Service,
		Dependencies.ID(ResourceTypes.RequestNaming),
	),
	ResourceTypes.IAMUser:  NewResourceDescriptor(users.Service),
	ResourceTypes.IAMGroup: NewResourceDescriptor(groups.Service),
	ResourceTypes.DDUPool:  NewResourceDescriptor(ddupool.Service),
	ResourceTypes.ProcessGroupAnomalies: NewResourceDescriptor(
		pg_anomalies.Service,
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.ProcessGroupAlerting: NewResourceDescriptor(
		processgroupalerting.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.DatabaseAnomaliesV2: NewResourceDescriptor(
		database_anomalies_v2.Service,
		Coalesce(Dependencies.ServiceMethod),
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessMonitoringRule: NewResourceDescriptor(
		customprocessmonitoring.Service,
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessMonitoring: NewResourceDescriptor(
		processmonitoring.Service,
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessAvailability: NewResourceDescriptor(
		processavailability.Service,
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.AdvancedProcessGroupDetectionRule: NewResourceDescriptor(advanceddetectionrule.Service),
	ResourceTypes.ConnectivityAlerts: NewResourceDescriptor(
		connectivityalerts.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.DeclarativeGrouping: NewResourceDescriptor(
		declarativegrouping.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.HostMonitoring: NewResourceDescriptor(
		hostmonitoring.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.HostProcessGroupMonitoring: NewResourceDescriptor(
		hostprocessgroupmonitoring.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.RUMIPLocations: NewResourceDescriptor(ipmappings.Service),
	ResourceTypes.CustomAppEnablement: NewResourceDescriptor(
		rumcustomenablement.Service,
		Dependencies.ID(ResourceTypes.MobileApplication),
	),
	ResourceTypes.MobileAppEnablement: NewResourceDescriptor(
		rummobileenablement.Service,
		Dependencies.ID(ResourceTypes.MobileApplication),
	),
	ResourceTypes.WebAppEnablement: NewResourceDescriptor(
		rumwebenablement.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.RUMProcessGroup: NewResourceDescriptor(
		rumprocessgroup.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.RUMProviderBreakdown:  NewResourceDescriptor(rumproviderbreakdown.Service),
	ResourceTypes.UserExperienceScore:   NewResourceDescriptor(userexperiencescore.Service),
	ResourceTypes.WebAppResourceCleanup: NewResourceDescriptor(webappresourcecleanup.Service),
	ResourceTypes.UpdateWindows:         NewResourceDescriptor(updatewindows.Service),
	ResourceTypes.ProcessGroupDetectionFlags: NewResourceDescriptor(
		detectionflags.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessGroupMonitoring: NewResourceDescriptor(
		processgroupmonitoring.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.ProcessGroupSimpleDetection: NewResourceDescriptor(processgroupsimpledetection.Service),
	ResourceTypes.LogMetrics:                  NewResourceDescriptor(schemalesslogmetric.Service),
	ResourceTypes.BrowserMonitorPerformanceThresholds: NewResourceDescriptor(
		browserperformancethresholds.Service,
		Coalesce(Dependencies.SyntheticTest),
	),
	ResourceTypes.HttpMonitorPerformanceThresholds: NewResourceDescriptor(
		httpperformancethresholds.Service,
		Coalesce(Dependencies.HttpCheck),
	),
	ResourceTypes.HttpMonitorCookies: NewResourceDescriptor(
		httpcookies.Service,
		Coalesce(Dependencies.HttpCheck),
	),
	ResourceTypes.SessionReplayWebPrivacy: NewResourceDescriptor(
		sessionreplaywebprivacy.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.SessionReplayResourceCapture: NewResourceDescriptor(
		sessionreplayresourcecapture.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.UsabilityAnalytics: NewResourceDescriptor(
		analytics.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.SyntheticAvailability: NewResourceDescriptor(availability.Service),
	ResourceTypes.BrowserMonitorOutageHandling: NewResourceDescriptor(
		browseroutagehandling.Service,
		Coalesce(Dependencies.SyntheticTest),
	),
	ResourceTypes.HttpMonitorOutageHandling: NewResourceDescriptor(
		httpoutagehandling.Service,
		Coalesce(Dependencies.HttpCheck),
	),
	ResourceTypes.CloudAppWorkloadDetection:      NewResourceDescriptor(workloaddetection.Service),
	ResourceTypes.MainframeTransactionMonitoring: NewResourceDescriptor(txmonitoring.Service),
	ResourceTypes.MonitoredTechnologiesApache: NewResourceDescriptor(
		apache.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesDotNet: NewResourceDescriptor(
		dotnet.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesGo: NewResourceDescriptor(
		golang.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesIIS: NewResourceDescriptor(
		iis.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesJava: NewResourceDescriptor(
		java.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesNGINX: NewResourceDescriptor(
		nginx.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesNodeJS: NewResourceDescriptor(
		nodejs.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesOpenTracing: NewResourceDescriptor(
		opentracingnative.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesPHP: NewResourceDescriptor(
		php.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesVarnish: NewResourceDescriptor(
		varnish.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesWSMB: NewResourceDescriptor(
		wsmb.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.ProcessVisibility: NewResourceDescriptor(
		processvisibility.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup)),
	ResourceTypes.FailureDetectionEnvironmentParameter: NewResourceDescriptor(environmentparameters.Service),
}

var BlackListedResources = []ResourceType{
	ResourceTypes.MaintenanceWindow, // legacy
	ResourceTypes.ManagementZone,    // legacy
	ResourceTypes.Notification,      // legacy
	ResourceTypes.AlertingProfile,   // legacy
	ResourceTypes.Dashboard,         // taken care of dynatrace_json_dashboard
	ResourceTypes.IAMUser,           // not sure whether to migrate
	ResourceTypes.IAMGroup,          // not sure whether to migrate

	// excluding by default
	ResourceTypes.JSONDashboard, // may replace dynatrace_dashboard in the future
	ResourceTypes.DashboardSharing,

	ResourceTypes.ProcessGroupAnomalies, // there could be 100k process groups

	ResourceTypes.WebAppEnablement,             // overlap with ResourceTypes.MobileApplication
	ResourceTypes.MobileAppEnablement,          // overlap with ResourceTypes.MobileApplication
	ResourceTypes.CustomAppEnablement,          // overlap with ResourceTypes.MobileApplication
	ResourceTypes.SessionReplayWebPrivacy,      // overlap with ResourceTypes.ApplicationDataPrivacy
	ResourceTypes.SessionReplayResourceCapture, // overlap with ResourceTypes.WebApplication
	ResourceTypes.BrowserMonitorOutageHandling, // overlap with ResourceTypes.BrowserMonitor
	ResourceTypes.HttpMonitorOutageHandling,    // overlap with ResourceTypes.HTTPMonitor
}

func Service(credentials *settings.Credentials, resourceType ResourceType) settings.CRUDService[settings.Settings] {
	return AllResources[resourceType].Service(credentials)
}

// func DSService(credentials *settings.Credentials, dataSourceType DataSourceType) settings.RService[settings.Settings] {
// 	return AllDataSources[dataSourceType].Service(credentials)
// }
