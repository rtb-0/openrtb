package response

import (
	"encoding/json"
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// The event trackers response is an array of objects and specifies the types of events the bidder wishes to track and the URLs/information to track them (OpenRTB Native 1.2 §5.8).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-8
type EventTracker struct {
	// Type of event to track.
	Event native1.EventType `json:"event"`
	// Type of tracking requested See Event Tracking Methods table.
	Method native1.EventTrackingMethod `json:"method"`
	// The URL of the image or js.
	URL string `json:"url,omitempty"`
	// To be agreed individually with the exchange, an array of key:value objects for custom tracking, for example the account number of the DSP with a tracking company.
	CustomData json.RawMessage `json:"customdata,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
