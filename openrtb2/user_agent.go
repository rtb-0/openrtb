package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// Structured user agent information from User-Agent Client Hints (OpenRTB 2.6 §3.2.29).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectuseragent
type UserAgent struct {
	// Each BrandVersion object (see Section 3.2.30) identifies a browser or similar software component.
	Browsers []BrandVersion `json:"browsers,omitempty"`
	// A BrandVersion object (see Section 3.2.30) that identifies the user agent’s execution platform / OS.
	Platform *BrandVersion `json:"platform,omitempty"`
	// 1 if the agent prefers a “mobile” version of the content, if available, i.e. optimized for small screens or touch input.
	Mobile *int8 `json:"mobile,omitempty"`
	// Device’s major binary architecture, e.g. “x86” or “arm”.
	Architecture string `json:"architecture,omitempty"`
	// Device’s bitness, e.g. “64” for 64-bit architecture.
	Bitness string `json:"bitness,omitempty"`
	// Device model.
	Model string `json:"model,omitempty"`
	// The source of data used to create this object, List: User-Agent Source in AdCOM 1.0.
	Source adcom1.UserAgentSource `json:"source,omitempty"`
	// Placeholder for advertising-system specific extensions to this object.
	Ext common.Ext `json:"ext,omitempty"`
}
