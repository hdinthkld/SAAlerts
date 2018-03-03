package nvdcve

import (
	"os"
)

var httpSetup HTTPSettings

type CVEDownloadError struct {
}

type HTTP struct {
	Settings HTTPSettings
}

// HTTPSettings is a structure will enables the various paramters used in downloading the CVEs to be modified
type HTTPSettings struct {
	BaseURL      string
	baseName     string
	suffix       string
	minCVEYear   int
	modifiedName string
	updatedName  string
	useCache     bool // Not yet implemented
}

// DefaultHTTPSettings provides a default set of HTTP parameters which is use to retrieve the CVE files
var DefaultHTTPSettings = HTTPSettings{
	BaseURL:      "https://static.nvd.nist.gov/feeds/xml/cve/2.0/",
	baseName:     "nvdcve-2.0",
	suffix:       ".xml.gz",
	minCVEYear:   2002,
	modifiedName: "modified",
	updatedName:  "updated",
	useCache:     false, // Not yet implemented
}

func New() *HTTP {
	h := new(HTTP)
	h.Settings = DefaultHTTPSettings
	return h
}

/*
DownloadAllCVEs is a convenience function that will download all the CVEs ever published
to the path specified by basename which must exist before running the function.
*/
func (h HTTP) DownloadAllCVEs(path string, basename string) (err int) {
	return 0
}

/*
DownloadCVEsLatest is a convenience function that will attempt to download only the most recently
published CVEs. It calls the DownloadCves function for the 'latest' file whilst contains all
updated CVEs for the last 7 days.
*/
func (h HTTP) DownloadCVEsLatest(path string, basename string) (err int) {
	return 0
}

/*
DownloadCVEsByYear function takes the year in which to download CVEs as an argument
and attempts to download the XML from MITRE server using HTTP.
*/
func (h HTTP) DownloadCVEsByYear(path string, basename string, year int) (err int) {
	return 0
}

/*
DownloadCVEs is ...
*/
func (h HTTP) DownloadCVEs(path string, filename string, cveurl string) (err error) {
	var e error

	if _, err := os.Stat(path); os.IsNotExist(err) {
		e = new(CVEDownloadError)
		return e
	}

	return e

}

func (e CVEDownloadError) Error() string {
	return "Could not download CVE entries"
}
