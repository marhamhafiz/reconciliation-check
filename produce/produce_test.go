package produce

import (
	"regexp"
	"testing"
)

// TestDateRange calls prepare.DateRange with a string, checking
// for a valid return value.
func TestDateRange(t *testing.T) {
	map_data := []map[string]string{
		{
			"No":    "1",
			"Place": "tokyo",
			"Date":  "2021-07-14",
		},
		{
			"No":    "2",
			"Place": "paris",
			"Date":  "2021-07-02",
		},
		{
			"No":    "3",
			"Place": "london",
			"Date":  "2021-07-08",
		},
	}
	value_test := "Date From 02-Jul-2021 to 14-Jul-2021" ////should be pass
	// value_test := "Date From 14-Jul-2021 to 08-Jul-2021" //should be error test
	want := regexp.MustCompile(`\b` + value_test + `\b`)
	msg := DateRangeReport(map_data)
	if !want.MatchString(msg) {
		t.Fatalf(`DateRangeReport(map_data) = %q, want match for %#q, nil`, msg, want)
	}
}
