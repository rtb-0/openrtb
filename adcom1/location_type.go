package adcom1

// LocationType represents options to indicate how the geographic information was determined (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_locationtypes
type LocationType int8

// Options to indicate how the geographic information was determined.
const (
	LocationGPS          LocationType = 1 // GPS/Location Services
	LocationIP           LocationType = 2 // IP Address
	LocationUserProvided LocationType = 3 // User Provided (e.g., registration data)
)
