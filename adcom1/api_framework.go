package adcom1

// APIFramework represents API frameworks either supported by a placement or required by an ad (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_apiframeworks
type APIFramework int16

// API frameworks either supported by a placement or required by an ad.
const (
	APIVPAID10 APIFramework = 1 // VPAID 1.0
	APIVPAID20 APIFramework = 2 // VPAID 2.0
	APIMRAID10 APIFramework = 3 // MRAID 1.0
	APIORMMA   APIFramework = 4 // ORMMA
	APIMRAID20 APIFramework = 5 // MRAID 2.0
	APIMRAID30 APIFramework = 6 // MRAID 3.0
	APIOMID10  APIFramework = 7 // OMID 1.0
	APISIMID10 APIFramework = 8 // SIMID 1.0
	APISIMID11 APIFramework = 9 // SIMID 1.1
)
