package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// LinkAsset object identifies the native asset as a link asset and is used to define navigation for call-to-action or other activations (i.e., clicks) A link asset can be independent or associated with the overall native ad (i.e., a default for all assets) (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_linkasset
type LinkAsset struct {
	// Landing URL of the clickable link.
	URL string `json:"url,omitempty"`
	// Fallback URL for deep-link to be used if the URL specified in url is not supported by the device.
	URLFB string `json:"urlfb,omitempty"`
	// List of third-party tracker URLs to be fired on click.
	Trkr []string `json:"trkr,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
