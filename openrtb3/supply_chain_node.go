package openrtb3

import "github.com/rtb-0/openrtb/common"

// SupplyChainNode is one entity in a SupplyChain.
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/master/supplychainobject.md#openrtb-object-supplychainnode
type SupplyChainNode struct {
	// Canonical domain of the advertising system this node belongs to.
	ASI string `json:"asi"`
	// Seller or reseller account id within that advertising system.
	SID string `json:"sid"`
	// OpenRTB request id issued by this seller.
	RID string `json:"rid,omitempty"`
	// Name of the company paid for inventory under this seller id.
	Name string `json:"name,omitempty"`
	// Business domain of the entity represented by this node.
	Domain string `json:"domain,omitempty"`
	// 1 when this node is involved in the flow of payment.
	HP *int8 `json:"hp,omitempty"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
