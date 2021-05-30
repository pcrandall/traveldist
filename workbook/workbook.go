package workbook

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	excelColumn   int
	excelFile     *excelize.File
	oldFilename   string
	writeFileName string
	err           error
)

func checkErr(err error, str string) {
	if err != nil {
		if str != "" {
			log.Println(str, err)
		} else {
			log.Println(err)
		}
	}
}

func WriteFile(row string, sheetname string, total byte) {
	excelRow, _ := strconv.Atoi(row)
	writeCoord := BuildCoordinate(excelColumn, excelRow)
	excelFile.SetCellValue(sheetname, writeCoord, total)
}

func FindWorkbook(sheetname string) {

	newpath := filepath.Join(".", "old")
	if _, err = os.Stat(newpath); os.IsNotExist(err) {
		os.MkdirAll(newpath, os.ModePerm)
	}

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
			excelColumn, err = FindLastColumn(match[0], sheetname)
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

func SaveFile() {
	oldFilename = writeFileName // store filename for comparison later
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		log.Panic(err)
	}

	if writeFileName == "" { // check that the new filename is not empty, if so use old filename
		err := excelFile.SaveAs(oldFilename)
		checkErr(err, "Saving file error")
	} else {
		err := excelFile.SaveAs(writeFileName)
		checkErr(err, "Saving file error")
		err = os.Remove(oldFilename)
		checkErr(err, "Removing old file error")
	}
}

func BuildCoordinate(n int, row int) string {
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

func FindLastColumn(inFile string, sheetName string) (write int, err error) {
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

// Copy the src file to dst. Any existing file will be overwritten and will not copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
