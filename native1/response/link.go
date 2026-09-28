package response

import (
	"github.com/rtb-0/openrtb/common"
)

// Used for ‘call to action’ assets, or other links from the Native ad (OpenRTB Native 1.2 §5.7).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-7
type Link struct {
	// Landing URL of the clickable link.
	URL string `json:"url"`
	// List of third-party tracker URLs to be fired on click of the URL.
	ClickTrackers []string `json:"clicktrackers,omitempty"`
	// Fallback URL for deeplink.
	Fallback string `json:"fallback,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
