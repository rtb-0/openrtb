package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// DisplayFormat object represents an allowed set of parameters for a banner display ad and often appears as an array when multiple sizes are permitted (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_displayformat
type DisplayFormat struct {
	// Absolute width of the creative in units specified by DisplayPlacement.unit.
	W int64 `json:"w,omitempty"`
	// Absolute height of the creative in units specified by DisplayPlacement.unit.
	H int64 `json:"h,omitempty"`
	// Relative width of the creative when expressing size as a ratio.
	WRatio int8 `json:"wratio,omitempty"`
	// Relative height of the creative when expressing size as a ratio.
	HRatio int8 `json:"hratio,omitempty"`
	// Directions in which the creative is permitted to expand.
	ExpDir []ExpandableDirection `json:"expdir,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
