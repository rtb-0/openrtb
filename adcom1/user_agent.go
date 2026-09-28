package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// UserAgent represents Structured user agent information provided when client supports User-Agent Client Hints (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_useragent
type UserAgent struct {
	// Each BrandVersion object identifies a browser or similar software component.
	Browsers []BrandVersion `json:"browsers,omitempty"`
	// Refer to Object: BrandVersion that identifies the user agent’s execution platform / OS.
	Platform *BrandVersion `json:"platform,omitempty"`
	// 1 if the agent prefers a “mobile” version of the content, if available, i.e. optimized for small screens or touch input.
	Mobile int8 `json:"mobile,omitempty"`
	// Device’s major binary architecture, e.g. “x86” or “arm”.
	Architecture string `json:"architecture,omitempty"`
	// Device’s bitness, e.g. “64” for 64-bit architecture.
	Bitness string `json:"bitness,omitempty"`
	// Device model.
	Model string `json:"model,omitempty"`
	// The source of data used to create this object.
	Source UserAgentSource `json:"source,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
