package environmentparameter

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FailureDetectionEnvironmentParameter struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	HTTPResponseCodes HTTPResponseCodes `json:"httpResponseCodes"`
	BrokenLinks       BrokenLinks       `json:"brokenLinks"`
	ExceptionRules    ExceptionRules    `json:"exceptionRules"`
}

func (me *FailureDetectionEnvironmentParameter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metrics": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Metrics",
			Elem: &schema.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"log_monitoring": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Log Monitoring",
			Elem: &schema.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"serverless": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Serverless",
			Elem: &schema.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"events": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Events",
			Elem: &schema.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"traces": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Traces",
			Elem: &schema.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
	}
}

func (me *DDUPool) MarshalHCL(properties hcl.Properties) error {

	if me.MetricsPool.LimitEnabled {
		if err := properties.Encode("metrics", &me.MetricsPool); err != nil {
			return err
		}
	}
	if me.LogMonitoringPool.LimitEnabled {
		if err := properties.Encode("log_monitoring", &me.LogMonitoringPool); err != nil {
			return err
		}
	}
	if me.ServerlessPool.LimitEnabled {
		if err := properties.Encode("serverless", &me.ServerlessPool); err != nil {
			return err
		}
	}
	if me.EventsPool.LimitEnabled {
		if err := properties.Encode("events", &me.EventsPool); err != nil {
			return err
		}
	}
	if me.TracesPool.LimitEnabled {
		if err := properties.Encode("traces", &me.TracesPool); err != nil {
			return err
		}
	}

	return nil
}

func (me *DDUPool) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"metrics":        &me.MetricsPool,
		"log_monitoring": &me.LogMonitoringPool,
		"serverless":     &me.ServerlessPool,
		"events":         &me.EventsPool,
		"traces":         &me.TracesPool,
	})
}
