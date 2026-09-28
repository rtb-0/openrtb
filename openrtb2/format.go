package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object represents an allowed size (i.e., height and width combination) or Flex Ad parameters for a banner impression (OpenRTB 2.6 §3.2.10).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectformat
type Format struct {
	// Width in device independent pixels (DIPS).
	W int64 `json:"w,omitempty"`
	// Height in device independent pixels (DIPS).
	H int64 `json:"h,omitempty"`
	// Relative width when expressing size as a ratio
	WRatio int64 `json:"wratio,omitempty"`
	// Relative height when expressing size as a ratio.
	HRatio int64 `json:"hratio,omitempty"`
	// The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
	WMin int64 `json:"wmin,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
