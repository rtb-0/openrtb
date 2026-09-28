package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Publisher object describes the publisher of the media in which ads will be displayed (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_publisher
type Publisher struct {
	// Vendor-specific unique publisher identifier, as used in ads.txt files.
	ID string `json:"id,omitempty"`
	// Displayable name of the publisher.
	Name string `json:"name,omitempty"`
	// Highest level domain of the publisher (e.g., “publisher.com”).
	Domain string `json:"domain,omitempty"`
	// Array of content categories that describe the publisher using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`
	// The taxonomy in use for the cat attribute.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
