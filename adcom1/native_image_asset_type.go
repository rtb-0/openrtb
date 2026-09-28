package adcom1

// NativeImageAssetType represents image asset types (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_nativeimageassettypes
type NativeImageAssetType int16

// Common image asset types.
const (
	// Icon: Icon image. Maximum height at least 50 device independent pixels (DIPS); aspect ratio 1:1.
	ImageAssetIcon NativeImageAssetType = 1
	// Main: Large image preview for the ad. At least one of 2 size variants required:
	//
	// Small: Maximum height at least 627 DIPS; maximum width at least 627, 836, or 1198 DIPS (i.e., aspect ratios of 1:1, 4:3, or 1.91:1, respectively).
	//
	// Large: Maximum height at least 200 DIPS; maximum width at least 200, 267, or 382 DIPS (i.e., aspect ratios of 1:1, 4:3, or 1.91:1, respectively).
	ImageAssetMain NativeImageAssetType = 3
)
