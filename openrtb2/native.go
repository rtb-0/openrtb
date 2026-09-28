package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
	nreq "github.com/rtb-0/openrtb/native1/request"
)

// This object represents a native type impression (OpenRTB 2.6 §3.2.9).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectnative
type Native struct {
	// Request payload complying with the Native Ad Specification.
	// JSON stores this as a string containing the native request object.
	Request common.StringEncoded[nreq.Request] `json:"request"`
	// Version of the Dynamic Native Ads API to which request complies; highly recommended for efficient parsing.
	Ver string `json:"ver,omitempty"`
	// List of supported API frameworks for this impression.
	API []adcom1.APIFramework `json:"api,omitempty"`
	// Blocked creative attributes.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
