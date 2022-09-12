package reconcile

import (
	"regexp"
	"strconv"
	"testing"
)

// TestIsExistValueCheck calls prepare.IsExistValueCheck with a string, checking
// for a valid return value.
func TestIsExistValueCheck(t *testing.T) {
	bank_data := []map[string]string{
		{
			"ID":          "1",
			"Amount":      "123123",
			"Description": "tokyo",
			"Date":        "2021-07-14",
		},
		{
			"ID":          "2",
			"Amount":      "45646",
			"Description": "paris",
			"Date":        "2021-07-02",
		},
		{
			"ID":          "3",
			"Amount":      "546443",
			"Description": "london",
			"Date":        "2021-07-08",
		},
	}
	id_test := "2" ////should be pass
	// id_test := "7" //should be error test

	foud_test := "true"
	value_test := "paris" ////should be pass id_test id = "2"
	// value_test := "jakarta" ////should be error test
	want_id := regexp.MustCompile(`\b` + foud_test + `\b`)
	want_value := regexp.MustCompile(`\b` + value_test + `\b`)
	found, value := IsExistValueCheck(id_test, bank_data)
	if !want_id.MatchString(strconv.FormatBool(found)) || !want_value.MatchString(value["Description"]) {
		t.Fatalf(`IsExistValueCheck = %q, want_map match for %#q and %q, want_header match for %#q nil`, strconv.FormatBool(found), want_id, value["Description"], want_value)
	}
}

// TestCheckDataForm calls prepare.IsCheckDataFormk with a string, checking
// for a valid return value.
func TestCheckDataForm(t *testing.T) {
	own_data := map[string]string{
		"ID":    "1",
		"Amt":   "123123",
		"Descr": "jakarta",
		"Date":  "2021-07-14",
	}

	bank_data := map[string]string{
		"ID":          "1",
		"Amount":      "123123",
		"Description": "tokyo",
		"Date":        "2021-07-14",
	}

	value_test := "tokyo" ////should be pass
	// value_test := "jakarta" ////should be error test
	want_value := regexp.MustCompile(`\b` + value_test + `\b`)
	note := CheckDataForm(own_data, bank_data)
	if !want_value.MatchString(note) {
		t.Fatalf(`IsExistValueCheck = %q, want_map match for %#q nil`, note, want_value)
	}
}

// TestProduceNewReconcileData calls prepare.ProduceNewReconcileData with a string, checking
// for a valid return value.
func TestProduceNewReconcileData(t *testing.T) {
	each_header := []string{
		"ID",
		"Amt",
		"Descr",
		"Date",
	}

	new_data := map[string]string{
		"ID":    "1",
		"Amt":   "123123",
		"Descr": "tokyo",
		"Date":  "2021-07-14",
	}

	value_test := "1,123123,tokyo,2021-07-14" ////should be pass
	// value_test := "1,,tokyo,2021-07-14" ////should be error test
	want_value := regexp.MustCompile(`\b` + value_test + `\b`)
	newline := ProduceNewReconcileData(each_header, new_data)
	if !want_value.MatchString(newline) {
		t.Fatalf(`ProduceNewReconcileData = %q, want_map match for %#q nil`, newline, want_value)
	}
}

// TestReconcilData calls prepare.ReconcilData with a string, checking
// for a valid return value.
func TestReconcilData(t *testing.T) {
	bank_data := []map[string]string{
		{
			"ID":          "1",
			"Amount":      "123123",
			"Description": "tokyo",
			"Date":        "2021-07-14",
		},
		{
			"ID":          "2",
			"Amount":      "45646",
			"Description": "paris",
			"Date":        "2021-07-02",
		},
		{
			"ID":          "3",
			"Amount":      "546443",
			"Description": "jakarta",
			"Date":        "2021-07-08",
		},
	}

	own_data := []map[string]string{
		{
			"ID":    "1",
			"Amt":   "123123",
			"Descr": "tokyo",
			"Date":  "2021-07-14",
		},
		{
			"ID":    "2",
			"Amt":   "45646",
			"Descr": "paris",
			"Date":  "2021-07-02",
		},
		{
			"ID":    "3",
			"Amt":   "546",
			"Descr": "jakarta",
			"Date":  "2021-07-08",
		},
	}

	own_header := "ID,Amt,Descr,Date"

	// value_test := "1,123123,tokyo,2021-07-14,Found" ////should be pass
	value_test := "1,123123,tokyo,2021-07-14,Missed" ////should be error test

	disc_test := "1" ////should be pass
	// disc_test := "0" ////should be error test
	want_value := regexp.MustCompile(`\b` + value_test + `\b`)
	want_disc := regexp.MustCompile(`\b` + disc_test + `\b`)
	newline, disc := ReconcilData(bank_data, own_data, own_header)

	//you can check with disc with amount, date, remarks
	//if use descr made the value +1 cause of case in tes TestCheckDataForm has been counter it once
	if !want_value.MatchString(newline[1]) || !want_disc.MatchString(strconv.Itoa(disc["amount"])) {
		t.Fatalf(`ReconcilData = %q, want_value match for %#q nil and %q, want_disc match for %#q nil`, newline[1], want_value, strconv.Itoa(disc["amount"]), want_disc)
	}
}
