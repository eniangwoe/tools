package utils

import "testing"
import "regexp"
import "strconv"

func TestTrimPathName(t *testing.T) {
	var testURL = "//seg1////seg2"
	var res = TrimPathName(testURL)
	if res != "/seg1/seg2" {
		t.Errorf("trim url %v", testURL)
		t.Errorf("expected %v", "/seg1/seg2")
		t.Errorf("     got %v", res)
	}
}

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	var length = 32
	var token = RandStringBytesMaskImprSrc(length)
	if len(token) != length {
		t.Errorf("expected string length %v", length)
		t.Errorf("                   got %v", len(token))
	}
	var reg = regexp.MustCompile(`^[\da-zA-Z]{` + strconv.Itoa(length) + `}$`)
	var succ = reg.Match([]byte(token))
	if !succ {
		t.Errorf("unexpected token %v", token)
	}
}
