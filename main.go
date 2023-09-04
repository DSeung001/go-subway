package main

import (
	"subway/subway"
)

// 지하철 역별 데이터
// https://data.kric.go.kr/rips/M_01_01/detail.do?id=1

func main() {
	port := 8080

	subway.Start(port)
}
