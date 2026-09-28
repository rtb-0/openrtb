package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Segment objects are essentially key-value pairs that convey specific units of data (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_segment
type Segment struct {
	// ID of the data segment specific to the data provider.
	ID string `json:"id,omitempty"`
	// Displayable name of the data segment specific to the data provider.
	Name string `json:"name,omitempty"`
	// String representation of the data segment value.
	Value string `json:"value,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
