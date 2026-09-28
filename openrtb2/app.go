package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object should be included if the ad supported content is a non-browser application (typically in mobile) as opposed to a website (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectapp
type App struct {
	// Exchange-specific app ID.
	ID string `json:"id,omitempty"`
	// App name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`
	// The store ID of the app in an app store.
	Bundle string `json:"bundle,omitempty"`
	// Domain of the app (e.g., “mygame.foo.com”).
	Domain string `json:"domain,omitempty"`
	// App store URL for an installed app; for IQG 2.1 compliance.
	StoreURL string `json:"storeurl,omitempty"`
	// The taxonomy in use.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// Array of IAB content categories of the app.
	Cat []string `json:"cat,omitempty"`
	// Array of IAB content categories that describe the current section of the app.
	SectionCat []string `json:"sectioncat,omitempty"`
	// Array of IAB content categories that describe the current page or view of the app.
	PageCat []string `json:"pagecat,omitempty"`
	// Application version.
	Ver string `json:"ver,omitempty"`
	// Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy *int8 `json:"privacypolicy,omitempty"`
	// 0 = app is free, 1 = the app is a paid version.
	Paid *int8 `json:"paid,omitempty"`
	// Details about the Publisher (Section 3.2.15) of the app.
	Publisher *Publisher `json:"publisher,omitempty"`
	// Details about the Content (Section 3.2.16) within the app
	Content *Content `json:"content,omitempty"`
	// Comma separated list of keywords about the app.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// A domain to be used for inventory authorization in the case of inventory sharing arrangements between an app owner and content owner.
	InventoryPartnerDomain string `json:"inventorypartnerdomain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
