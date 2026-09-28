package adcom1

// StartDelay represents video or audio start delay (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_startdelaymodes
type StartDelay int64

// Options for the video or audio start delay.
const (
	StartPreRoll  StartDelay = 0  // Pre-Roll
	StartMidRoll  StartDelay = -1 // Generic Mid-Roll
	StartPostRoll StartDelay = -2 // Generic Post-Roll
)

// Ptr returns pointer to own value.
func (d StartDelay) Ptr() *StartDelay {
	return &d
}

// Val safely dereferences pointer, returning default value (StartDelayPreRoll) for nil.
func (d *StartDelay) Val() StartDelay {
	if d == nil {
		return StartPreRoll
	}
	return *d
}
