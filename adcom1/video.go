package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Video object provides additional detail about an ad specifically for video ads (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_video
type Video[A any] struct {
	// Mime type(s) of the ad creative(s) (e.g., “video/mp4”).
	MIME []string `json:"mime,omitempty"`
	// API required by the ad if applicable.
	API []APIFramework `json:"api,omitempty"`
	// Subtype of video creative.
	CType MediaCreativeSubtype `json:"ctype,omitempty"`
	// Duration of the video creative in seconds.
	Dur int64 `json:"dur,omitempty"`
	// Video markup (e.g., VAST). JSON stores this as a string. A parses that string.
	Adm common.StringEncoded[A] `json:"adm,omitzero"`
	// Optional means of retrieving markup by reference; a URL that returns video markup (e.g., VAST).
	CURL string `json:"curl,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
