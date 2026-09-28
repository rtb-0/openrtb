package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Companion object is used in video and audio placements to specify an associated or so-called companion display ad (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_companion
type Companion struct {
	// Identifier of the companion ad; unique within this placement.
	ID string `json:"id,omitempty"`
	// Indicates the companion ad rendering mode relative to the associated video or audio ad, where 0 = concurrent, 1 = end-card.
	VCm int8 `json:"vcm,omitempty"`
	// Display specification object representing the companion ad.
	Display *DisplayPlacement `json:"display,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
