package openrtb2

// SSAI indicates if server-side ad insertion (e.g., stitching an ad into an audio or video stream) is in use and the impact of this on asset and tracker retrieval (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectimp
type AdInsertion int8

const (
	AdInsertUnknown                 AdInsertion = 0 // status unknown
	AdInsertClient                  AdInsertion = 1 // all client-side (i.e., not server-side)
	AdInsertServerStitchClientTrack AdInsertion = 2 // assets stitched server-side but tracking pixels fired client-side
	AdInsertServer                  AdInsertion = 3 // all server-side
)
