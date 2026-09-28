package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// VideoPlacement object signals that the placement may be a video placement and provides additional detail about permitted video ads (e.g., VAST) (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_videoplacement
type VideoPlacement struct {
	// Placement subtype.
	PType VideoPlacementSubtype `json:"ptype,omitempty"`
	// Placement position on screen.
	Pos PlacementPosition `json:"pos,omitempty"`
	// Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll placements.
	Delay StartDelay `json:"delay,omitempty"`
	// Indicates if the placement imposes ad skippability, where 0 = no, 1 = yes.
	Skip int8 `json:"skip,omitempty"`
	// The placement allows creatives of total duration greater than this number of seconds to be skipped; only applicable if the ad is skippable.
	SkipMin int64 `json:"skipmin,omitempty"`
	// Number of seconds a creative must play before the placement enables skipping; only applicable if the ad is skippable.
	SkipAfter int64 `json:"skipafter,omitempty"`
	// Playback method in use for this placement.
	PlayMethod PlaybackMethod `json:"playmethod,omitempty"`
	// The event that causes playback to end for this placement.
	PlayEnd PlaybackCessationMode `json:"playend,omitempty"`
	// Indicates the click type of the placement.
	ClkType ClickType `json:"clktype,omitempty"`
	// Array of supported mime types (e.g., “video/mp4”).
	MIME []string `json:"mime,omitempty"`
	// List of supported APIs for this placement.
	API []APIFramework `json:"api,omitempty"`
	// Creative subtypes permitted for this placement.
	CType []MediaCreativeSubtype `json:"ctype,omitempty"`
	// Width of the placement in units specified by unit.
	W int64 `json:"w,omitempty"`
	// Height of the placement in units specified by unit.
	H int64 `json:"h,omitempty"`
	// Units of size used for w and h attributes.
	Unit SizeUnit `json:"unit,omitempty"`
	// Minimum creative duration in seconds.
	MinDur int64 `json:"mindur,omitempty"`
	// Maximum creative duration in seconds.
	MaxDur int64 `json:"maxdur,omitempty"`
	// Precise acceptable durations for video creatives in seconds.
	RqdDurs []int64 `json:"rqddurs,omitempty"`
	// Maximum extended creative duration if extension is allowed.
	MaxExt int64 `json:"maxext,omitempty"`
	// Minimum bit rate of the creative in Kbps.
	MinBitR int64 `json:"minbitr,omitempty"`
	// Maximum bit rate of the creative in Kbps.
	MaxBitR int64 `json:"maxbitr,omitempty"`
	// Array of supported creative delivery methods.
	Delivery []DeliveryMethod `json:"delivery,omitempty"`
	// The maximum number of ads that can be played in an ad pod.
	MaxSeq int64 `json:"maxseq,omitempty"`
	// Indicates the total amount of time in seconds that advertisers may fill for a “dynamic” video ad pod, or the dynamic portion of a “hybrid” ad pod.
	PodDur int64 `json:"poddur,omitempty"`
	// Unique identifier indicating that an impression opportunity belongs to a video ad pod.
	PodID int64 `json:"podid,omitempty"`
	// The sequence (position) of the video ad pod within a content stream.
	PodSeq PodSequence `json:"podseq,omitempty"`
	// For video ad pods, this value indicates that the seller can guarantee delivery against the indicated slot position in the pod.
	SlotInPod SlotPositionInPod `json:"slotinpod,omitempty"`
	// Minimum CPM per second.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`
	// Indicates if the creative must be linear, nonlinear, etc. If none specified, no restrictions are assumed.
	Linear LinearityMode `json:"linear,omitempty"`
	// Indicates if letterboxing of 4:3 creatives into a 16:9 window is allowed, where 0 = no, 1 = yes.
	Boxing int8 `json:"boxing,omitempty"`
	// Array of objects indicating that companion ads are available and providing the specifications thereof.
	Comp []Companion `json:"comp,omitempty"`
	// Supported companion ad types; recommended if companion ads are specified in comp.
	CompType []CompanionType `json:"comptype,omitempty"`
	// Directions in which the creative (video placement) is permitted to expand.
	ExpDir []ExpandableDirection `json:"expdir,omitempty"`
	// Directions in which the creative (video overlay) is permitted to expand.
	OverlayExpDir []ExpandableDirection `json:"overlayexpdir,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
