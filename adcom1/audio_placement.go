package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// AudioPlacement object signals that the placement may be an audio placement and provides additional detail about permitted audio ads (e.g., DAAST) (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_audioplacement
type AudioPlacement struct {
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
	// Type of audio feed of this placement.
	Feed FeedType `json:"feed,omitempty"`
	// Volume normalization mode of this placement.
	NVol VolumeNormalizationMode `json:"nvol,omitempty"`
	// Array of supported mime types (e.g., “audio/mp4”).
	MIME []string `json:"mime,omitempty"`
	// List of supported APIs for this placement.
	API []APIFramework `json:"api,omitempty"`
	// Creative subtypes permitted for this placement.
	CType []MediaCreativeSubtype `json:"ctype,omitempty"`
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
	// Array of objects indicating that companion ads are available and providing the specifications thereof.
	Comp []Companion `json:"comp,omitempty"`
	// Supported companion ad types; recommended if companion ads are specified in comp.
	CompType []CompanionType `json:"comptype,omitempty"`
	// Directions in which the creative (video overlay) is permitted to expand.
	OverlayExpDir []ExpandableDirection `json:"overlayexpdir,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
