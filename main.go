package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// set all discrepancies counter variable
var disc_counter, disc_miss, disc_amount, disc_date, disc_desc int

// function for open and readfile, and return array of string
func readFile(file_name string) []string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	return text
}

// function for mapping array of string that already return from readFile to be array of map data
func mappingArrayOfData(array_data []string) (map[string][]map[string]string, string) {
	var data_header = array_data[0]
	var each_header = strings.Split(data_header, ",")
	mapping_data := make(map[string][]map[string]string)

	//mapping slice of string from array string
	for i := 1; i < len(array_data); i++ {
		var each_value = strings.Split(array_data[i], ",")
		aMap := make(map[string]string)
		for i, s := range each_value {
			aMap[each_header[i]] = s
		}
		mapping_data["data"] = append(mapping_data["data"], aMap)
	}
	return mapping_data, data_header
}

// function to read and mapping data, then return array of map data list
func setDataList(filename string) ([]map[string]string, string) {
	dataread := readFile(filename)
	setdata, stringHeader := mappingArrayOfData(dataread)

	return setdata["data"], stringHeader
}

// function to check if a value exist in the map
func isExistValueCheck(own_id string, bank_data []map[string]string) (bool, map[string]string) {
	//check each bank_data, are contains ID that same with own_id or not
	for _, value := range bank_data {
		if value["ID"] == own_id {
			return true, value
		}
	}
	return false, nil
}

// function to check, is the data form same or not
func checkDataForm(own_data map[string]string, bank_data map[string]string) string {
	// Remarks_note
	var remark_note string
	if own_data["Amt"] != bank_data["Amount"] {
		remark_note = remark_note + "&differs-on-amount" + "(Bank=" + bank_data["Amount"] + ")"
		disc_counter++
		disc_amount++
	}

	if own_data["Descr"] != bank_data["Description"] {
		remark_note = remark_note + "&differs-on-desc" + "(Bank=" + "/" + bank_data["Description"] + ")"
		disc_counter++
		disc_desc++
	}

	if own_data["Date"] != bank_data["Date"] {
		remark_note = remark_note + "&differs-on-date" + "(Bank=" + "/" + bank_data["Date"] + ")"
		disc_counter++
		disc_date++
	}

	return remark_note
}

// function produce every new line for reconcil data result
func produceNewReconcileData(each_header []string, value map[string]string) string {
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
func reconcilData(bank_data []map[string]string, own_data []map[string]string, own_header string) []string {
	var reconcil_results []string

	reconcil_results = append(reconcil_results, own_header+",Remarks") //add new header "Remarks"

	var each_header = strings.Split(reconcil_results[0], ",")
	for _, value := range own_data {
		isExist, exist_data_bank := isExistValueCheck(value["ID"], bank_data)
		if isExist {
			value["Remarks"] = "Found" //if value exist remarks as "Found"
			new_remark_note := checkDataForm(value, exist_data_bank)
			value["Remarks"] = value["Remarks"] + new_remark_note
		} else {
			value["Remarks"] = "Missed" //if not exist remarks as "Missed"
			disc_counter++
			disc_miss++
		}

		newline_record := produceNewReconcileData(each_header, value) //processing new line record data
		reconcil_results = append(reconcil_results, newline_record)
	}

	return reconcil_results
}

// function to create .csv reconcile report
func createReconcilReport(file_name string, reconcil_data []string) {
	csvFile, err := os.Create(file_name)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, value := range reconcil_data {
		_ = csvwriter.Write(strings.Split(value, ","))
	}

	csvwriter.Flush()
	csvFile.Close()
}

// function to check and set data range report
func dateRangeReport(own_data []map[string]string) string {
	var date_from, date_to time.Time
	for i, value := range own_data {
		ownDate, _ := time.Parse("2006-01-02", value["Date"])
		if i == 0 {
			date_from = ownDate
			date_to = ownDate
		}

		if ownDate.Before(date_from) {
			date_from = ownDate
		}

		if ownDate.After(date_to) {
			date_to = ownDate
		}
	}

	return "Date From " + date_from.String() + " to " + date_to.String()
}

// function to create summary report after recouncile data between bank_data and own_data
func createSummaryReport(file_name string, bank_data []map[string]string, own_data []map[string]string) {
	date_range := dateRangeReport(own_data)
	f, err := os.Create(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Data Recouncile Summary Report\nDate Range : " + date_range + "\nNumber of Source Record Processed, " + strconv.Itoa(len(own_data)) + " for Own Data & " + strconv.Itoa(len(bank_data)) + " Bank Data\nData Discrepancies are " + strconv.Itoa(disc_counter) + " consisting of " + strconv.Itoa(disc_miss) + " missing data, " + strconv.Itoa(disc_amount) + " differs amount, " + strconv.Itoa(disc_desc) + " differs description, " + strconv.Itoa(disc_date) + " differs date")

	if err2 != nil {
		log.Fatal(err2)
	}
}

func main() {
	// set bank data
	bank_data, _ := setDataList("source.csv")

	//set own data
	own_data, own_header := setDataList("proxy.csv")

	//reconciliation own data towards bank data
	reconcil_data := reconcilData(bank_data, own_data, own_header)

	//create reconcil report in .csv file
	createReconcilReport("answer_a.csv", reconcil_data)

	//create summary report in .txt file
	createSummaryReport("answer_b.txt", bank_data, own_data)
}
