package util

import (
	"testing"
	"github.com/carsonsx/log4g"
)

func TestGetTodayUnix(t *testing.T) {
	b, e := GetTodayInterval()
	log4g.Info(b)
	log4g.Info(e)
}