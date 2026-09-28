package openrtb3

import (
	"github.com/rtb-0/openrtb/common"
)

// Deal object constitutes a specific deal that was struck a priori between a seller and a buyer (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_deal
type Deal struct {
	// A unique identifier for the deal.
	ID string `json:"id,omitempty"`
	// Minimum deal price for this item expressed in CPM.
	Flr float64 `json:"flr,omitempty"`
	// Currency of the flr attribute specified using ISO-4217 alpha codes.
	FlrCur string `json:"flrcur,omitempty"`
	// Optional override of the overall auction type of the request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in flr is the agreed upon deal price.
	AT AuctionType `json:"at,omitempty"`
	// Whitelist of buyer seats allowed to bid on this deal.
	WSeat []string `json:"wseat,omitempty"`
	// Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal.
	WADomain []string `json:"wadomain,omitempty"`
	// Optional exchange-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
