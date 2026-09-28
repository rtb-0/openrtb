package openrtb3

import (
	"github.com/rtb-0/openrtb/common"
)

// Metric object is associated with an item as an array of metrics (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_metric
type Metric struct {
	// Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Type string `json:"type"`
	// Number representing the value of the metric.
	Value float64 `json:"value,omitempty"`
	// Source of the value using exchange curated string names which should be published to bidders a priori.
	Vendor string `json:"vendor,omitempty"`
	// Optional exchange-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
