package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object represents the most general type of impression (OpenRTB 2.6 §3.2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbanner
type Banner struct {
	// Array of format objects (Section 3.2.10) representing the banner sizes permitted.
	Format []Format `json:"format,omitempty"`
	// Exact width in device independent pixels (DIPS); recommended if no format objects are specified.
	W *int64 `json:"w,omitempty"`
	// Exact height in device independent pixels (DIPS); recommended if no format objects are specified.
	H *int64 `json:"h,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 banners still decode.
	WMax int64 `json:"wmax,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 banners still decode.
	HMax int64 `json:"hmax,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 banners still decode.
	WMin int64 `json:"wmin,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 banners still decode.
	HMin int64 `json:"hmin,omitempty"`
	// Blocked banner ad types.
	BType []BannerAdType `json:"btype,omitempty"`
	// Blocked creative attributes.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`
	// Ad position on screen.
	Pos *adcom1.PlacementPosition `json:"pos,omitempty"`
	// Content MIME types supported.
	MIMEs []string `json:"mimes,omitempty"`
	// Indicates if the banner is in the top frame as opposed to an iframe, where 0 = no, 1 = yes.
	TopFrame int8 `json:"topframe,omitempty"`
	// Directions in which the banner may expand.
	ExpDir []adcom1.ExpandableDirection `json:"expdir,omitempty"`
	// List of supported API frameworks for this impression.
	API []adcom1.APIFramework `json:"api,omitempty"`
	// Unique identifier for this banner object.
	ID string `json:"id,omitempty"`
	// Relevant only for Banner objects used with a Video object (Section 3.2.7) in an array of companion ads.
	Vcm *int8 `json:"vcm,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
