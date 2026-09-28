package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Site object is used to define an ad supported website, in contrast to a non-browser application, for example (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_site
type Site struct {
	DistributionChannel
	// Domain of the site (e.g., “mysite.foo.com”).
	Domain string `json:"domain,omitempty"`
	// Array of content categories describing the site using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`
	// Array of content categories describing the current section of the site using IDs from the taxonomy indicated in cattax.
	SectCat []string `json:"sectcat,omitempty"`
	// Array of content categories describing the current page or view of the site using IDs from the taxonomy indicated in cattax.
	PageCat []string `json:"pagecat,omitempty"`
	// The taxonomy in use for the cat, sectcat and pagecat attributes.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`
	// Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	PrivPolicy int8 `json:"privpolicy,omitempty"`
	// Comma-separated list of keywords about the app.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// URL of the page within the site.
	Page string `json:"page,omitempty"`
	// Referrer URL that caused navigation to the current page.
	Ref string `json:"ref,omitempty"`
	// Search string that caused navigation to the current page.
	Search string `json:"search,omitempty"`
	// Indicates if the site has been programmed to optimize layout when viewed on mobile devices, where 0 = no, 1 = yes.
	Mobile int8 `json:"mobile,omitempty"`
	// Indicates if the page is built with AMP HTML, where 0 = no, 1 = yes.
	AMP int8 `json:"amp,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
