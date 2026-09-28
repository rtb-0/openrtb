package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object is associated with an impression as an array of metrics (OpenRTB 2.6 §3.2.5).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectmetric
type Metric struct {
	// Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Type string `json:"type"`
	// Number representing the value of the metric.
	Value float64 `json:"value,omitempty"`
	// Source of the value using exchange curated string names which should be published to bidders a priori.
	Vendor string `json:"vendor,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
