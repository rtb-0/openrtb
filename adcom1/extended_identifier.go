package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// ExtendedIdentifier support in the OpenRTB specification allows buyers to use audience data in real-time bidding (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_eids
type ExtendedIdentifier struct {
	// Source or technology provider responsible for the set of included IDs.
	Source string `json:"source,omitempty"`
	// Array of extended ID UID objects from the given source.
	UIDs []ExtendedIdentifierUID `json:"uids,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
