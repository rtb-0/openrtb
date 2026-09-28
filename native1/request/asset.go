package request

import (
	"github.com/rtb-0/openrtb/common"
)

// The main container object for each asset requested or supported by Exchange on behalf of the rendering client (OpenRTB Native 1.2 §4.2).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-2
type Asset struct {
	// Unique asset ID, assigned by exchange.
	ID int64 `json:"id"`
	// Set to 1 if asset is required (exchange will not accept a bid without it)
	Required int8 `json:"required,omitempty"`
	// Title object for title assets.
	Title *Title `json:"title,omitempty"`
	// Image object for image assets.
	Img *Image `json:"img,omitempty"`
	// Video object for video assets.
	Video *Video `json:"video,omitempty"`
	// Data object for brand name, description, ratings, prices etc. See DataObject definition.
	Data *Data `json:"data,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
