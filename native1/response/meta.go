// Package response holds OpenRTB Native 1.2 response objects.
package response

import "github.com/rtb-0/openrtb/common"

// Protocol describes this package.
var Protocol = common.Protocol{
	Name:    "OpenRTB Native",
	Version: "1.2",
	Docs: []common.Doc{
		{Title: "OpenRTB Native Ads 1.2", URL: "https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md"},
	},
}
