package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object is the private marketplace container for direct deals between buyers and sellers that may pertain to this impression (OpenRTB 2.6 §3.2.11).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectpmp
type PMP struct {
	// Indicator of auction eligibility to seats named in the Direct Deals object, where 0 = all bids are accepted, 1 = bids are restricted to the deals specified and the terms thereof.
	PrivateAuction int8 `json:"private_auction,omitempty"`
	// Array of Deal (Section 3.2.12) objects that convey the specific deals applicable to this impression.
	Deals []Deal `json:"deals,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
