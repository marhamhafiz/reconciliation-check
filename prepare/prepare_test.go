package prepare

import (
	"regexp"
	"testing"
)

// TestReadFile calls prepare.ReadFile with a string, checking
// for a valid return value.
func TestReadFile(t *testing.T) {
	value_test := "Name,Age,Number" ////should be pass
	// value_test := "Diaz,21,10" //should be error test
	want := regexp.MustCompile(`\b` + value_test + `\b`)
	msg := ReadFile("sample.csv")
	if !want.MatchString(msg[0]) {
		t.Fatalf(`ReadFile("sample.csv") = %q, want match for %#q, nil`, msg, want)
	}
}

// TestMappingArrayOfData calls prepare.MappingArrayOfData with a []string, checking
// for a valid return value.
func TestMappingArrayOfData(t *testing.T) {
	read_value := []string{"no,place,verif", "1,new york,t", "2,tokyo,f"}
	map_test := map[string]string{
		"no":    "2",
		"place": "tokyo",
		"verif": "f",
	}
	header_test := "no,place,verif" //should be pass
	// header_test := "no,place,verif" //should be error test

	want_map := regexp.MustCompile(`\b` + map_test["place"] + `\b`)
	want_header := regexp.MustCompile(`\b` + header_test + `\b`)
	map_data, header := MappingArrayOfData(read_value)
	if !want_map.MatchString(map_data["data"][1]["place"]) || !want_header.MatchString(header) {
		t.Fatalf(`MappingArrayOfData = %q, want_map match for %#q and %q, want_header match for %#q nil`, map_data["data"][1]["place"], want_map, header, want_header)
	}
}

// TestSetDataList calls prepare.SetDataList with a string, checking
// for a valid return value.
func TestSetDataList(t *testing.T) {
	map_test := "23" //should be pass
	// map_test := "19" //should be error test

	header_test := "Name,Age,Number" //should be pass
	// header_test := "no,place,verif" //should be error test

	want_map := regexp.MustCompile(`\b` + map_test + `\b`)
	want_header := regexp.MustCompile(`\b` + header_test + `\b`)
	map_data, header := SetDataList("sample.csv")
	if !want_map.MatchString(map_data[0]["Age"]) || !want_header.MatchString(header) {
		t.Fatalf(`SetDataList("sample.csv") = %q, want_map match for %#q and %q, want_header match for %#q nil`, map_data[0]["Age"], want_map, header, want_header)
	}
}
