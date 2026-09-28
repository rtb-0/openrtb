package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Network describes the network an ad will be displayed on (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_network
type Network struct {
	// A unique identifier assigned by the publisher.
	ID string `json:"id,omitempty"`
	// Network the content is on (e.g., a TV network like "ABC").
	Name string `json:"name,omitempty"`
	// The primary domain of the network (e.g., “abc.com” in the case of the network ABC).
	Domain string `json:"domain,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
