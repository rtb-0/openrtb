package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// ExtendedIdentifierUID is defined by AdCOM 1.0 (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_eid_uids
type ExtendedIdentifierUID struct {
	// Cookie or platform-native identifier.
	ID string `json:"id,omitempty"`
	// Type of user agent the match is from.
	AType AgentType `json:"atype,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
