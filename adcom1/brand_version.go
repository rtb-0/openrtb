package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// BrandVersion provides further identification based on User-Agent Client Hints (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_brandversion
type BrandVersion struct {
	// A brand identifier, for example, “Chrome” or “Windows”.
	Brand string `json:"brand,omitempty"`
	// A sequence of version components, in descending hierarchical order (major, minor, micro, …).
	Version []string `json:"version,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
