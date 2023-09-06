package subway

import (
	"os"
	"subway/utils"
)

const resourceDirPath = "subway/resource"

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
