// Package native1 holds the OpenRTB Native 1.2 enumerations.
package native1

import "github.com/rtb-0/openrtb/common"

// Info describes this package.
// The name Protocol is the creative-protocol enumeration in this package.
var Info = common.Protocol{
	Name:    "OpenRTB Native",
	Version: "1.2",
	Docs: []common.Doc{
		{Title: "OpenRTB Native Ads 1.2", URL: "https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md"},
	},
}
