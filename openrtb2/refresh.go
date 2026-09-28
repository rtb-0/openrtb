package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// Refresh is defined by OpenRTB 2.6 (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectrefresh
type Refresh struct {
	// A RefSettings object (see Section 3.2.34) describing the mechanics of how an ad placement automatically refreshes.
	RefSettings []RefSettings `json:"refsettings,omitempty"`
	// The number of times this ad slot had been refreshed since last page load.
	Count *int `json:"count,omitempty"`
	// Placeholder for vendor specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
