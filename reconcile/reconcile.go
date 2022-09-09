package reconcile

import (
	"strings"
)

// set all discrepancies counter variable
// var disc_counter, disc_miss, disc_amount, disc_date, disc_desc int
var discrepanciies = map[string]int{
	"counter": 0,
	"miss":    0,
	"amount":  0,
	"descr":   0,
	"date":    0,
}

// function to check if a value exist in the map
func IsExistValueCheck(own_id string, bank_data []map[string]string) (bool, map[string]string) {
	//check each bank_data, are contains ID that same with own_id or not
	for _, value := range bank_data {
		if value["ID"] == own_id {
			return true, value
		}
	}
	return false, nil
}

// function to check, is the data form same or not
func CheckDataForm(own_data map[string]string, bank_data map[string]string) string {
	// Remarks_note
	var remark_note string
	if own_data["Amt"] != bank_data["Amount"] {
		remark_note = remark_note + "&differs-on-amount" + "(Bank=" + bank_data["Amount"] + ")"
		discrepanciies["counter"]++
		discrepanciies["amount"]++
		// disc_counter++
		// disc_amount++
	}

	if own_data["Descr"] != bank_data["Description"] {
		remark_note = remark_note + "&differs-on-desc" + "(Bank=" + "/" + bank_data["Description"] + ")"
		discrepanciies["counter"]++
		discrepanciies["descr"]++
		// disc_counter++
		// disc_desc++
	}

	if own_data["Date"] != bank_data["Date"] {
		remark_note = remark_note + "&differs-on-date" + "(Bank=" + "/" + bank_data["Date"] + ")"
		discrepanciies["counter"]++
		discrepanciies["date"]++
		// disc_counter++
		// disc_date++
	}

	return remark_note
}

// function produce every new line for reconcil data result
func ProduceNewReconcileData(each_header []string, value map[string]string) string {
	var temp_new_line string
	for i := range each_header {
		if i == 0 {
			temp_new_line = value[each_header[i]]
		} else {
			temp_new_line = temp_new_line + "," + value[each_header[i]]
		}
	}

	return temp_new_line
}

// reconciliation check between bank_data and own_data
func ReconcilData(bank_data []map[string]string, own_data []map[string]string, own_header string) ([]string, map[string]int) {
	var reconcil_results []string

	reconcil_results = append(reconcil_results, own_header+",Remarks") //add new header "Remarks"

	var each_header = strings.Split(reconcil_results[0], ",")
	for _, value := range own_data {
		isExist, exist_data_bank := IsExistValueCheck(value["ID"], bank_data)
		if isExist {
			value["Remarks"] = "Found" //if value exist remarks as "Found"
			new_remark_note := CheckDataForm(value, exist_data_bank)
			value["Remarks"] = value["Remarks"] + new_remark_note
		} else {
			value["Remarks"] = "Missed" //if not exist remarks as "Missed"
			discrepanciies["counter"]++
			discrepanciies["miss"]++
		}

		newline_record := ProduceNewReconcileData(each_header, value) //processing new line record data
		reconcil_results = append(reconcil_results, newline_record)
	}

	return reconcil_results, discrepanciies
}
