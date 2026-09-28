package adcom1

// LinearityMode represents options for media linearity, typically for video (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_linearitymodes
type LinearityMode int8

// Options for media linearity, typically for video.
const (
	LinearityLinear    LinearityMode = 1 // Linear (i.e., In-Stream such as Pre-Roll, Mid-Roll, Post-Roll)
	LinearityNonLinear LinearityMode = 2 // Non-Linear (i.e., Overlay)
)
