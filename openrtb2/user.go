package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object contains information known or derived about the human user of the device (i.e., the audience for advertising) (OpenRTB 2.6 §3.2.20).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectuser
type User struct {
	// Exchange-specific ID for the user.
	ID string `json:"id,omitempty"`
	// Buyer-specific ID for the user as mapped by the exchange for the buyer.
	BuyerUID string `json:"buyeruid,omitempty"`
	// Year of birth as a 4-digit integer.
	Yob int64 `json:"yob,omitempty"`
	// Gender, where “M” = male, “F” = female, “O” = known to be other (i.e., omitted is unknown).
	Gender string `json:"gender,omitempty"`
	// Comma separated list of keywords, interests, or intent.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// Optional feature to pass bidder data that was set in the exchange’s cookie.
	CustomData string `json:"customdata,omitempty"`
	// Location of the user’s home base defined by a Geo object (Section 3.2.19).
	Geo *Geo `json:"geo,omitempty"`
	// Additional user data.
	Data []Data `json:"data,omitempty"`
	// When GDPR regulations are in effect this attribute contains the Transparency and Consent Framework’s Consent String data structure.
	Consent string `json:"consent,omitempty"`
	// Details for support of a standard protocol for multiple third party identity providers (Section 3.2.27)
	EIDs []EID `json:"eids,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
