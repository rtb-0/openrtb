package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// VideoAsset object identifies the native asset as a video asset (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_videoasset
type VideoAsset struct {
	// Video markup (e.g., VAST document) for the asset.
	AdM string `json:"adm,omitempty"`
	// A URL that returns the video markup (e.g., VAST document) for the asset.
	CURL string `json:"curl,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
