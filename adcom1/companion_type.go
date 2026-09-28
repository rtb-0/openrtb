package adcom1

// CompanionType represents options to indicate markup types allowed for companion ads that apply to video and audio ads (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_companiontypes
type CompanionType int8

// options to indicate markup types allowed for companion ads that apply to video and audio ads.
const (
	CompanionStatic CompanionType = 1 // Static Resource
	CompanionHTML   CompanionType = 2 // HTML Resource
	CompanionIFrame CompanionType = 3 // iframe Resource
)
