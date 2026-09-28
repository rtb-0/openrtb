package request

import (
	"github.com/rtb-0/openrtb/common"
)

// The Title object is to be used for title element of the Native ad (OpenRTB Native 1.2 §4.3).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-3
type Title struct {
	// Maximum length of the text in the title element.
	Len int64 `json:"len"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
