package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// This object describes an ad placement or impression being auctioned (OpenRTB 2.6 §3.2.4).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectimp
type Imp struct {
	// A unique identifier for this impression within the context of the bid request (typically, starts with 1 and increments).
	ID string `json:"id"`
	// An array of Metric object (Section 3.2.5).
	Metric []Metric `json:"metric,omitempty"`
	// A Banner object (Section 3.2.6); required if this impression is offered as a banner ad opportunity.
	Banner *Banner `json:"banner,omitempty"`
	// A Video object (Section 3.2.7); required if this impression is offered as a video ad opportunity.
	Video *Video `json:"video,omitempty"`
	// An Audio object (Section 3.2.8); required if this impression is offered as an audio ad opportunity.
	Audio *Audio `json:"audio,omitempty"`
	// A Native object (Section 3.2.9); required if this impression is offered as a native ad opportunity.
	Native *Native `json:"native,omitempty"`
	// A Pmp object (Section 3.2.11) containing any private marketplace deals in effect for this impression.
	PMP *PMP `json:"pmp,omitempty"`
	// Name of ad mediation partner, SDK technology, or player responsible for rendering ad (typically video or mobile).
	DisplayManager string `json:"displaymanager,omitempty"`
	// Version of ad mediation partner, SDK technology, or player responsible for rendering ad (typically video or mobile).
	DisplayManagerVer string `json:"displaymanagerver,omitempty"`
	// 1 = the ad is interstitial or full screen, 0 = not interstitial.
	Instl int8 `json:"instl,omitempty"`
	// Identifier for specific ad placement or ad tag that was used to initiate the auction.
	TagID string `json:"tagid,omitempty"`
	// Minimum bid for this impression expressed in CPM.
	BidFloor float64 `json:"bidfloor,omitempty"`
	// Currency specified using ISO-4217 alpha codes.
	BidFloorCur string `json:"bidfloorcur,omitempty"`
	// Indicates the type of browser opened upon clicking the creative in an app, where 0 = embedded, 1 = native.
	ClickBrowser *int8 `json:"clickbrowser,omitempty"`
	// Flag to indicate if the impression requires secure HTTPS URL creative assets and markup, where 0 = non-secure, 1 = secure.
	Secure *int8 `json:"secure,omitempty"`
	// Array of exchange-specific names of supported iframe busters.
	IframeBuster []string `json:"iframebuster,omitempty"`
	// Indicates whether the user receives a reward for viewing the ad, where 0 = no, 1 = yes.
	Rwdd int8 `json:"rwdd,omitempty"`
	// Indicates if server-side ad insertion (e.g., stitching an ad into an audio or video stream) is in use and the impact of this on asset and tracker retrieval, where 0 = status unknown, 1 = all client-side (i.e., not server-side), 2 = assets stitched server-side but tracking pixels fired client-side, 3 = all server-side.
	SSAI AdInsertion `json:"ssai,omitempty"`
	// Advisory as to the number of seconds that may elapse between the auction and the actual impression.
	Exp int64 `json:"exp,omitempty"`
	// A means of passing a multiplier in the bid request, representing the total quantity of impressions for adverts that display to more than one person.
	Qty *Qty `json:"qty,omitempty"`
	// Timestamp when the item is estimated to be fulfilled (e.g. when a DOOH impression will be displayed) in Unix format (i.e., milliseconds since the epoch).
	DT float64 `json:"dt,omitempty"`
	// Details about ad slots being refreshed automatically.
	Refresh *Refresh `json:"refresh,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
	// Popunder marks an impression that intentionally has no banner, video, audio, or native object.
	Popunder bool `json:"-"`
}
