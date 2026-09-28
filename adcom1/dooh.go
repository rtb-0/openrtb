package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// DOOH object is used to define an ad supported digital out-of-home (DOOH) experience such as a public kiosk or digital billboard (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_dooh
type DOOH struct {
	DistributionChannel
	// The type of out-of-home venue.
	Venue DOOHVenueType `json:"venue,omitempty"`
	// Indicates whether the DOOH placement is in a fixed location (e.g., kiosk, billboard, elevator) or is movable (e.g., taxi), where 1 = fixed, 2 = movable.
	Fixed int8 `json:"fixed,omitempty"`
	// The exposure time in seconds per view that the creative will be displayed before refreshing to the next creative.
	ETime int64 `json:"etime,omitempty"`
	// Minimum DPI for text-based creative elements to display clearly.
	DPI int64 `json:"dpi,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
