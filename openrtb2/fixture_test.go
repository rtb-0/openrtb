package openrtb2_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/rtb-0/openrtb/adm/popunder"
	"github.com/rtb-0/openrtb/adm/vast"
	"github.com/rtb-0/openrtb/adm/wrapper"
	"github.com/rtb-0/openrtb/common"
	nreq "github.com/rtb-0/openrtb/native1/request"
	nresp "github.com/rtb-0/openrtb/native1/response"
	"github.com/rtb-0/openrtb/openrtb2"
)

func readFixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestBannerFixture(t *testing.T) {
	var req openrtb2.BidRequest
	if err := json.Unmarshal(readFixture(t, "testdata/banner/request.json"), &req); err != nil {
		t.Fatal(err)
	}
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}
	if req.Imp[0].Banner == nil || req.Imp[0].Banner.W == nil || *req.Imp[0].Banner.W != 728 {
		t.Fatalf("banner width = %v", req.Imp[0].Banner)
	}

	var resp openrtb2.BidResponse[string]
	if err := json.Unmarshal(readFixture(t, "testdata/banner/response.json"), &resp); err != nil {
		t.Fatal(err)
	}
	if err := resp.Validate(); err != nil {
		t.Fatal(err)
	}
	if resp.SeatBid[0].Bid[0].Adm.Value == "" {
		t.Fatal("empty banner adm")
	}
	out, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb2.BidResponse[string]
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.SeatBid[0].Bid[0].Adm.Value != resp.SeatBid[0].Bid[0].Adm.Value {
		t.Fatalf("adm round trip %q", again.SeatBid[0].Bid[0].Adm.Value)
	}
}

func TestNativeFixture(t *testing.T) {
	var req openrtb2.BidRequest
	if err := json.Unmarshal(readFixture(t, "testdata/native/request.json"), &req); err != nil {
		t.Fatal(err)
	}
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}
	native := req.Imp[0].Native.Request.Value
	if native.Ver != "1.2" || len(native.Assets) != 3 {
		t.Fatalf("native request ver=%q assets=%d", native.Ver, len(native.Assets))
	}

	var resp openrtb2.BidResponse[nresp.Response]
	if err := json.Unmarshal(readFixture(t, "testdata/native/response.json"), &resp); err != nil {
		t.Fatal(err)
	}
	if err := resp.Validate(); err != nil {
		t.Fatal(err)
	}
	got := resp.SeatBid[0].Bid[0].Adm.Value
	if got.Link.URL != "https://example.com/landing" || len(got.Assets) != 1 {
		t.Fatalf("native adm %+v", got)
	}

	object := []byte(`{"id":"n","seatbid":[{"bid":[{"id":"b","impid":"i","price":1,"adm_native":{"ver":"1.2","link":{"url":"https://example.com/obj"},"assets":[{"id":1,"title":{"text":"T"}}]}}]}]}`)
	var obj openrtb2.BidResponse[string]
	if err := json.Unmarshal(object, &obj); err != nil {
		t.Fatal(err)
	}
	if obj.SeatBid[0].Bid[0].AdMNative == nil || obj.SeatBid[0].Bid[0].AdMNative.Link.URL != "https://example.com/obj" {
		t.Fatalf("adm_native %+v", obj.SeatBid[0].Bid[0].AdMNative)
	}
}

func TestPopunderFixture(t *testing.T) {
	var req openrtb2.BidRequest
	if err := json.Unmarshal(readFixture(t, "testdata/popunder/request.json"), &req); err != nil {
		t.Fatal(err)
	}
	if err := req.Validate(); err == nil {
		t.Fatal("expected missing format")
	}
	if err := req.Validate(openrtb2.AllowMissingFormat()); err != nil {
		t.Fatal(err)
	}

	raw := readFixture(t, "testdata/popunder/response.json")
	var resp openrtb2.BidResponse[popunder.Ad]
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatal(err)
	}
	ad := resp.SeatBid[0].Bid[0].Adm.Value
	if ad.URL != "http://google.com" {
		t.Fatalf("url %q", ad.URL)
	}
	out, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb2.BidResponse[popunder.Ad]
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.SeatBid[0].Bid[0].Adm.Value.Raw != ad.Raw {
		t.Fatalf("raw round trip %q", again.SeatBid[0].Bid[0].Adm.Value.Raw)
	}

	var urlResp openrtb2.BidResponse[popunder.Ad]
	if err := json.Unmarshal(readFixture(t, "testdata/popunder/response-url.json"), &urlResp); err != nil {
		t.Fatal(err)
	}
	if urlResp.SeatBid[0].Bid[0].Adm.Value.URL != "https://example.com/pop" {
		t.Fatalf("url adm %+v", urlResp.SeatBid[0].Bid[0].Adm.Value)
	}
}

