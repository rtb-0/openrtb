package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object is composed of a set of nodes where each node represents a specific entity that participates in the transacting of inventory (OpenRTB 2.6 §3.2.25).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectsupplychain
type SupplyChain struct {
	// Flag indicating whether the chain contains all nodes involved in the transaction leading back to the owner of the site, app or other medium of the inventory, where 0 = no, 1 = yes.
	Complete int8 `json:"complete"`
	// Array of SupplyChainNode objects in the order of the chain.
	Nodes []SupplyChainNode `json:"nodes"`
	// Version of the supply chain specification in use, in the format of "major.minor".
	Ver string `json:"ver"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
