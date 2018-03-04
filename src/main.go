package main

import (
	"fmt"
	"nvdcve"
	// "strconv"
)

func main() {
	nh := nvdcve.New()
	nh.Settings.BaseURL = "http://127.1.1.1/feeds/xml/cve/2.0/"
	fmt.Println(nh.Settings.MaxCVEYear)

	// Try to download all files to the specified directory and file prefix
	err := nh.DownloadAllCVEs("c:/temp", "nvdcve-2.0")
	fmt.Println(err)
}
