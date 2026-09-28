package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object describes the nature and behavior of the entity that is the source of the bid request upstream from the exchange (OpenRTB 2.6 §3.2.2).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectsource
type Source struct {
	// Entity responsible for the final impression sale decision, where 0 = exchange, 1 = upstream source.
	FD *int8 `json:"fd,omitempty"`
	// Transaction ID that must be common across all participants in this bid request (e.g., potentially multiple exchanges).
	TID string `json:"tid,omitempty"`
	// Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol v1.0.
	PChain string `json:"pchain,omitempty"`
	// This object represents both the links in the supply chain as well as an indicator whether or not the supply chain is complete.
	SChain *SupplyChain `json:"schain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
