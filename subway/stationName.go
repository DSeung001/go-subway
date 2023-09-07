package subway

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"subway/utils"
)

const resourceDirPath = "subway/resource"

func getStationList() {
	names := getStationNames()
	for _, name := range names {
		// 파일은 읽어지지만 가져온 역정보 xls에 문제가 있어서 읽혀지지가 않음
		fmt.Println(resourceDirPath + "/" + name)
		file, err := excelize.OpenFile(resourceDirPath + "/" + name)
		utils.HandleErr(err)

		// file 정리
		defer func() {
			// Close the spreadsheet.
			if err := file.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		// 시트명을 역명으로 내용 읽기
		rows, err := file.GetRows("역명")
		utils.HandleErr(err)
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
	}
}

func getStationNames() []string {
	var fileNames []string

	// 폴더 열기
	dir, err := os.Open(resourceDirPath)
	utils.HandleErr(err)
	defer func(dir *os.File) {
		err := dir.Close()
		utils.HandleErr(err)
	}(dir)

	// 디렉토리 내 파일 목록 읽기
	fileInfos, err := dir.Readdir(-1)
	utils.HandleErr(err)

	// 파일 이름을 리스트에 넣기
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	return fileNames
}
