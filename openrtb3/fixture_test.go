package openrtb3_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/rtb-0/openrtb/adm/popunder"
	"github.com/rtb-0/openrtb/adm/vast"
	"github.com/rtb-0/openrtb/openrtb3"
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
	var body openrtb3.Body[string]
	if err := json.Unmarshal(readFixture(t, "testdata/banner/request.json"), &body); err != nil {
		t.Fatal(err)
	}
	req := body.OpenRTB.Request
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}
	display := req.Item[0].Spec.Placement.Display
	if display == nil || display.W != 728 || display.H != 90 {
		t.Fatalf("display %+v", display)
	}
	if body.OpenRTB.DomainVer != "1.0" || req.Context.Site.Domain != "example.com" {
		t.Fatalf("request %+v", req)
	}

	var respBody openrtb3.Body[string]
	if err := json.Unmarshal(readFixture(t, "testdata/banner/response.json"), &respBody); err != nil {
		t.Fatal(err)
	}
	resp := respBody.OpenRTB.Response
	if err := resp.Validate(); err != nil {
		t.Fatal(err)
	}
	adm := resp.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value
	if adm != "<div>banner</div>" {
		t.Fatalf("adm %q", adm)
	}
	out, err := json.Marshal(respBody)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb3.Body[string]
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value != adm {
		t.Fatal("adm round trip")
	}
}

func TestNativeFixture(t *testing.T) {
	var body openrtb3.Body[string]
	if err := json.Unmarshal(readFixture(t, "testdata/native/request.json"), &body); err != nil {
		t.Fatal(err)
	}
	if err := body.OpenRTB.Request.Validate(); err != nil {
		t.Fatal(err)
	}
	assets := body.OpenRTB.Request.Item[0].Spec.Placement.Display.NativeFmt.Asset
	if len(assets) != 2 || assets[0].Title == nil || assets[0].Title.Len != 50 {
		t.Fatalf("nativefmt %+v", assets)
	}

	var respBody openrtb3.Body[string]
	if err := json.Unmarshal(readFixture(t, "testdata/native/response.json"), &respBody); err != nil {
		t.Fatal(err)
	}
	if err := respBody.OpenRTB.Response.Validate(); err != nil {
		t.Fatal(err)
	}
	native := respBody.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Display.Native
	if native == nil || native.Ver != "1.2" || native.Link.URL != "https://example.com/landing" || len(native.Assets) != 1 {
		t.Fatalf("native %+v", native)
	}
}

func TestPopunderFixture(t *testing.T) {
	var body openrtb3.Body[popunder.Ad]
	if err := json.Unmarshal(readFixture(t, "testdata/popunder/request.json"), &body); err != nil {
		t.Fatal(err)
	}
	if err := body.OpenRTB.Request.Validate(); err == nil {
		t.Fatal("expected missing spec")
	}
	if err := body.OpenRTB.Request.Validate(openrtb3.AllowMissingFormat()); err != nil {
		t.Fatal(err)
	}

	var respBody openrtb3.Body[popunder.Ad]
	if err := json.Unmarshal(readFixture(t, "testdata/popunder/response.json"), &respBody); err != nil {
		t.Fatal(err)
	}
	ad := respBody.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value
	if ad.URL != "http://google.com" {
		t.Fatalf("url %q", ad.URL)
	}
	out, err := json.Marshal(respBody)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb3.Body[popunder.Ad]
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value.Raw != ad.Raw {
		t.Fatalf("raw %q", again.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value.Raw)
	}

	var urlBody openrtb3.Body[popunder.Ad]
	if err := json.Unmarshal(readFixture(t, "testdata/popunder/response-url.json"), &urlBody); err != nil {
		t.Fatal(err)
	}
	if urlBody.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value.URL != "https://example.com/pop" {
		t.Fatal("url adm")
	}
}

func TestVideoAndAudioFixtures(t *testing.T) {
	var video openrtb3.Body[vast.Document]
	if err := json.Unmarshal(readFixture(t, "testdata/video/request.json"), &video); err != nil {
		t.Fatal(err)
	}
	if err := video.OpenRTB.Request.Validate(); err != nil {
		t.Fatal(err)
	}
	mimes := video.OpenRTB.Request.Item[0].Spec.Placement.Video.MIME
	if len(mimes) != 1 || mimes[0] != "video/mp4" {
		t.Fatalf("mimes %v", mimes)
	}

	var audio openrtb3.Body[string]
	if err := json.Unmarshal(readFixture(t, "testdata/audio/request.json"), &audio); err != nil {
		t.Fatal(err)
	}
	if err := audio.OpenRTB.Request.Validate(); err != nil {
		t.Fatal(err)
	}
	if audio.OpenRTB.Request.Item[0].Spec.Placement.Audio.MIME[0] != "audio/mpeg" {
		t.Fatal("audio mime")
	}

	var resp openrtb3.Body[vast.Document]
	if err := json.Unmarshal(readFixture(t, "testdata/video/response.json"), &resp); err != nil {
		t.Fatal(err)
	}
	doc := resp.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Video.Adm.Value
	if doc.Version != "2.0" || len(doc.Ads) != 1 || !doc.Ads[0].Wrapper || doc.Ads[0].AdTagURI == "" {
		t.Fatalf("vast %+v", doc)
	}
	out, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var again openrtb3.Body[vast.Document]
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if again.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Video.Adm.Value.Raw != doc.Raw {
		t.Fatal("vast raw round trip")
	}

	var audioResp openrtb3.Body[string]
	if err := json.Unmarshal(readFixture(t, "testdata/audio/response.json"), &audioResp); err != nil {
		t.Fatal(err)
	}
	if audioResp.OpenRTB.Response.SeatBid[0].Bid[0].Media.Ad.Audio.Adm.Value != `<DAAST version="1.0"></DAAST>` {
		t.Fatal("audio adm")
	}
}
