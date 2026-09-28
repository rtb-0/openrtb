package native1

// Below is a list of common image asset element types of native advertising at the time of writing this spec (OpenRTB Native 1.2 §7.5).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#7-5
// Values of 500 and greater are exchange-specific.
type ImageAssetType int16

const (
	ImageAssetTypeIcon ImageAssetType = 1 // Icon; Icon image; Optional. Max height: at least 50; aspect ratio: 1:1
	ImageAssetTypeLogo ImageAssetType = 2 // Logo; Logo image for the brand/app. Deprecated in version 1.2 - use type 1 Icon.
	// Main; Large image preview for the ad. At least one of 2 size variants required:
	//   Small Variant:
	//     max height: at least 200
	//     max width: at least 200, 267, or 382
	//     aspect ratio: 1:1, 4:3, or 1.91:1
	//   Large Variant:
	//     max height: at least 627
	//     max width: at least 627, 836, or 1198
	//     aspect ratio: 1:1, 4:3, or 1.91:1
	ImageAssetTypeMain ImageAssetType = 3
)
