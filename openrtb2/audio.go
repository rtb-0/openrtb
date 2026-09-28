package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object represents an audio type impression (OpenRTB 2.6 §3.2.8).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectaudio
type Audio struct {
	// Content MIME types supported (e.g., “audio/mp4”).
	MIMEs []string `json:"mimes"`
	// Minimum audio ad duration in seconds.
	MinDuration int64 `json:"minduration,omitempty"`
	// Maximum audio ad duration in seconds.
	MaxDuration int64 `json:"maxduration,omitempty"`
	// Indicates the total amount of time that advertisers may fill for a "dynamic" audio ad pod, or the dynamic portion of a "hybrid" ad pod.
	PodDur int64 `json:"poddur,omitempty"`
	// Array of supported audio protocols.
	Protocols []adcom1.MediaCreativeSubtype `json:"protocols,omitempty"`
	// Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll ad placements.
	StartDelay *adcom1.StartDelay `json:"startdelay,omitempty"`
	// Precise acceptable durations for audio creatives in seconds.
	RqdDurs []int64 `json:"rqddurs,omitempty"`
	// Unique identifier indicating that an impression opportunity belongs to an audioad pod.
	PodID string `json:"podid,omitempty"`
	// The sequence (position) of the audio ad pod within a content stream.
	PodSeq adcom1.PodSequence `json:"podseq,omitempty"`
	// If multiple ad impressions are offered in the same bid request, the sequence number will allow for the coordinated delivery of multiple creatives.
	Sequence int64 `json:"sequence,omitempty"`
	// For audio ad pods, this value indicates that the seller can guarantee delivery against the indicated sequence.
	SlotInPod adcom1.SlotPositionInPod `json:"slotinpod,omitempty"`
	// Minimum CPM per second.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`
	// Blocked creative attributes.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`
	// Maximum extended ad duration if extension is allowed.
	MaxExtended int64 `json:"maxextended,omitempty"`
	// Minimum bit rate in Kbps (kilobits per second).
	MinBitrate int64 `json:"minbitrate,omitempty"`
	// Maximum bit rate in Kbps (kilobits per second).
	MaxBitrate int64 `json:"maxbitrate,omitempty"`
	// Supported delivery methods (e.g., streaming, progressive).
	Delivery []adcom1.DeliveryMethod `json:"delivery,omitempty"`
	// Array of Banner objects (Section 3.2.6) if companion ads are available.
	CompanionAd []Banner `json:"companionad,omitempty"`
	// List of supported API frameworks for this impression.
	API []adcom1.APIFramework `json:"api,omitempty"`
	// Supported companion ad types.
	CompanionType []adcom1.CompanionType `json:"companiontype,omitempty"`
	// The maximum number of ads that can be played in an ad pod.
	MaxSeq int64 `json:"maxseq,omitempty"`
	// Type of audio feed.
	Feed adcom1.FeedType `json:"feed,omitempty"`
	// Indicates if the ad is stitched with audio content or delivered independently, where 0 = no, 1 = yes.
	Stitched *int8 `json:"stitched,omitempty"`
	// Volume normalization mode.
	NVol *adcom1.VolumeNormalizationMode `json:"nvol,omitempty"`
	// An array of DurFloors objects (Section 3.2.35) indicating the floor prices for video creatives of various durations that the buyer may bid with.
	DurFloors []DurFloors `json:"durfloors,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
