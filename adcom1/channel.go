package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Channel describes the channel an ad will be displayed on (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_channel
type Channel struct {
	// A unique identifier assigned by the publisher.
	ID string `json:"id,omitempty"`
	// Channel the content is on (e.g., a local channel like "WABC-TV").
	Name string `json:"name,omitempty"`
	// The primary domain of the channel (e.g., “abc7ny.com” in the case of the local channel WABC-TV).
	Domain string `json:"domain,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
