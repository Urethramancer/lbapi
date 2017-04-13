// Utility functions used while parsing the somewhat convoluted
// JSON returned by LogicBoxes API calls.
package lbapi

import (
	"fmt"
	"strconv"
	"time"
)

func parseBool(b string) bool {
	x, err := strconv.ParseBool(b)
	if err != nil {
		return false
	}
	return x
}

func atoi(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func parseDate(s string) time.Time {
	var year, mon, day int
	var hour, min, sec, ns int
	var z int
	c, err := fmt.Sscanf(s, "%04d-%02d-%02d %02d:%02d:%02d.%d+%d", &year, &mon, &day, &hour, &min, &sec, &ns, &z)
	if c != 8 || err != nil {
		return time.Now()
	}
	return time.Date(year, time.Month(mon), day, hour, min, sec, ns, time.UTC)
}
