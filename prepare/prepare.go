package prepare

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// function for open and readfile, and return array of string
func ReadFile(file_name string) []string {
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
func MappingArrayOfData(array_data []string) (map[string][]map[string]string, string) {
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
func SetDataList(filename string) ([]map[string]string, string) {
	dataread := ReadFile(filename)
	setdata, stringHeader := MappingArrayOfData(dataread)

	return setdata["data"], stringHeader
}
