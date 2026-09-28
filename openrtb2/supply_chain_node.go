package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object is associated with a SupplyChain object as an array of nodes (OpenRTB 2.6 §3.2.26).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectsupplychainnode
type SupplyChainNode struct {
	// The canonical domain name of the SSP, Exchange, Header Wrapper, etc system that bidders connect to.
	ASI string `json:"asi"`
	// The identifier associated with the seller or reseller account within the advertising system.
	SID string `json:"sid"`
	// The OpenRTB RequestId of the request as issued by this seller.
	RID string `json:"rid,omitempty"`
	// The name of the company (the legal entity) that is paid for inventory transacted under the given seller_ID.
	Name string `json:"name,omitempty"`
	// The business domain name of the entity represented by this node.
	Domain string `json:"domain,omitempty"`
	// Indicates whether this node will be involved in the flow of payment for the inventory.
	HP *int8 `json:"hp,omitempty"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
