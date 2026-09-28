package adcom1

// SlotPositionInPod identifies the slot position in pod field, for use in audio and video ad pods (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_slotpositioninpod
type SlotPositionInPod int8

// SlotPositionInPod options.
const (
	SlotPosLast        SlotPositionInPod = -1 // Last ad in the pod.
	SlotPosAny         SlotPositionInPod = 0  // Any ad in the pod.
	SlotPosFirst       SlotPositionInPod = 1  // First ad in the pod.
	SlotPosFirstOrLast SlotPositionInPod = 2  // First or Last ad in the pod.
)
