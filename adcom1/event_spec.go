package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// EventSpec object specifies a type of ad tracking event and which methods of tracking are available for it (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_eventspec
type EventSpec struct {
	// Type of supported ad tracking event.
	Type EventType `json:"type,omitempty"`
	// Array of supported event tracking methods for this event type.
	Method []EventTrackingMethod `json:"method,omitempty"`
	// Event tracking APIs available for use; only relevant for JavaScript method trackers.
	API []APIFramework `json:"api,omitempty"`
	// Array of domains, top two levels only (e.g., “tracker.com”), that constitute a restriction list of JavaScript trackers.
	JSTrk []string `json:"jstrk,omitempty"`
	// Sense of the jstrk restriction list, where 0 = block list, 1 = whitelist.
	WJS int8 `json:"wjs,omitempty"`
	// Array of domains, top two levels only (e.g., “tracker.com”), that constitute a restriction list of pixel image trackers.
	PxTrk []string `json:"pxtrk,omitempty"`
	// Sense of the pxtrk restriction list, where 0 = block list, 1 = whitelist.
	WPx int8 `json:"wpx,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
