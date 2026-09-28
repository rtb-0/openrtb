package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Regs object contains any known legal, governmental, or industry regulations that are in effect (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_regs
type Regs struct {
	// Flag indicating if COPPA regulations apply, where 0 = no, 1 = yes.
	COPPA int8 `json:"coppa,omitempty"`
	// Flag indicating if GDPR regulations apply, where 0 = no, 1 = yes.
	GDPR int8 `json:"gdpr,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
