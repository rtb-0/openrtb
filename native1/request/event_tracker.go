package request

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// The event trackers object specifies the types of events the bidder can request to be tracked in the bid response, and which types of tracking are available for each event type, and is included as an array in the request (OpenRTB Native 1.2 §4.7).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-7
type EventTracker struct {
	// Type of event available for tracking.
	Event native1.EventType `json:"event"`
	// Array of the types of tracking available for the given event.
	Methods []native1.EventTrackingMethod `json:"methods"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
