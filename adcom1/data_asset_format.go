package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// DataAssetFormat object is used to provide native asset format specifications for a data element (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_dataassetformat
type DataAssetFormat struct {
	// The type of data asset supported.
	Type NativeDataAssetType `json:"type,omitempty"`
	// The maximum allowed length of the data value.
	Len int64 `json:"len,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
