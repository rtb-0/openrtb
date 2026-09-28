package response

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// Corresponds to the Image Object in the request (OpenRTB Native 1.2 §5.4).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-4
type Image struct {
	// Required for assetsurl or dcourl responses, not required for embedded asset responses.
	Type native1.ImageAssetType `json:"type,omitempty"`
	// URL of the image asset
	URL string `json:"url"`
	// Width of the image in pixels.
	W int64 `json:"w,omitempty"`
	// Height of the image in pixels.
	H int64 `json:"h,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
