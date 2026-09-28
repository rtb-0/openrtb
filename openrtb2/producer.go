package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object defines the producer of the content in which the ad will be shown (OpenRTB 2.6 §3.2.17).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectproducer
type Producer struct {
	// Content producer or originator ID.
	ID string `json:"id,omitempty"`
	// Content producer or originator name (e.g., “Warner Bros”).
	Name string `json:"name,omitempty"`
	// The taxonomy in use.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// Array of IAB content categories that describe the content producer.
	Cat []string `json:"cat,omitempty"`
	// Highest level domain of the content producer (e.g., “producer.com”).
	Domain string `json:"domain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
