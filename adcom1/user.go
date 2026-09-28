package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// User object contains information known or derived about the human user of the device (i.e., the audience for advertising) (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_user
type User struct {
	// Vendor-specific ID for the user.
	ID string `json:"id,omitempty"`
	// Buyer-specific ID for the user as mapped by an exchange for the buyer.
	BuyerUID string `json:"buyeruid,omitempty"`
	// Year of birth as a 4-digit integer.
	YOB int64 `json:"yob,omitempty"`
	// Gender, where “M” = male, “F” = female, “O” = known to be other (i.e., omitted is unknown).
	Gender string `json:"gender,omitempty"`
	// Comma-separated list of keywords, interests, or intent.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// GDPR consent string if applicable, complying with the comply with the IAB standard Consent String Format in the Transparency and Consent Framework technical specifications.
	Consent string `json:"consent,omitempty"`
	// Location of the user's home base (i.e., not necessarily their current location).
	Geo *Geo `json:"geo,omitempty"`
	// Additional user data.
	Data []Data `json:"data,omitempty"`
	// Extended (third-party) identifiers for this user.
	EIDs []ExtendedIdentifier `json:"eids,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
