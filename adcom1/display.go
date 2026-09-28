package adcom1

import (
	"github.com/rtb-0/openrtb/common"
	nresp "github.com/rtb-0/openrtb/native1/response"
)

// Display object provides additional detail about an ad specifically for display ads (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_display
type Display[A any] struct {
	// Mime type of the ad (e.g., “image/jpeg”).
	MIME string `json:"mime,omitempty"`
	// API required by the ad if applicable.
	API []APIFramework `json:"api,omitempty"`
	// Subtype of display creative.
	CType DisplayCreativeSubtype `json:"ctype,omitempty"`
	// Absolute width of the creative in device independent pixels (DIPS), typically for non-native ads.
	W int64 `json:"w,omitempty"`
	// Absolute height of the creative in device independent pixels (DIPS), typically for non-native ads.
	H int64 `json:"h,omitempty"`
	// Relative width of the creative when expressing size as a ratio, typically for non-native ads.
	WRatio int8 `json:"wratio,omitempty"`
	// Relative height of the creative when expressing size as a ratio, typically for non-native ads.
	HRatio int8 `json:"hratio,omitempty"`
	// URL of a page informing the user about a buyer's targeting activity.
	Priv string `json:"priv,omitempty"`
	// General display markup (e.g., HTML, AMPHTML) if not using a structured alternative (e.g., banner, native).
	// JSON stores this as a string. A parses that string.
	Adm common.StringEncoded[A] `json:"adm,omitzero"`
	// Optional means of retrieving display markup by reference; a URL that can return HTML, AMPHTML, or a collection native Asset object and their subordinates).
	CURL string `json:"curl,omitempty"`
	// Structured banner image object, recommended for simple banner creatives.
	Banner *Banner `json:"banner,omitempty"`
	// Structured native object, recommended for native ads.
	// This is an OpenRTB Native 1.2 response object, not a JSON string.
	Native *nresp.Response `json:"native,omitempty"`
	// Array of events that the advertiser or buying platform wants to track.
	Event []Event `json:"event,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
