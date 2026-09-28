package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object should be included if the ad supported content is a website as opposed to a non-browser application or Digital Out of Home (DOOH) inventory (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectsite
type Site struct {
	// Exchange-specific site ID.
	ID string `json:"id,omitempty"`
	// Site name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`
	// Domain of the site (e.g., “mysite.foo.com”).
	Domain string `json:"domain,omitempty"`
	// The taxonomy in use.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// Array of IABTL content categories of the site.
	Cat []string `json:"cat,omitempty"`
	// Array of IABTL content categories that describe the current section of the site.
	SectionCat []string `json:"sectioncat,omitempty"`
	// Array of IABTL content categories that describe the current page or view of the site.
	PageCat []string `json:"pagecat,omitempty"`
	// URL of the page where the impression will be shown.
	Page string `json:"page,omitempty"`
	// Referrer URL that caused navigation to the current page
	Ref string `json:"ref,omitempty"`
	// Search string that caused navigation to the current page.
	Search string `json:"search,omitempty"`
	// Indicates if the site has been programmed to optimize layout when viewed on mobile devices, where 0 = no, 1 = yes.
	Mobile *int8 `json:"mobile,omitempty"`
	// Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy *int8 `json:"privacypolicy,omitempty"`
	// Details about the Publisher (Section 3.2.15) of the site.
	Publisher *Publisher `json:"publisher,omitempty"`
	// Details about the Content (Section 3.2.16) within the site.
	Content *Content `json:"content,omitempty"`
	// Comma separated list of keywords about the site.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// A domain to be used for inventory authorization in the case of inventory sharing arrangements between a site owner and content owner.
	InventoryPartnerDomain string `json:"inventorypartnerdomain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
