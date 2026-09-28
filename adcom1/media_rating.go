package adcom1

// MediaRating represents media ratings used in describing content based on the TAG Inventory Quality Guidelines (IQG) v2.1 categorization (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_mediaratings
type MediaRating int8

// Media ratings used in describing content based on the TAG Inventory Quality Guidelines (IQG) v2.1 categorization.
const (
	MediaRatingAll    MediaRating = 1 // All Audiences
	MediaRatingOver12 MediaRating = 2 // Everyone Over Age 12
	MediaRatingMature MediaRating = 3 // Mature Audiences
)
