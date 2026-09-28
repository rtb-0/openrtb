package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// DisplayPlacement object signals that the placement may be a display placement (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_displayplacement
type DisplayPlacement struct {
	// Placement position on screen.
	Pos PlacementPosition `json:"pos,omitempty"`
	// Indicates if this is an interstitial placement, where 0 = no, 1 = yes.
	Instl int8 `json:"instl,omitempty"`
	// Indicates if the placement will be loaded into an iframe or not, where 0 = unfriendly iframe or unknown, 1 = top frame, friendly iframe, or SafeFrame.
	TopFrame int8 `json:"topframe,omitempty"`
	// Array of iframe busters supported by this placement.
	IfrBust []string `json:"ifrbust,omitempty"`
	// Indicates the click type of this placement.
	ClkType ClickType `json:"clktype,omitempty"`
	// AMPHTML rendering treatment for AMP ads in this placement, where 1 = early loading, 2 = standard loading.
	AMPRen int8 `json:"ampren,omitempty"`
	// The display placement type.
	PType DisplayPlacementType `json:"ptype,omitempty"`
	// The context of the placement.
	Context DisplayContextType `json:"context,omitempty"`
	// Array of supported mime types (e.g., “image/jpeg”, “image/gif”).
	MIME []string `json:"mime,omitempty"`
	// List of supported APIs.
	API []APIFramework `json:"api,omitempty"`
	// Creative subtypes permitted.
	CType []DisplayCreativeSubtype `json:"ctype,omitempty"`
	// Width of the placement in units specified by unit.
	W int64 `json:"w,omitempty"`
	// Width of the placement in units specified by unit.
	H int64 `json:"h,omitempty"`
	// Unit of size used for placement size (i.e., w and h attributes).
	Unit SizeUnit `json:"unit,omitempty"`
	// Indicator of whether or not the placement supports a buyer-specific privacy notice URL, where 0 = no, 1 = yes.
	Priv int8 `json:"priv,omitempty"`
	// Array of objects that govern the attributes (e.g., sizes) of a banner display placement.
	DisplayFmt []DisplayFormat `json:"displayfmt,omitempty"`
	// This object specified the required and permitted assets and attributes of a native display placement.
	NativeFmt *NativeFormat `json:"nativefmt,omitempty"`
	// Array of supported ad tracking events.
	Event []EventSpec `json:"event,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
