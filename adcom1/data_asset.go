package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// DataAsset object identifies the native asset as a data asset (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_dataasset
type DataAsset struct {
	// A formatted string of data to be displayed (e.g., “5 stars”, “3.4 stars out of 5”, “$10”, etc.).
	Value string `json:"value"`
	// The length of the value contents.
	Len int64 `json:"len,omitempty"`
	// The type of data represented by this asset.
	Type NativeDataAssetType `json:"type,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
