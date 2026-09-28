package adcom1

// DOOHMultiplierMeasurementSourceType identifies the types of entities that provide quantity measurement for impression multipliers, which are common in DOOH (Digital Out of Home) advertising (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_doohmultipliermeasurementmourcetypes
type DOOHMultiplierMeasurementSourceType int8

// MultiplierMeasurementSourceType options.
const (
	MultiplierUnknown                   DOOHMultiplierMeasurementSourceType = 0
	MultiplierMeasurementVendorProvided DOOHMultiplierMeasurementSourceType = 1
	MultiplierPublisherProvided         DOOHMultiplierMeasurementSourceType = 2
	MultiplierExchangeProvided          DOOHMultiplierMeasurementSourceType = 3
)
