package utils

import (
	"errors"
	"testing"

	"github.com/kdada/tinygo/util"
)

func TestInfo(t *testing.T) {
	var succReturn = NewSuccessReturn("succ")
	if succReturn.Code != 0 {
		t.Errorf("expected succ return code %v", 0)
		t.Errorf("                      got %v", succReturn.Code)
	}

	var failReturn = NewFailReturn(errors.New("failed return"))
	if failReturn.Code != 60000 {
		t.Errorf("expected fail return code %v", 60000)
		t.Errorf("                      got %v", failReturn.Code)
	}

	failReturn = NewFailReturn("failed return")
	if failReturn.Code != 60000 {
		t.Errorf("expected fail return code %v", 60000)
		t.Errorf("                      got %v", failReturn.Code)
	}

	failReturn = NewFailReturn(util.ErrorMacAddrNotFound)
	if failReturn.Code != 60000 {
		t.Errorf("expected fail return code %v", 60000)
		t.Errorf("                      got %v", failReturn.Code)
	}

	failReturn = NewFailReturn(11)
	if failReturn.Code != 60001 {
		t.Errorf("expected fail return code %v", 60001)
		t.Errorf("                      got %v", failReturn.Code)
	}
}
