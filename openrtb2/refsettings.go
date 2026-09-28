package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// Information on how often and what triggers an ad slot being refreshed (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectrefsettings
type RefSettings struct {
	// The type of the declared auto refresh.
	RefType adcom1.AutoRefreshTrigger `json:"reftype,omitempty"`
	// The minimum refresh interval in seconds.
	MinInt int `json:"minint,omitempty"`
	// Placeholder for vendor specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
