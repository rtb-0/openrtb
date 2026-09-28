package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// ImageAssetFormat object is used to provide native asset format specifications for an image element (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_imageassetformat
type ImageAssetFormat struct {
	// The type of image asset supported.
	Type NativeImageAssetType `json:"type,omitempty"`
	// Array of supported mime types (e.g., “image/jpeg”, “image/gif”).
	MIME []string `json:"mime,omitempty"`
	// Absolute width of the image asset in device independent pixels (DIPS).
	W int64 `json:"w,omitempty"`
	// Absolute height of the image asset in device independent pixels (DIPS).
	H int64 `json:"h,omitempty"`
	// The minimum requested absolute width of the image in device independent pixels (DIPS).
	WMin int64 `json:"wmin,omitempty"`
	// The minimum requested absolute height of the image in device independent pixels (DIPS).
	HMin int64 `json:"hmin,omitempty"`
	// Relative width of the image asset when expressing size as a ratio.
	WRatio int8 `json:"wratio,omitempty"`
	// Relative height of the image asset when expressing size as a ratio.
	HRatio int8 `json:"hratio,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