func TestVideoAndAudioFixtures(t *testing.T) {
	var req openrtb2.BidRequest
	if err := json.Unmarshal(readFixture(t, "testdata/video/request.json"), &req); err != nil {
		t.Fatal(err)
	}
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}

	var audio openrtb2.BidRequest
	if err := json.Unmarshal(readFixture(t, "testdata/audio/request.json"), &audio); err != nil {
		t.Fatal(err)
	}
	if err := audio.Validate(); err != nil {
		t.Fatal(err)
	}

	raw := readFixture(t, "testdata/vast/response.json")
	var resp openrtb2.BidResponse[vast.Document]
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatal(err)
	}
	doc := resp.SeatBid[0].Bid[0].Adm.Value
	if doc.Version != "2.0" || len(doc.Ads) != 1 || !doc.Ads[0].Wrapper || doc.Ads[0].AdTagURI == "" {
		t.Fatalf("vast %+v", doc)
	}
	out, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb2.BidResponse[vast.Document]
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.SeatBid[0].Bid[0].Adm.Value.Raw != doc.Raw {
		t.Fatal("vast raw changed")
	}
}

func TestWrapperFixture(t *testing.T) {
	cases := []struct {
		file string
		kind wrapper.Kind
	}{
		{"testdata/banner/response.json", wrapper.KindRaw},
		{"testdata/native/response.json", wrapper.KindNative},
		{"testdata/popunder/response.json", wrapper.KindPopunder},
		{"testdata/popunder/response-url.json", wrapper.KindPopunder},
		{"testdata/vast/response.json", wrapper.KindVAST},
	}
	for _, tc := range cases {
		var resp openrtb2.BidResponse[wrapper.Markup]
		if err := json.Unmarshal(readFixture(t, tc.file), &resp); err != nil {
			t.Fatalf("%s: %v", tc.file, err)
		}
		got := resp.SeatBid[0].Bid[0].Adm.Value
		if got.Kind != tc.kind {
			t.Fatalf("%s kind %d", tc.file, got.Kind)
		}
		if got.Raw == "" {
			t.Fatalf("%s empty raw", tc.file)
		}
	}
}

func TestImpressionBuilder(t *testing.T) {
	nativeReq := nreq.Request{Ver: "1.2", Assets: []nreq.Asset{{ID: 1, Title: &nreq.Title{Len: 80}}}}
	req := openrtb2.NewBidRequest("id").
		Auction(openrtb2.SecondPrice).
		SetSite(openrtb2.Site{Domain: "example.com"}).
		AddImp(openrtb2.BannerImp("1", 728, 90).Pos(1).SetSecure()).
		AddImp(openrtb2.NativeImp("2", nativeReq).SetBidFloor(0.0001, "USD")).
		AddImp(openrtb2.VideoImp("3", openrtb2.Video{MIMEs: []string{"video/mp4"}})).
		AddImp(openrtb2.PopunderImp("4").SetBidFloor(0.0005, "USD").SetSecure())
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}
	if req.AT != openrtb2.SecondPrice || len(req.Imp) != 4 || !req.Imp[3].Popunder {
		t.Fatalf("built %+v", req)
	}
	out, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb2.BidRequest
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.Imp[1].Native.Request.Value.Ver != "1.2" {
		t.Fatalf("native ver %q", again.Imp[1].Native.Request.Value.Ver)
	}
	if err := again.Validate(openrtb2.AllowMissingFormat()); err != nil {
		t.Fatal(err)
	}
}

func TestWrapperAdCOM(t *testing.T) {
	body := []byte(`{"id":"1","seatbid":[{"bid":[{"id":"b","impid":"i","price":1,"adm":"{\"id\":\"cr\",\"display\":{\"adm\":\"<b>x</b>\",\"w\":1}}"}]}]}`)
	var resp openrtb2.BidResponse[wrapper.Markup]
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	got := resp.SeatBid[0].Bid[0].Adm.Value
	if got.Kind != wrapper.KindAdCOM || got.AdCOM == nil || got.AdCOM.Display.Adm.Value != "<b>x</b>" {
		t.Fatalf("%+v", got)
	}
}

func TestProtocolMeta(t *testing.T) {
	if openrtb2.Protocol.Version != "2.6" || len(openrtb2.Protocol.Docs) != 2 {
		t.Fatalf("%+v", openrtb2.Protocol)
	}
	_ = common.Encode("x")
}
