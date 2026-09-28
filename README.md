# openrtb

Go types for OpenRTB 2.6, OpenRTB 3.0, Native 1.2, and AdCOM 1.0. Module `github.com/rtb-0/openrtb`, standard library only.

`bid.adm` is a JSON string. `BidResponse[A]` parses that string into `A`.

```go
var native openrtb2.BidResponse[nresp.Response]
var pop openrtb2.BidResponse[popunder.Ad]
var video openrtb2.BidResponse[vast.Document]
var any openrtb2.BidResponse[wrapper.Markup]
var html openrtb2.BidResponse[string]

json.Unmarshal(body, &native)
native.SeatBid[0].Bid[0].Adm.Value.Link.URL
```

`native1` and `adcom1.Ad` accept either a JSON object or a JSON string containing that object. Popunder, VAST, and the wrapper marshal back to the original adm text. `adm_native` stays an object.

| Package        | Contents                                                                            |
| -------------- | ----------------------------------------------------------------------------------- |
| `openrtb2`     | OpenRTB 2.6, including 2.5 fields. `Protocol` links to both specs.                  |
| `openrtb3`     | OpenRTB 3.0. Markup is `media.ad.display.adm`.                                      |
| `native1`      | Native 1.2 enums. Requests are `native1/request`, responses are `native1/response`. |
| `adcom1`       | AdCOM 1.0 objects and lists.                                                        |
| `common`       | `StringEncoded[A]`, `Ext`, protocol metadata.                                       |
| `errs`         | Reusable errors: `Err`, `ErrVar`, `Error`.                                          |
| `adm/popunder` | Popunder URL or `<popunderAd>` XML.                                                 |
| `adm/vast`     | VAST version, inline or wrapper, tag URI, impressions, media files.                 |
| `adm/wrapper`  | Detects native, popunder, VAST, AdCOM, or leaves the text raw.                      |

Specs:

- [OpenRTB 2.6](https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md)
- [OpenRTB 2.5](https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.5.md)
- [OpenRTB 3.0](https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md)
- [Native 1.2](https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md)
- [AdCOM 1.0](https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md)

```go
req := openrtb2.NewBidRequest("id").
    Auction(openrtb2.SecondPrice).
    SetSite(openrtb2.Site{Domain: "example.com"}).
    AddImp(openrtb2.BannerImp("1", 728, 90).Pos(1).SetSecure()).
    AddImp(openrtb2.NativeImp("2", nativeReq).SetBidFloor(0.0001, "USD")).
    AddImp(openrtb2.VideoImp("3", openrtb2.Video{MIMEs: []string{"video/mp4"}})).
    AddImp(openrtb2.PopunderImp("4").SetBidFloor(0.0005, "USD"))

if err := req.Validate(); err != nil {
    // popunder traffic parsed from the wire has no format object;
    // call Validate(openrtb2.AllowMissingFormat()) or build it with PopunderImp
}
```

`ext` is `json.RawMessage`. Enumerations are `int8`, or `int16` where the list allows exchange codes of 500 and greater.
