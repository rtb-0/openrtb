package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Producer object defines the producer of the content in which ad will be displayed (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_producer
type Producer struct {
	// Vendor-specific unique producer identifier.
	ID string `json:"id,omitempty"`
	// Displayable name of the producer.
	Name string `json:"name,omitempty"`
	// Highest level domain of the producer (e.g., “producer.com”).
	Domain string `json:"domain,omitempty"`
	// Array of content categories that describe the producer using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`
	// The taxonomy in use for the cat attribute.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
