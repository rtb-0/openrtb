package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Event object specifies a type of event that the advertiser or buying platform wants to track along with the information required to do so (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_event
type Event struct {
	// Type of event to track.
	Type EventType `json:"type"`
	// Method of tracking requested.
	Method EventTrackingMethod `json:"method"`
	// The APIs being used by the tracker; only relevant when the tracking method is JavaScript.
	API []APIFramework `json:"api,omitempty"`
	// The URL of the tracking pixel or JavaScript tag, respectively.
	URL string `json:"url,omitempty"`
	// An array of key-value pairs to support vendor-specific data required for custom tracking.
	CData map[string]string `json:"cdata,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
