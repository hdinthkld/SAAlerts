package main

import (
	"fmt"
	"nvdcve"
)

func main() {
	nh := nvdcve.New()
	err := nh.DownloadCVEs("c:\\temp", "test", nh.Settings.BaseURL+"-"+string(2018)+".xml.gz")
	fmt.Println(err)
}
