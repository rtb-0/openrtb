package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object should be included if the ad supported content is a Digital Out-Of-Home screen (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectdooh
type DOOH struct {
	// Exchange provided id for a placement or logical grouping of placements.
	ID string `json:"id,omitempty"`
	// Name of the DOOH placement.
	Name string `json:"name,omitempty"`
	// The type of out-of-home venue.
	VenueType []string `json:"venuetype,omitempty"`
	// The venue taxonomy in use.
	VenueTypeTax *adcom1.DOOHVenueTaxonomy `json:"venuetypetax,omitempty"`
	// Details about the publisher of the placement.
	Publisher *Publisher `json:"publisher,omitempty"`
	// Domain of the inventory owner (e.g., “mysite.foo.com”)
	Domain string `json:"domain,omitempty"`
	// Comma separated list of keywords about the DOOH placement.
	Keywords string `json:"keywords,omitempty"`
	// Details about the Content within the DOOH placement.
	Content *Content `json:"content,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
