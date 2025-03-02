package download

import (
	"fmt"
	"os"
	"strings"

	"github.com/dtcookie/dynatrace/api/config/applications/web"
	"github.com/dtcookie/dynatrace/api/config/dashboards"
	"github.com/dtcookie/dynatrace/api/config/metrics/calculated/service"
	service_naming "github.com/dtcookie/dynatrace/api/config/naming/services"
	"github.com/dtcookie/dynatrace/api/config/requestattributes"
	privlocations "github.com/dtcookie/dynatrace/api/config/synthetic/locations"
)

func creds() (string, string) {
	environmentURL := os.Getenv("DYNATRACE_TARGET_ENV_URL")
	environmentURL = strings.TrimSuffix(strings.TrimSuffix(environmentURL, " "), "/")
	apiToken := os.Getenv("DYNATRACE_TARGET_API_TOKEN")

	if len(environmentURL) == 0 || len(apiToken) == 0 {
		environmentURL = os.Getenv("DYNATRACE_ENV_URL")
		environmentURL = strings.TrimSuffix(strings.TrimSuffix(environmentURL, " "), "/")
		apiToken = os.Getenv("DYNATRACE_API_TOKEN")
	}
	return environmentURL, apiToken
}

var MissingRequiredScope = false

var InterventionInfoMap = map[string]InterventionStruct{
	"dynatrace_dashboard": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				dashboard := resource.RESTObject.(*dashboards.Dashboard)
				if dashboard.Metadata.Owner != nil && *dashboard.Metadata.Owner == "Dynatrace" {
					resource.ReqInter.Type = InterventionTypes.ReqAttn
					resource.ReqInter.Message = []string{"ATTENTION " + "Dashboards owned by Dynatrace are automatically excluded to prevent duplicates of OOTB dashboards. Please return to the dashboards folder if this is a custom dashboard."}
					continue
				}

				if !MissingRequiredScope {
					dbId := dashboard.ID
					dashboard.ID = nil

					environmentURL, apiToken := creds()

					client := dashboards.NewService(environmentURL+"/api/config/v1", apiToken)
					errors := client.Validate(dashboard)

					dashboard.ID = dbId
					if len(errors) > 0 {
						if strings.Contains(errors[0], "Token is missing required scope. Use one of: WriteConfig (Write configuration)") {
							MissingRequiredScope = true
						} else {
							resource.ReqInter.Type = InterventionTypes.Flawed
							errors[0] = "ATTENTION " + strings.ReplaceAll(errors[0], "\n", "")
							resource.ReqInter.Message = errors
							continue
						}
					}
				}
			}
		},
	},
	// "dynatrace_custom_anomalies": {
	// 	Move: func(resName string, resourceData ResourceData) {
	// 		for _, resource := range resourceData[resName] {
	// 			config := resource.RESTObject.(*metricevents.MetricEvent)
	// 			configID := config.ID
	// 			config.ID = nil

	// 			environmentURL, apiToken := creds()
	// 			client := metricevents.NewService(environmentURL+"/api/config/v1", apiToken)
	// 			var errors []string
	// 			err := client.Validate(config)
	// 			config.ID = configID

	// 			if err != nil {
	// 				if env, ok := err.(*rest.Error); ok {
	// 					if len(env.ConstraintViolations) > 0 {
	// 						errs := []string{}
	// 						for _, violation := range env.ConstraintViolations {
	// 							errs = append(errs, violation.Message)
	// 						}
	// 						errors = errs
	// 					} else {
	// 						errors = []string{err.Error()}
	// 					}
	// 				} else {
	// 					errors = []string{err.Error()}
	// 				}
	// 				if len(errors) > 0 && !strings.Contains(errors[0], "Token is missing required scope. Use one of: WriteConfig (Write configuration)") {
	// 					resource.ReqInter.Type = InterventionTypes.Flawed
	// 					errors[0] = "ATTENTION " + strings.ReplaceAll(errors[0], "\n", "")
	// 					resource.ReqInter.Message = errors
	// 					continue
	// 				}
	// 			}
	// 		}
	// 	},
	// },
	"dynatrace_calculated_service_metric": {
		Move: func(resName string, resourceData ResourceData) {
			reqConditions := []string{"SERVICE_DISPLAY_NAME", "SERVICE_PUBLIC_DOMAIN_NAME", "SERVICE_WEB_APPLICATION_ID", "SERVICE_WEB_CONTEXT_ROOT", "SERVICE_WEB_SERVER_NAME", "SERVICE_WEB_SERVICE_NAME", "SERVICE_WEB_SERVICE_NAMESPACE", "REMOTE_SERVICE_NAME", "REMOTE_ENDPOINT", "AZURE_FUNCTIONS_SITE_NAME", "AZURE_FUNCTIONS_FUNCTION_NAME", "CTG_GATEWAY_URL", "CTG_SERVER_NAME", "ACTOR_SYSTEM", "ESB_APPLICATION_NAME", "SERVICE_TAG", "SERVICE_TYPE", "PROCESS_GROUP_TAG", "PROCESS_GROUP_NAME"}
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				dataObj := resource.RESTObject.(*service.CalculatedServiceMetric)
				if len(dataObj.ManagementZones) == 0 && dataObj.Conditions != nil {
					var found bool
					for _, condition := range dataObj.Conditions {
						for _, reqCondition := range reqConditions {
							if string(condition.Attribute) == reqCondition {
								found = true
							}
						}
					}
					if !found {
						resource.ReqInter.Type = InterventionTypes.Flawed
						resource.ReqInter.Message = []string{"ATTENTION " + "The metric needs to either get limited by specifying a Management Zone or by specifying one or more conditions related to SERVICE_DISPLAY_NAME, SERVICE_PUBLIC_DOMAIN_NAME, SERVICE_WEB_APPLICATION_ID, SERVICE_WEB_CONTEXT_ROOT, SERVICE_WEB_SERVER_NAME, SERVICE_WEB_SERVICE_NAME, SERVICE_WEB_SERVICE_NAMESPACE, REMOTE_SERVICE_NAME, REMOTE_ENDPOINT, AZURE_FUNCTIONS_SITE_NAME, AZURE_FUNCTIONS_FUNCTION_NAME, CTG_GATEWAY_URL, CTG_SERVER_NAME, ACTOR_SYSTEM, ESB_APPLICATION_NAME, SERVICE_TAG, SERVICE_TYPE, PROCESS_GROUP_TAG or PROCESS_GROUP_NAME"}
					}
				}

			}
		},
	},
	"dynatrace_credentials": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				resource.ReqInter.Type = InterventionTypes.ReqAttn
				resource.ReqInter.Message = []string{"ATTENTION " + "REST API didn't provide credentials"}
			}
		},
	},
	"dynatrace_aws_credentials": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				resource.ReqInter.Type = InterventionTypes.ReqAttn
				resource.ReqInter.Message = []string{"ATTENTION " + "REST API didn't provide credentials"}
			}
		},
	},
	"dynatrace_cloudfoundry_credentials": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				resource.ReqInter.Type = InterventionTypes.ReqAttn
				resource.ReqInter.Message = []string{"ATTENTION " + "REST API didn't provide credentials"}
			}
		},
	},
	"dynatrace_azure_credentials": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				resource.ReqInter.Type = InterventionTypes.ReqAttn
				resource.ReqInter.Message = []string{"ATTENTION " + "REST API didn't provide credentials"}
			}
		},
	},
	"dynatrace_k8s_credentials": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				resource.ReqInter.Type = InterventionTypes.ReqAttn
				resource.ReqInter.Message = []string{"ATTENTION " + "REST API didn't provide credentials"}
			}
		},
	},
	"dynatrace_synthetic_location": {
		StripIDs: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				dataObj := resource.RESTObject.(*privlocations.PrivateSyntheticLocation)
				if len(dataObj.Nodes) > 0 {
					dataObj.Nodes = []string{}
				}
			}
		},
	},
	"dynatrace_request_attribute": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				dataObj := resource.RESTObject.(*requestattributes.RequestAttribute)
				if len(dataObj.DataSources) > 0 {
					for _, dataSource := range dataObj.DataSources {
						if dataSource.Scope != nil {
							if dataSource.Scope.HostGroup != nil {
								resource.ReqInter.Type = InterventionTypes.ReqAttn
								resource.ReqInter.Message = []string{fmt.Sprintf("ATTENTION Data Source refers to entity '%s' - which may not exist on the target environment", *dataSource.Scope.HostGroup)}
								break
							} else if dataSource.Scope.ProcessGroup != nil {
								resource.ReqInter.Type = InterventionTypes.ReqAttn
								resource.ReqInter.Message = []string{fmt.Sprintf("ATTENTION Data Source refers to entity '%s' - which may not exist on the target environment", *dataSource.Scope.ProcessGroup)}
								break
							}
						}

					}
				}
			}
		},
	},
	"dynatrace_service_naming": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				dataObj := resource.RESTObject.(*service_naming.NamingRule)
				if strings.Contains(dataObj.Format, "ProcessGroup:Environment:") {
					formatSnippet := dataObj.Format[strings.Index(dataObj.Format, "ProcessGroup:Environment:")+len("ProcessGroup:Environment:"):]
					idx := strings.Index(formatSnippet, "}")
					if idx >= 0 {
						formatSnippet = formatSnippet[:idx]
						resource.ReqInter.Type = InterventionTypes.ReqAttn
						resource.ReqInter.Message = []string{fmt.Sprintf("ATTENTION {ProcessGroup:Environment:%s} may not exist on the target environment", formatSnippet)}
					}
				}
			}
		},
	},
	"dynatrace_web_application": {
		Move: func(resName string, resourceData ResourceData) {
			for _, resource := range resourceData[resName] {
				if resource.ReqInter.Evaluated {
					return
				}
				resource.ReqInter.Evaluated = true
				dataObj := resource.RESTObject.(*web.ApplicationConfig)
				if string(dataObj.Type) == "BROWSER_EXTENSION_INJECTED" {
					if dataObj.URLInjectionPattern == nil {
						resource.ReqInter.Type = InterventionTypes.ReqAttn
						resource.ReqInter.Message = []string{"ATTENTION REST API didn't provide the URL Injection Pattern - which is required for type BROWSER_EXTENSION_INJECTED"}
					}
				}
				if len(dataObj.UserActionAndSessionProperties) > 20 {
					resource.ReqInter.Type = InterventionTypes.ReqAttn
					resource.ReqInter.Message = append(resource.ReqInter.Message, "ATTENTION You cannot have more than 20 useraction/session properties")
				}
			}
		},
	},
}

type InterventionStruct struct {
	Move     func(string, ResourceData)
	StripIDs func(string, ResourceData)
}

func (me ResourceData) RequiresIntervention(dlConfig DownloadConfig) error {
	for resName := range me {
		if _, exists := InterventionInfoMap[resName]; exists {
			if InterventionInfoMap[resName].Move != nil {
				InterventionInfoMap[resName].Move(resName, me)
			}
			if dlConfig.Migrate && InterventionInfoMap[resName].StripIDs != nil {
				InterventionInfoMap[resName].StripIDs(resName, me)
			}
		}
	}
	return nil
}
