package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// A programmatic impression is often referred to as a ‘spot’ in digital out-of-home and CTV, with an impression being a unique member of the audience viewing it (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectqty
type Qty struct {
	// The quantity of billable events which will be deemed to have occurred if this item is purchased.
	Multiplier float64 `json:"multiplier,omitempty"`
	// The source type of the quantity measurement, ie.
	SourceType adcom1.DOOHMultiplierMeasurementSourceType `json:"sourcetype,omitempty"`
	// The top level business domain name of the measurement vendor providing the quantity measurement.
	Vendor string `json:"vendor,omitempty"`
	// Placeholder for vendor specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
