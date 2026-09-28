package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object contains any legal, governmental, or industry regulations that the sender deems applicable to the request (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectregs
type Regs struct {
	// Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	COPPA int8 `json:"coppa,omitempty"`
	// Flag that indicates whether or not the request is subject to GDPR regulations 0 = No, 1 = Yes, omission indicates Unknown.
	GDPR *int8 `json:"gdpr,omitempty"`
	// Communicates signals regarding consumer privacy under US privacy regulation.
	USPrivacy string `json:"us_privacy,omitempty"`
	// Contains the Global Privacy Platform’s consent string.
	GPP string `json:"gpp,omitempty"`
	// Array of the section(s) of the string which should be applied for this transaction.
	GPPSID []int8 `json:"gpp_sid,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
