package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Banner object describes a basic banner creative (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_banner
type Banner struct {
	// A URL that will return the image.
	Img string `json:"img"`
	// Destination link if the image is activated (e.g., clicked); not applicable in some contexts (e.g., DOOH) and its inclusion does not guarantee it will be supported.
	Link *LinkAsset `json:"link,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
