package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// Segment objects are essentially key-value pairs that convey specific units of data (OpenRTB 2.6 §3.2.22).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectsegment
type Segment struct {
	// ID of the data segment specific to the data provider.
	ID string `json:"id,omitempty"`
	// Name of the data segment specific to the data provider.
	Name string `json:"name,omitempty"`
	// String representation of the data segment value.
	Value string `json:"value,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
