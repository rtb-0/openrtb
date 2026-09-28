package response

import (
	"github.com/rtb-0/openrtb/common"
)

// The native1 object is the top level JSON object which identifies a native response (OpenRTB Native 1.2 §5.1).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-1
type Response struct {
	// Version of the Native Markup version in use.
	Ver string `json:"ver,omitempty"`
	// List of native1 ad’s assets.
	Assets []Asset `json:"assets,omitempty"`
	// URL of an alternate source for the assets object.
	AssetsURL string `json:"assetsurl,omitempty"`
	// URL where a dynamic creative specification may be found for populating this ad, per the Dynamic Content Ads Specification.
	DCOURL string `json:"dcourl,omitempty"`
	// Destination Link.
	Link Link `json:"link"`
	// Array of impression tracking URLs, expected to return a 1x1 image or 204 response - typically only passed when using 3rd party trackers.
	ImpTrackers []string `json:"imptrackers,omitempty"`
	// Optional JavaScript impression tracker.
	JSTracker string `json:"jstracker,omitempty"`
	// Array of tracking objects to run with the ad, in response to the declared supported methods in the request.
	EventTrackers []EventTracker `json:"eventtrackers,omitempty"`
	// If support was indicated in the request, URL of a page informing the user about the buyer’s targeting activity.
	Privacy string `json:"privacy,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
