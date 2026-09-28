package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object allows sellers to specify price floors for video and audio creatives, whose price varies based on time (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectdurfloors
type DurFloors struct {
	// An integer indicating the low end of a duration range.
	MinDur int64 `json:"mindur,omitempty"`
	// An integer indicating the high end of a duration range.
	MaxDur int64 `json:"maxdur,omitempty"`
	// Minimum bid for a given impression opportunity, if bidding with a creative in this duration range, expressed in CPM.
	BidFloor float64 `json:"bidfloor,omitempty"`
	// Placeholder for vendor specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
