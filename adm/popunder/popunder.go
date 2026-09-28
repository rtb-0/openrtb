// Package popunder parses a popunder adm value.
// The value is either an absolute URL or an XML document:
//
//	<?xml version="1.0" encoding="ISO-8859-1"?>
//	<ad><popunderAd><url><![CDATA[http://example.com]]></url></popunderAd></ad>
package popunder

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"strings"

	"github.com/rtb-0/openrtb/common"
)

// Ad is a popunder click target.
// Raw is the adm text that produced URL, and marshal writes Raw back unchanged.
type Ad struct {
	URL string
	Raw string
}

type xmlAd struct {
	XMLName xml.Name `xml:"ad"`
	Pop     struct {
		URL string `xml:"url"`
	} `xml:"popunderAd"`
}

// UnmarshalJSON reads a JSON string that holds a URL or the popunder XML document.
func (a *Ad) UnmarshalJSON(data []byte) error {
	raw, err := common.UnwrapJSONString(data)
	if err != nil {
		return err
	}
	text := string(bytes.TrimSpace(raw))
	a.Raw = text
	if strings.HasPrefix(text, "http://") || strings.HasPrefix(text, "https://") {
		a.URL = text
		return nil
	}
	dec := xml.NewDecoder(strings.NewReader(text))
	dec.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return input, nil
	}
	var doc xmlAd
	if err := dec.Decode(&doc); err != nil {
		return err
	}
	a.URL = strings.TrimSpace(doc.Pop.URL)
	return nil
}

// MarshalJSON writes the original adm text as a JSON string.
func (a Ad) MarshalJSON() ([]byte, error) {
	if a.Raw != "" {
		return common.MarshalString(a.Raw)
	}
	return common.MarshalString(a.URL)
}

// URLAd builds a popunder from an absolute URL.
func URLAd(url string) Ad {
	return Ad{URL: url, Raw: url}
}

// Compile-time check that Ad can sit in a JSON string field.
var _ json.Unmarshaler = (*Ad)(nil)
