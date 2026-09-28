package adcom1

// ProductionQuality represents content quality (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_productionqualities
type ProductionQuality int8

// Options for content quality.
const (
	ProductionUnknown      ProductionQuality = 0 // Unknown
	ProductionProfessional ProductionQuality = 1 // Professionally Produced
	ProductionProsumer     ProductionQuality = 2 // Prosumer
	ProductionUser         ProductionQuality = 3 // User Generated (UGC)
)
