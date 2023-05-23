package exportExcel

import (
	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"strconv"
)

var cells = []string{"A", "B", "C", "D", "E", "F", "J", "K"}

func WriteStreamToExcel(stream string, filePath string, id int) {
	f, err := excelize.OpenFile(filePath)

	if err != nil {
		log.Error(err)
		return
	}

	err = f.Save()
	if err != nil {
		log.Error(err)
		return
	}

	/*	for _, cell := range cells {


		for i := 0; i < len(stream); i = i + 2 {
			cellValue := string(stream[i]) + string(stream[i+1])
			fmt.Println("Cell Value: ", cellValue)
			f.SetCellValue("Sheet1", cellName, cellValue)
		}

	}*/

	cellIndex := -1

	for i := 0; i < len(stream); i = i + 2 {
		cellValue := string(stream[i]) + string(stream[i+1])
		cellIndex++
		var cellName string = cells[cellIndex] + strconv.Itoa(id)
		f.SetCellValue("Sheet1", cellName, cellValue)
	}

	f.Save()

}
