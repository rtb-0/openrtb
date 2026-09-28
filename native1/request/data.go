package request

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// The Data Object is to be used for all non-core elements of the native1 unit such as Brand Name, Ratings, Review Count, Stars, Download count, descriptions etc. It is also generic for future native1 elements not contemplated at the time of the writing of this document (OpenRTB Native 1.2 §4.6).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-6
type Data struct {
	// Type ID of the element supported by the publisher.
	Type native1.DataAssetType `json:"type"`
	// Maximum length of the text in the element’s response.
	Len int64 `json:"len,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
