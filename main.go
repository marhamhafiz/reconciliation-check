package main

import (
	"reconciliation-check/prepare"
	"reconciliation-check/produce"
	"reconciliation-check/reconcile"
)

func main() {
	// set bank data
	bank_data, _ := prepare.SetDataList("source.csv")

	//set own data
	own_data, own_header := prepare.SetDataList("proxy.csv")

	//reconciliation own data towards bank data
	reconcil_data, discrepancies := reconcile.ReconcilData(bank_data, own_data, own_header)

	//create reconcil report in .csv file
	produce.CreateReconcilReport("answer_a.csv", reconcil_data)

	//create summary report in .txt file
	produce.CreateSummaryReport("answer_b.txt", bank_data, own_data, discrepancies)
}
