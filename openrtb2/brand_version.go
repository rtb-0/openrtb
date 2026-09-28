package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// Further identification based on User-Agent Client Hints, the BrandVersion object is used to identify a device’s browser or similar software component, and the user agent’s execution platform or operating system (OpenRTB 2.6 §3.2.30).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbrandversion
type BrandVersion struct {
	// A brand identifier, for example, "Chrome" or "Windows".
	Brand string `json:"brand"`
	// A sequence of version components, in descending hierarchical order (major, minor, micro, …)
	Version []string `json:"version,omitempty"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
