package adcom1

// PodSequence identifies the pod sequence field, for use in audio and video content streams with one or more ad pods (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_podsequence
type PodSequence int8

// PodSequence options.
const (
	PodSeqLast  PodSequence = -1 // Last pod in the content stream.
	PodSeqAny   PodSequence = 0  // Any pod in the content stream.
	PodSeqFirst PodSequence = 1  // First pod in the content stream.
)
