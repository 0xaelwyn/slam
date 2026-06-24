package webserver

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var resolvedTimezone = "UTC"

// SetTimezone sets the timezone used in resolved alert timestamps.
func SetTimezone(tz string) {
	resolvedTimezone = tz
}

// Hash returns the sha256 for  string
func Hash(key string) string {
	h := sha256.New()
	// hash.Hash.Write never returns an error.
	//nolint: errcheck
	h.Write([]byte(string(key)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func timeNowToDateTimeFormatted() string {
	loc, err := time.LoadLocation(resolvedTimezone)
	if err != nil {
		loc = time.UTC
	}
	return time.Now().In(loc).Format("Jan 2, 2006 at 3:04 PM")
}
