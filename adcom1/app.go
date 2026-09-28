package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// App object is used to define an ad supported non-browser application, in contrast to a typical website, example (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_app
type App struct {
	DistributionChannel
	// Domain of the app (e.g., “mygame.foo.com”).
	Domain string `json:"domain,omitempty"`
	// Array of content categories describing the app using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`
	// Array of content categories describing the current section of the app using IDs from the taxonomy indicated in cattax.
	SectCat []string `json:"sectcat,omitempty"`
	// Array of content categories describing the current page or view of the app using IDs from the taxonomy indicated in cattax.
	PageCat []string `json:"pagecat,omitempty"`
	// The taxonomy in use for the cat, sectcat and pagecat attributes.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`
	// Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	PrivPolicy int8 `json:"privpolicy,omitempty"`
	// Comma-separated list of keywords about the app.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// Bundle or package name of the app (e.g., “com.foo.mygame”) and should NOT be app store IDs (e.g., not iTunes store IDs).
	Bundle string `json:"bundle,omitempty"`
	// The ID of the app in an app store (e.g., Apple iTunes, Google Play).
	StoreID string `json:"storeid,omitempty"`
	// App store URL for an installed app; for IQG 2.1 compliance.
	StoreURL string `json:"storeurl,omitempty"`
	// Application version.
	Ver string `json:"ver,omitempty"`
	// Indicator of whether or not this is a paid app, where 0 = free, 1 = paid.
	Paid int8 `json:"paid,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
