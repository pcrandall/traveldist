package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func findWorkbook() {
	// find current workbook
	re := regexp.MustCompile(`^([a-zA-Z0-9\s_\\.\-\(\):])+\.(xlsx|xlsm)$`)
	files, err := ioutil.ReadDir(".")
	checkErr(err, "")

	for _, file := range files {
		match := re.FindStringSubmatch(file.Name())

		if len(match) > 0 {
			dest := "./old/" + match[0]

			err = Copy(match[0], dest)

			if err != nil {
				log.Println("Latest travel distance file must be in the same directory as Travel Distance.exe")
				log.Panic(err)
			}
			excelColumn, err = findLastColumn(match[0], config.SheetName)
			writeFileName = match[0]
		}
	}

	excelFile, err = excelize.OpenFile(writeFileName)
	if err != nil {
		log.Println("Latest travel distance file must be in the same directory as Travel Distance.exe")
		log.Panic(err)
		return
	}
}

func buildCoordinate(n int, row int) string {
	// https://www.geeksforgeeks.org/find-excel-column-name-given-number
	str, result := "", ""
	for n > 0 {
		rem := n % 26
		if rem == 0 {
			str += string(90) // string(90) = "Z"
			n = (n / 26) - 1
		} else {
			str += string(rem + 64)
			n /= 26
		}
	}
	for _, v := range str {
		result = string(v) + result
	}
	// fmt.Println("coordinate here: ", result+strconv.Itoa(row))
	return result + strconv.Itoa(row)
}

func findLastColumn(inFile string, sheetName string) (write int, err error) {
	s, err := excelize.OpenFile(inFile)

	if err != nil {
		return -1, err
	}
	rows := s.GetRows(sheetName)
	// fmt.Println("row here: ", rows[5])
	// fmt.Println("row here: ", len(rows[5]))
	val := ""
	row := rows[5]

	for i := len(row) - 1; val == ""; i-- {
		val = string(row[i])
		if val != "" {
			excelColumn = i + 1
		}
	}
	excelColumn++
	// fmt.Println("writeColumn here: ", writeColumn)
	// fmt.Println("writeColumn here: ", buildCoordinate(writeColumn, 6))
	return excelColumn, err
}
