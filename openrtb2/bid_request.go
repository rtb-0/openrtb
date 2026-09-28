package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// The top-level bid request object contains an exchange unique bid request or auction ID (OpenRTB 2.6 §3.2.1).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbidrequest
type BidRequest struct {
	// ID of the bid request, assigned by the exchange, and unique for the exchange’s subsequent tracking of the responses.
	ID string `json:"id"`
	// Array of Imp objects (Section 3.2.4) representing the impressions offered.
	Imp []Imp `json:"imp"`
	// Details via a Site object (Section 3.2.13) about the publisher’s website.
	Site *Site `json:"site,omitempty"`
	// Details via an App object (Section 3.2.14) about the publisher’s app (i.e., non-browser applications).
	App *App `json:"app,omitempty"`
	// This object should be included if the ad supported content is a Digital Out-Of-Home screen.
	DOOH *DOOH `json:"dooh,omitempty"`
	// Details via a Device object (Section 3.2.18) about the user’s device to which the impression will be delivered.
	Device *Device `json:"device,omitempty"`
	// Details via a User object (Section 3.2.20) about the human user of the device; the advertising audience.
	User *User `json:"user,omitempty"`
	// Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	Test int8 `json:"test,omitempty"`
	// Auction type, where 1 = First Price, 2 = Second Price Plus.
	// Values of 500 and greater are exchange-specific.
	AT AuctionType `json:"at,omitempty"`
	// Maximum time in milliseconds the exchange allows for bids to be received including Internet latency to avoid timeout.
	TMax int64 `json:"tmax,omitempty"`
	// Allowed list of buyer seats (e.g., advertisers, agencies) allowed to bid on this impression.
	WSeat []string `json:"wseat,omitempty"`
	// Block list of buyer seats (e.g., advertisers, agencies) restricted from bidding on this impression.
	BSeat []string `json:"bseat,omitempty"`
	// Flag to indicate if Exchange can verify that the impressions offered represent all of the impressions available in context (e.g., all on the web page, all video spots such as pre/mid/post roll) to support road-blocking.
	AllImps int8 `json:"allimps,omitempty"`
	// Array of allowed currencies for bids on this bid request using ISO-4217 alpha codes.
	Cur []string `json:"cur,omitempty"`
	// Allowed list of languages for creatives using ISO-639-1-alpha-2.
	WLang []string `json:"wlang,omitempty"`
	// Allowed list of languages for creatives using IETF BCP 47I.
	WLangB []string `json:"wlangb,omitempty"`
	// Allowed advertiser categories using the specified category taxonomy.
	ACat []string `json:"acat,omitempty"`
	// Blocked advertiser categories using the specified category taxonomy.
	BCat []string `json:"bcat,omitempty"`
	// The taxonomy in use for bcat.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// Block list of advertisers by their domains (e.g., “ford.com”).
	BAdv []string `json:"badv,omitempty"`
	// Block list of applications by their app store IDs.
	BApp []string `json:"bapp,omitempty"`
	// A Sorce object (Section 3.2.2) that provides data about the inventory source and which entity makes the final decision.
	Source *Source `json:"source,omitempty"`
	// A Regs object (Section 3.2.3) that specifies any industry, legal, or governmental regulations in force for this request.
	Regs *Regs `json:"regs,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
