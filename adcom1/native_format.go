package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// NativeFormat object refines a display placement to be specifically a native display placement (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_nativeformat
type NativeFormat struct {
	// Array of objects that specify the set of native assets and their permitted formats.
	Asset []AssetFormat `json:"asset"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
