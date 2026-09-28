package request

import (
	"encoding/json"

	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// Video is the OpenRTB video object used by a native asset (OpenRTB Native 1.2 §4.5).
// Companion ads and duration floors stay raw so this package does not import openrtb2.
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-5
type Video struct {
	// Content MIME types supported.
	MIMEs []string `json:"mimes"`
	// Minimum video ad duration in seconds.
	MinDuration int64 `json:"minduration,omitempty"`
	// Maximum video ad duration in seconds.
	MaxDuration int64 `json:"maxduration,omitempty"`
	// Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll ad placements.
	StartDelay *adcom1.StartDelay `json:"startdelay,omitempty"`
	// Maximum number of ads in a dynamic video ad pod.
	MaxSeq int64 `json:"maxseq,omitempty"`
	// Total time in seconds that advertisers may fill for a dynamic video ad pod.
	PodDur int64 `json:"poddur,omitempty"`
	// Supported video protocols.
	Protocols []adcom1.MediaCreativeSubtype `json:"protocols,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 video still decodes.
	Protocol adcom1.MediaCreativeSubtype `json:"protocol,omitempty"`
	// Width of the video player in device independent pixels.
	W *int64 `json:"w,omitempty"`
	// Height of the video player in device independent pixels.
	H *int64 `json:"h,omitempty"`
	// Identifier of the video ad pod.
	PodID string `json:"podid,omitempty"`
	// Sequence of the video ad pod within a content stream.
	PodSeq adcom1.PodSequence `json:"podseq,omitempty"`
	// Precise acceptable durations for video creatives in seconds.
	RqdDurs []int64 `json:"rqddurs,omitempty"`
	// Video placement type.
	Placement adcom1.VideoPlacementSubtype `json:"placement,omitempty"`
	// Video placement subtype.
	Plcmt adcom1.VideoPlcmtSubtype `json:"plcmt,omitempty"`
	// Linear or non-linear requirement.
	Linearity adcom1.LinearityMode `json:"linearity,omitempty"`
	// 0 = not skippable, 1 = skippable.
	Skip *int8 `json:"skip,omitempty"`
	// Videos longer than this many seconds can be skippable.
	SkipMin int64 `json:"skipmin,omitempty"`
	// Seconds before skipping is enabled.
	SkipAfter int64 `json:"skipafter,omitempty"`
	// Sequence number when several impressions are offered together.
	Sequence int8 `json:"sequence,omitempty"`
	// Guaranteed slot position in a pod.
	SlotInPod adcom1.SlotPositionInPod `json:"slotinpod,omitempty"`
	// Minimum CPM per second.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`
	// Blocked creative attributes.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`
	// Maximum extended ad duration in seconds.
	MaxExtended int64 `json:"maxextended,omitempty"`
	// Minimum bit rate in Kbps.
	MinBitRate int64 `json:"minbitrate,omitempty"`
	// Maximum bit rate in Kbps.
	MaxBitRate int64 `json:"maxbitrate,omitempty"`
	// 0 = letter-boxing is not allowed, 1 = allowed.
	BoxingAllowed *int8 `json:"boxingallowed,omitempty"`
	// Playback methods that may be in use.
	PlaybackMethod []adcom1.PlaybackMethod `json:"playbackmethod,omitempty"`
	// Event that causes playback to end.
	PlaybackEnd adcom1.PlaybackCessationMode `json:"playbackend,omitempty"`
	// Supported delivery methods.
	Delivery []adcom1.DeliveryMethod `json:"delivery,omitempty"`
	// Ad position on screen.
	Pos *adcom1.PlacementPosition `json:"pos,omitempty"`
	// Companion ads, kept raw to avoid importing the OpenRTB banner object.
	CompanionAd []json.RawMessage `json:"companionad,omitempty"`
	// Supported API frameworks.
	API []adcom1.APIFramework `json:"api,omitempty"`
	// Supported VAST companion ad types.
	CompanionType []adcom1.CompanionType `json:"companiontype,omitempty"`
	// Pod deduplication settings.
	PodDedupe []adcom1.PodDedupe `json:"poddedupe,omitempty"`
	// Duration floors, kept raw to avoid importing OpenRTB.
	DurFloors []json.RawMessage `json:"durfloors,omitempty"`
	// Exchange-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
