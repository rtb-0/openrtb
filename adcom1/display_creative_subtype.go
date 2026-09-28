package adcom1

// DisplayCreativeSubtype represents subtypes of display ad creatives (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_creativesubtypesdisplay
type DisplayCreativeSubtype int8

// Subtypes of display ad creatives.
const (
	CreativeHTML   DisplayCreativeSubtype = 1 // HTML
	CreativeAMP    DisplayCreativeSubtype = 2 // AMPHTML
	CreativeImage  DisplayCreativeSubtype = 3 // Structured Image Object
	CreativeNative DisplayCreativeSubtype = 4 // Structured Native Object
)
