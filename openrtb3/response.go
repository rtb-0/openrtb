package openrtb3

import (
	"github.com/rtb-0/openrtb/common"
)

// Response object is the bid response object under the Openrtb root (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_response
type Response[A any] struct {
	// ID of the bid request to which this is a response; must match the request.id attribute.
	ID string `json:"id,omitempty"`
	// Bidder generated response ID to assist with logging/tracking.
	BidID string `json:"bidid,omitempty"`
	// Reason for not bidding if applicable (see List: No-Bid Reason Codes).
	NBR NoBidReason `json:"nbr,omitempty"`
	// Bid currency using ISO-4217 alpha codes.
	Cur string `json:"cur,omitempty"`
	// Allows bidder to set data in the exchange’s cookie, which can be retrieved on bid requests (refer to cdata in Object: Request) if supported by the exchange.
	CData string `json:"cdata,omitempty"`
	// Array of Seatbid objects; 1+ required if a bid is to be made.
	SeatBid []SeatBid[A] `json:"seatbid,omitempty"`
	// Optional demand source specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
