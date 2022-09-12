package produce

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// function to create .csv reconcile report
func CreateReconcilReport(file_name string, reconcil_data []string) {
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
func DateRangeReport(own_data []map[string]string) string {
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

	return "Date From " + date_from.Format("02-Jan-2006") + " to " + date_to.Format("02-Jan-2006")
}

// function to create summary report after recouncile data between bank_data and own_data
func CreateSummaryReport(file_name string, bank_data []map[string]string, own_data []map[string]string, discrepancies map[string]int) {
	date_range := DateRangeReport(own_data)
	f, err := os.Create(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Data Recouncile Summary Report\nDate Range : " + date_range + "\nNumber of Source Record Processed, " + strconv.Itoa(len(own_data)) + " for Own Data & " + strconv.Itoa(len(bank_data)) + " Bank Data\nData Discrepancies are " + strconv.Itoa(discrepancies["counter"]) + " consisting of " + strconv.Itoa(discrepancies["miss"]) + " missing data, " + strconv.Itoa(discrepancies["amount"]) + " differs amount, " + strconv.Itoa(discrepancies["descr"]) + " differs description, " + strconv.Itoa(discrepancies["date"]) + " differs date")

	if err2 != nil {
		log.Fatal(err2)
	}
}
