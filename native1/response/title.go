package response

import (
	"github.com/rtb-0/openrtb/common"
)

// Corresponds to the Title Object in the request, with the value filled in (OpenRTB Native 1.2 §5.3).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-3
type Title struct {
	// The text associated with the text element.
	Text string `json:"text"`
	// The length of the title being provided.
	Len int64 `json:"len,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
