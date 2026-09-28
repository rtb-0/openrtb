package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Audio object provides additional detail about an ad specifically for audio ads (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_audio
type Audio[A any] struct {
	// Mime type(s) of the ad creative(s) (e.g., “audio/mp4”).
	MIME []string `json:"mime,omitempty"`
	// API required by the ad if applicable.
	API []APIFramework `json:"api,omitempty"`
	// Subtype of audio creative.
	CType MediaCreativeSubtype `json:"ctype,omitempty"`
	// Duration of the audio creative in seconds.
	Dur int64 `json:"dur,omitempty"`
	// Audio markup (e.g., DAAST). JSON stores this as a string. A parses that string.
	Adm common.StringEncoded[A] `json:"adm,omitzero"`
	// Optional means of retrieving markup by reference; a URL that returns audio markup (e.g., DAAST).
	CURL string `json:"curl,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
