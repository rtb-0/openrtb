package response

import (
	"github.com/rtb-0/openrtb/common"
)

// Corresponds to the Asset Object in the request (OpenRTB Native 1.2 §5.2).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-2
type Asset struct {
	// Optional if assetsurl/dcourl is being used; required if embedded asset is being used.
	ID *int64 `json:"id,omitempty"`
	// Set to 1 if asset is required.
	Required int8 `json:"required,omitempty"`
	// Title object for title assets.
	Title *Title `json:"title,omitempty"`
	// Image object for image assets.
	Img *Image `json:"img,omitempty"`
	// Video object for video assets.
	Video *Video `json:"video,omitempty"`
	// Data object for ratings, prices etc. Asset object may contain only one of title, img, data or video.
	Data *Data `json:"data,omitempty"`
	// Link object for call to actions.
	Link *Link `json:"link,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification Bidders are encouraged not to use asset.ext for exchanging text assets.
	Ext common.Ext `json:"ext,omitempty"`
}
