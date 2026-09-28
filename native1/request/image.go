package request

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// The Image object to be used for all image elements of the Native ad such as Icons, Main Image, etc. Recommended sizes and aspect ratios are included in the Image Asset Types section (OpenRTB Native 1.2 §4.4).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-4
type Image struct {
	// Type ID of the image element supported by the publisher.
	Type native1.ImageAssetType `json:"type,omitempty"`
	// Width of the image in pixels.
	W int64 `json:"w,omitempty"`
	// The minimum requested width of the image in pixels.
	WMin int64 `json:"wmin,omitempty"`
	// Height of the image in pixels.
	H int64 `json:"h,omitempty"`
	// The minimum requested height of the image in pixels.
	HMin int64 `json:"hmin,omitempty"`
	// Whitelist of content MIME types supported.
	MIMEs []string `json:"mimes,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
