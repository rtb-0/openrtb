package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Ad object is the root of a structure that defines in instance of advertising media (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_ad
type Ad[A any] struct {
	// ID of the creative; unique at least throughout the scope of a vendor (e.g., an exchange or buying platform).
	ID string `json:"id"`
	// Advertiser domain; top two levels only (e.g., “ford.com”).
	ADomain []string `json:"adomain,omitempty"`
	// When the product of the ad is an app, the unique ID of that app as a bundle or package name (e.g., “com.foo.mygame”).
	Bundle []string `json:"bundle,omitempty"`
	// URL without cache-busting to an image that is representative of the ad content for cursory level ad quality checking.
	IURL string `json:"iurl,omitempty"`
	// Array of content categories describing the ad using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`
	// The taxonomy in use for the cat attribute.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`
	// Language of the creative using ISO-639-1-alpha-2.
	Lang string `json:"lang,omitempty"`
	// Set of attributes describing the creative.
	Attr []CreativeAttribute `json:"attr,omitempty"`
	// Flag to indicate if the creative is secure (i.e., uses HTTPS for all assets and markup), where 0 = no, 1 = yes.
	Secure int8 `json:"secure,omitempty"`
	// Media rating per IQG guidelines.
	MRating MediaRating `json:"mrating,omitempty"`
	// Timestamp of the original instantiation of this ad (i.e., this object or any of its children) in Unix format (i.e., milliseconds since the epoch).
	Init int64 `json:"init,omitempty"`
	// Timestamp of most recent modification to this ad (i.e., this object or any of its children other than the Audit object) in Unix format (i.e., milliseconds since the epoch).
	LastMod int64 `json:"lastmod,omitempty"`
	// Media Subtype Object that indicates this is a display ad and provides additional detail as such.
	Display *Display[A] `json:"display,omitempty"`
	// Media Subtype Object that indicates this is a video ad and provides additional detail as such.
	Video *Video[A] `json:"video,omitempty"`
	// Media Subtype Object that indicates this is an audio ad and provides additional detail as such.
	Audio *Audio[A] `json:"audio,omitempty"`
	// An object depicting the audit status of the ad; typically part of a quality/safety review process.
	Audit *Audit[A] `json:"audit,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
