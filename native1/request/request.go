package request

import (
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/native1"
)

// The Native Object defines the native1 advertising opportunity available for bid via this bid request (OpenRTB Native 1.2 §4.1).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-1
type Request struct {
	// Version of the Native Markup version in use.
	Ver string `json:"ver,omitempty"`
	// The Layout ID of the native1 ad unit.
	Layout native1.Layout `json:"layout,omitempty"`
	// The Ad unit ID of the native1 ad unit.
	AdUnit native1.AdUnit `json:"adunit,omitempty"`
	// The context in which the ad appears.
	Context native1.ContextType `json:"context,omitempty"`
	// A more detailed context in which the ad appears.
	ContextSubType native1.ContextSubType `json:"contextsubtype,omitempty"`
	// The design/format/layout of the ad unit being offered.
	PlcmtType native1.PlacementType `json:"plcmttype,omitempty"`
	// The number of identical placements in this Layout.
	PlcmtCnt int64 `json:"plcmtcnt,omitempty"`
	// 0 for the first ad, 1 for the second ad, and so on.
	Seq int64 `json:"seq,omitempty"`
	// An array of Asset Objects.
	Assets []Asset `json:"assets"`
	// Whether the supply source / impression supports returning an assetsurl instead of an asset object.
	AURLSupport int8 `json:"aurlsupport,omitempty"`
	// Whether the supply source / impression supports returning a dco url instead of an asset object.
	DURLSupport int8 `json:"durlsupport,omitempty"`
	// Specifies what type of event tracking is supported - see Event Trackers Request Object
	EventTrackers []EventTracker `json:"eventtrackers,omitempty"`
	// Set to 1 when the native1 ad supports buyer-specific privacy notice.
	Privacy int8 `json:"privacy,omitempty"`
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext common.Ext `json:"ext,omitempty"`
}
