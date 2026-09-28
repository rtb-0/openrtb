package openrtb3

import "github.com/rtb-0/openrtb/common"

// SupplyChain is the chain of entities that participate in a transaction.
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/master/supplychainobject.md#openrtb-object-supplychain
type SupplyChain struct {
	// Flag indicating whether the chain contains all nodes back to the inventory owner, where 0 = no, 1 = yes.
	Complete int8 `json:"complete"`
	// Nodes in the order of the chain.
	Nodes []SupplyChainNode `json:"nodes"`
	// Version of the supply chain specification in use, in the format of "major.minor".
	Ver string `json:"ver"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
