package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// AssetFormat object represents the permitted specifications of a single asset of a native ad (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_assetformat
type AssetFormat struct {
	// Asset ID, unique within the scope of this placement specification.
	ID int64 `json:"id"`
	// Indicator of whether or not this asset is required, where 0 = no, 1 = yes.
	Req int8 `json:"req,omitempty"`
	// Asset Format Subtype Object that indicates this is specifying a title asset and provides additional detail as such.
	Title *TitleAssetFormat `json:"title,omitempty"`
	// Asset Format Subtype Object that indicates this is specifying an image asset and provides additional detail as such.
	Img *ImageAssetFormat `json:"img,omitempty"`
	// Asset Format Subtype Object, which leverages the VideoPlacement object, that indicates this is specifying a video asset and provides additional detail as such.
	Video *VideoPlacement `json:"video,omitempty"`
	// Asset Format Subtype Object that indicates this is specifying a data asset and provides additional detail as such.
	Data *DataAssetFormat `json:"data,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
