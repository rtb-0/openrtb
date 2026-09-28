package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object describes the network an ad will be displayed on (OpenRTB 2.6 §3.2.23).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectnetwork
type Network struct {
	// A unique identifier assigned by the publisher.
	ID string `json:"id,omitempty"`
	// Network the content is on (e.g., a TV network like “ABC").
	Name string `json:"name,omitempty"`
	// The primary domain of the network (e.g. "abc.com" in the case of the network ABC).
	Domain string `json:"domain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
