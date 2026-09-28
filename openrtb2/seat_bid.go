package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// A bid response can contain multiple SeatBid objects, each on behalf of a different bidder seat and each containing one or more individual bids (OpenRTB 2.6 §4.2.2).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectseatbid
type SeatBid[A any] struct {
	// Array of 1+ Bid objects (Section 4.2.3) each related to an impression.
	Bid []Bid[A] `json:"bid"`
	// ID of the buyer seat (e.g., advertiser, agency) on whose behalf this bid is made.
	Seat string `json:"seat,omitempty"`
	// 0 = impressions can be won individually; 1 = impressions must be won or lost as a group.
	Group int8 `json:"group,omitempty"`
	// Placeholder for bidder-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
