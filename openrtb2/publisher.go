package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object describes the entity who directly supplies inventory to and is paid by the exchange (OpenRTB 2.6 §3.2.15).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectpublisher
type Publisher struct {
	// Exchange-specific publisher ID.
	ID string `json:"id,omitempty"`
	// Publisher name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`
	// The taxonomy in use.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// Array of IAB content categories that describe the publisher.
	Cat []string `json:"cat,omitempty"`
	// Highest level domain of the publisher (e.g., “publisher.com”).
	Domain string `json:"domain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
