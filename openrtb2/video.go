package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object represents a video impression (OpenRTB 2.6 §3.2.7).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectvideo
type Video struct {
	// Content MIME types supported (e.g., “video/x-ms-wmv”, “video/mp4”).
	MIMEs []string `json:"mimes"`
	// Minimum video ad duration in seconds.
	MinDuration int64 `json:"minduration,omitempty"`
	// Maximum video ad duration in seconds.
	MaxDuration int64 `json:"maxduration,omitempty"`
	// Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll ad placements.
	StartDelay *adcom1.StartDelay `json:"startdelay,omitempty"`
	// Indicates the maximum number of ads that may be served into a “dynamic” video ad pod (where the precise number of ads is not predetermined by the seller).
	MaxSeq int64 `json:"maxseq,omitempty"`
	// Indicates the total amount of time in seconds that advertisers may fill for a “dynamic” video ad pod (See Section 7.6 for more details), or the dynamic portion of a “hybrid” ad pod.
	PodDur int64 `json:"poddur,omitempty"`
	// Array of supported video protocols.
	Protocols []adcom1.MediaCreativeSubtype `json:"protocols,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 video still decodes.
	Protocol adcom1.MediaCreativeSubtype `json:"protocol,omitempty"`
	// Width of the video player in device independent pixels (DIPS).
	W *int64 `json:"w,omitempty"`
	// Height of the video player in device independent pixels (DIPS).
	H *int64 `json:"h,omitempty"`
	// Unique identifier indicating that an impression opportunity belongs to a video ad pod.
	PodID string `json:"podid,omitempty"`
	// The sequence (position) of the video ad pod within a content stream.
	PodSeq adcom1.PodSequence `json:"podseq,omitempty"`
	// Precise acceptable durations for video creatives in seconds.
	RqdDurs []int64 `json:"rqddurs,omitempty"`
	// Video placement type for the impression.
	Placement adcom1.VideoPlacementSubtype `json:"placement,omitempty"`
	// Video placement type for the impression.
	Plcmt adcom1.VideoPlcmtSubtype `json:"plcmt,omitempty"`
	// Indicates if the impression must be linear, nonlinear, etc. If none specified, assume all are allowed.
	Linearity adcom1.LinearityMode `json:"linearity,omitempty"`
	// Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	Skip *int8 `json:"skip,omitempty"`
	// Videos of total duration greater than this number of seconds can be skippable; only applicable if the ad is skippable.
	SkipMin int64 `json:"skipmin,omitempty"`
	// Number of seconds a video must play before skipping is enabled; only applicable if the ad is skippable
	SkipAfter int64 `json:"skipafter,omitempty"`
	// If multiple ad impressions are offered in the same bid request, the sequence number will allow for the coordinated delivery of multiple creatives.
	Sequence int8 `json:"sequence,omitempty"`
	// For video ad pods, this value indicates that the seller can guarantee delivery against the indicated slot position in the pod.
	SlotInPod adcom1.SlotPositionInPod `json:"slotinpod,omitempty"`
	// Minimum CPM per second.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`
	// Blocked creative attributes.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`
	// Maximum extended ad duration if extension is allowed.
	MaxExtended int64 `json:"maxextended,omitempty"`
	// Minimum bit rate in Kbps (kilobits per second).
	MinBitRate int64 `json:"minbitrate,omitempty"`
	// Maximum bit rate in Kbps (kilobits per second).
	MaxBitRate int64 `json:"maxbitrate,omitempty"`
	// Indicates if letter-boxing of 4:3 content into a 16:9 window is allowed, where 0 = no, 1 = yes.
	BoxingAllowed *int8 `json:"boxingallowed,omitempty"`
	// Playback methods that may be in use.
	PlaybackMethod []adcom1.PlaybackMethod `json:"playbackmethod,omitempty"`
	// The event that causes playback to end.
	PlaybackEnd adcom1.PlaybackCessationMode `json:"playbackend,omitempty"`
	// Supported delivery methods (e.g., streaming, progressive).
	Delivery []adcom1.DeliveryMethod `json:"delivery,omitempty"`
	// Ad position on screen.
	Pos *adcom1.PlacementPosition `json:"pos,omitempty"`
	// Array of Banner objects (Section 3.2.6) if companion ads are available.
	CompanionAd []Banner `json:"companionad,omitempty"`
	// List of supported API frameworks for this impression.
	API []adcom1.APIFramework `json:"api,omitempty"`
	// Supported VAST companion ad types.
	CompanionType []adcom1.CompanionType `json:"companiontype,omitempty"`
	// Indicates pod deduplication settings that will be applied to bid responses.
	PodDedupe []adcom1.PodDedupe `json:"poddedupe,omitempty"`
	// An array of DurFloors objects (Section 3.2.35) indicating the floor prices for video creatives of various durations that the buyer may bid with.
	DurFloors []DurFloors `json:"durfloors,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
