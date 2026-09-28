package adcom1

// MatchMethod represents various ways an ID could be matched to an ad request, and if they pertain to a single property or app (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_idmatchmethod
type MatchMethod int8

// Various ways an ID could be matched to an ad request
const (
	MatchMethodUnknown           MatchMethod = 0
	MatchMethodNoMatch           MatchMethod = 1 // No matching has occurred. The associated ID came directly from a 3rd-party cookie or OS-provided resettable device ID for advertising (IFA).
	MatchMethodBrowserCookieSync MatchMethod = 2 // Real time cookie sync as described in Appendix C (Cookie Based ID Syncing) of OpenRTB 2.x.
	MatchMethodAuthenticated     MatchMethod = 3 // ID match was based on user authentication such as an email login or hashed PII.
	MatchMethodObserved          MatchMethod = 4 // ID match was based on a 1st party observation, but without user authentication (e.g. GUID, SharedID, Session IDs, CHIPS or other 1st party cookies contained in localStorage).
	MatchMethodInference         MatchMethod = 5 // ID match was inferred from linkage based on non-authenticated features across multiple browsers or devices (e.g. IP address and/or UserAgent).
)
