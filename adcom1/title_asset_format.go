package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// TitleAssetFormat object is used to provide native asset format specifications for a title element (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_titleassetformat
type TitleAssetFormat struct {
	// The maximum allowed length of the title value.
	Len int64 `json:"len,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
