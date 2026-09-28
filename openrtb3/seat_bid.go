package openrtb3

import (
	"github.com/rtb-0/openrtb/common"
)

// SeatBid is an OpenRTB seatbid object (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_seatbid
type SeatBid[A any] struct {
	// ID of the buyer seat on whose behalf this bid is made.
	Seat string `json:"seat,omitempty"`
	// For offers with multiple items, this flag Indicates if the bidder is willing to accept wins on a subset of bids or requires the full group as a package, where 0 = individual wins accepted; 1 = package win or loss only.
	Package int8 `json:"package,omitempty"`
	// Array of 1+ Bid objects each related to an item.
	Bid []Bid[A] `json:"bid,omitempty"`
	// Optional demand source specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
