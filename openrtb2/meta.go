// Package openrtb2 implements OpenRTB 2.6, including the 2.5 object model.
package openrtb2

import "github.com/rtb-0/openrtb/common"

// Protocol describes this package.
var Protocol = common.Protocol{
	Name:    "OpenRTB",
	Version: "2.6",
	Docs: []common.Doc{
		{Title: "OpenRTB 2.6", URL: "https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md"},
		{Title: "OpenRTB 2.5", URL: "https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.5.md"},
	},
}
