package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// Extended identifiers support in the OpenRTB specification allows buyers to use audience data in real-time bidding (OpenRTB 2.6 §3.2.27).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objecteid
type EID struct {
	// The canonical domain name of the entity (publisher, publisher monetization company, SSP, Exchange, Header Wrapper, etc.) that caused the ID array element to be added.
	Inserter string `json:"inserter,omitempty"`
	// Technology providing the match method as defined in mm.
	Matcher string `json:"matcher,omitempty"`
	// Match method used by the matcher.
	MM adcom1.MatchMethod `json:"mm,omitempty"`
	// Canonical domain of the ID.
	Source string `json:"source,omitempty"`
	// Array of extended ID UID objects from the given source.
	UIDs []UID `json:"uids,omitempty"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
