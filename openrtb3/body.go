package openrtb3

// Body is a top-level wrapper for OpenRTB requests (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_openrtb
type Body[A any] struct {
	OpenRTB OpenRTB[A] `json:"openrtb"`
}
