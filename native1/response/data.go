package response

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// Corresponds to the Data Object in the request, with the value filled in (OpenRTB Native 1.2 §5.5).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-5
type Data struct {
	// Required for assetsurl/dcourl responses, not required for embedded asset responses.
	Type native1.DataAssetType `json:"type,omitempty"`
	// Required for assetsurl/dcourl responses, not required for embedded asset responses.
	Len int64 `json:"len,omitempty"`
	// The optional formatted string name of the data type to be displayed.
	Label string `json:"label,omitempty"`
	// The formatted string of data to be displayed.
	Value string `json:"value"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
