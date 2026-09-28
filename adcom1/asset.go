package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Asset object is the container for each asset comprising a native ad (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_asset
type Asset struct {
	// The value of AssetFormat.id if this ad references a specific native placement defined by a Placement object and its structure.
	ID int64 `json:"id,omitempty"`
	// Indicates if the asset is required to be displayed, where 0 = no, 1 = yes.
	Req int8 `json:"req,omitempty"`
	// Asset Subtype Object that indicates this is a title asset and provides additional detail as such.
	Title *TitleAsset `json:"title,omitempty"`
	// Asset Subtype Object that indicates this is an image asset and provides additional detail as such.
	Image *ImageAsset `json:"image,omitempty"`
	// Asset Subtype Object that indicates this is a video asset and provides additional detail as such.
	Video *VideoAsset `json:"video,omitempty"`
	// Asset Subtype Object that indicates this is a data asset and provides additional detail as such.
	Data *DataAsset `json:"data,omitempty"`
	// Asset Subtype Object that indicates this is a link asset and provides additional detail as such.
	Link *LinkAsset `json:"link,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
