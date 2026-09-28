package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Placement object represents the properties of a placement and the characteristics of ads permitted to be rendering within them (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_placement
type Placement struct {
	// Identifier for specific ad placement or ad tag; unique within a distribution channel.
	TagID string `json:"tagid,omitempty"`
	// Indicates if server-side ad insertion (e.g., stitching an ad into an audio or video stream) is in use and the impact of this on asset and tracker retrieval, where 0 = status unknown, 1 = all client-side (i.e., not server-side), 2 = assets stitched server-side but tracking pixels fired client-side, 3 = all server-side.
	SSAI int8 `json:"ssai,omitempty"`
	// Name of ad mediation partner, SDK technology, or player responsible for rendering ad (typically video, audio, or mobile); used by some ad servers to customize ad code by partner.
	SDK string `json:"sdk,omitempty"`
	// Version of the SDK specified in the sdk attribute.
	SDKVer string `json:"sdkver,omitempty"`
	// Indicates if this is a rewarded placement, where 0 = no, 1 = yes.
	Reward int8 `json:"reward,omitempty"`
	// Whitelist of permitted languages of the creative using ISO-639-1-alpha-2.
	WLang []string `json:"wlang,omitempty"`
	// Flag to indicate if the creative is secure (i.e., uses HTTPS for all assets and markup), where 0 = no, 1 = yes.
	Secure int8 `json:"secure,omitempty"`
	// Indicates if including markup is supported (i.e., the various adm attributes throughout the Placement structure), where 0 = no, 1 = yes.
	AdMX int8 `json:"admx,omitempty"`
	// Indicates if retrieving markup via URL reference is supported (i.e., the various curl attributes throughout the Placement structure), where 0 = no, 1 = yes.
	CURLX int8 `json:"curlx,omitempty"`
	// Placement Subtype Object that indicates that this may be a display placement and provides additional detail related thereto.
	Display *DisplayPlacement `json:"display,omitempty"`
	// Placement Subtype Object that indicates that this may be a video placement and provides additional detail related thereto.
	Video *VideoPlacement `json:"video,omitempty"`
	// Placement Subtype Object that indicates that this may be an audio placement and provides additional detail related thereto.
	Audio *AudioPlacement `json:"audio,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
