package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// ImageAsset object identifies the native asset as a image asset (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_imageasset
type ImageAsset struct {
	// A URL that returns the image for the asset.
	URL string `json:"url,omitempty"`
	// Width of the image asset in device independent pixels (DIPS).
	W int64 `json:"w,omitempty"`
	// Height of the image asset in device independent pixels (DIPS).
	H int64 `json:"h,omitempty"`
	// The type of image represented by this asset.
	Type NativeImageAssetType `json:"type,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
