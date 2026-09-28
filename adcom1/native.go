package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Native object is the root of a structure that defines a native display ad (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_native
type Native struct {
	// Default destination link for the native ad overall; used if an asset is activated (e.g., clicked) that doesn't specify it's own destination link.
	Link *LinkAsset `json:"link,omitempty"`
	// Array of assets that comprise the native ad.
	Asset []Asset `json:"asset,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
