package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object describes the channel an ad will be displayed on (OpenRTB 2.6 §3.2.24).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectchannel
type Channel struct {
	// A unique identifier assigned by the publisher.
	ID string `json:"id,omitempty"`
	// Channel the content is on (e.g., a local channel like “WABC-TV").
	Name string `json:"name,omitempty"`
	// The primary domain of the channel (e.g. "abc7ny.com" in the case of the local channel WABC-TV).
	Domain string `json:"domain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
