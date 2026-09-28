package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object contains a single user identifier provided as part of extended identifiers (OpenRTB 2.6 §3.2.28).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectuid
type UID struct {
	// The identifier for the user.
	ID string `json:"id,omitempty"`
	// Type of user agent the ID is from.
	AType adcom1.AgentType `json:"atype,omitempty"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
