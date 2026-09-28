package adcom1

// ClickType represents types of creative activation (i.e., click) behavior types (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_clicktypes
type ClickType int16

// Types of creative activation (i.e., click) behavior types.
const (
	ClickNonClickable ClickType = 0 // Non-Clickable
	ClickUnknown      ClickType = 1 // Clickable - Details Unknown
	ClickEmbedded     ClickType = 2 // Clickable - Embedded Browser/Webview
	ClickNative       ClickType = 3 // Clickable - Native Browser
)
