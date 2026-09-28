package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// TitleAsset object identifies the native asset as a title asset, which is essentially just a plain text string with specified length (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_titleasset
type TitleAsset struct {
	// The text content of the text element.
	Text string `json:"text,omitempty"`
	// The length of the contents of the text attribute.
	Len int64 `json:"len,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
