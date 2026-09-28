package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/openrtb3"
)

// This object is the top-level bid response object (i.e., the unnamed outer JSON object) (OpenRTB 2.6 §4.2.1).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbidresponse
type BidResponse[A any] struct {
	// ID of the bid request to which this is a response.
	ID string `json:"id"`
	// Array of seatbid objects; 1+ required if a bid is to be made.
	SeatBid []SeatBid[A] `json:"seatbid,omitempty"`
	// Bidder generated response ID to assist with logging/tracking.
	BidID string `json:"bidid,omitempty"`
	// Bid currency using ISO-4217 alpha codes.
	Cur string `json:"cur,omitempty"`
	// Optional feature to allow a bidder to set data in the exchange’s cookie.
	CustomData string `json:"customdata,omitempty"`
	// Reason for not bidding.
	NBR *openrtb3.NoBidReason `json:"nbr,omitempty"`
	// Placeholder for bidder-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
