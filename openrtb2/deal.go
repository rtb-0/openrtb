package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object constitutes a specific deal that was struck a priori between a buyer and a seller (OpenRTB 2.6 §3.2.12).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectdeal
type Deal struct {
	// A unique identifier for the direct deal.
	ID string `json:"id"`
	// Minimum bid for this impression expressed in CPM.
	BidFloor float64 `json:"bidfloor,omitempty"`
	// Currency specified using ISO-4217 alpha codes.
	BidFloorCur string `json:"bidfloorcur,omitempty"`
	// Optional override of the overall auction type of the bid request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in bidfloor is the agreed upon deal price.
	AT int64 `json:"at,omitempty"`
	// Whitelist of buyer seats (e.g., advertisers, agencies) allowed to bid on this deal.
	WSeat []string `json:"wseat,omitempty"`
	// Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal.
	WADomain []string `json:"wadomain,omitempty"`
	// Indicates that the deal is of type guaranteed and the bidder must bid on the deal, where 0 = not a guaranteed deal, 1 = guaranteed deal.
	Guar int8 `json:"guar,omitempty"`
	// Minimum CPM per second.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`
	// Container for floor price by duration information, to be used if a given deal is eligible for video or audio demand.
	DurFloors []DurFloors `json:"durfloors,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
