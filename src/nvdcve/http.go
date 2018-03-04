package nvdcve

import (
	"io"
	"net/http"
	"os"
	"strings"
	"strconv"
	"time"
)

var httpSetup HTTPSettings

type CVEDownloadError struct {
	Url string
}

type HTTP struct {
	Settings HTTPSettings
}

// HTTPSettings is a structure will enables the various paramters used in downloading the CVEs to be modified
type HTTPSettings struct {
	BaseURL      string
	BaseName     string
	Suffix       string
	MinCVEYear   int
	MaxCVEYear	int
	ModifiedName string
	RecentName  string
}

// DefaultHTTPSettings provides a default set of HTTP parameters which is use to retrieve the CVE files
var DefaultHTTPSettings = HTTPSettings{
	BaseURL:      "https://static.nvd.nist.gov/feeds/xml/cve/2.0/",
	BaseName:     "nvdcve-2.0",
	Suffix:       ".xml.gz",
	MinCVEYear:   2002,
	MaxCVEYear:   time.Now().Year(),
	ModifiedName: "modified",
	RecentName:   "recent",
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
func (h HTTP) DownloadAllCVEs(path string, basename string) (err error) {

	var years []string // Will store all the CVE years that we need to download

	// Build a slice of all CVE years that need to be downloaded
	for y:=h.Settings.MinCVEYear; y <= h.Settings.MaxCVEYear; y++ {
		years = append(years, strconv.Itoa(y))
	}

	// Add the "special" files that are not year based
	years = append(years, h.Settings.RecentName)
	years = append(years, h.Settings.ModifiedName)

	// Loops over all the defined years and downloads the file from the defined web server
	for _,d:= range(years) {
		fn := basename + "-" + d + h.Settings.Suffix
		cveurl := h.Settings.BaseURL +  h.Settings.BaseName + "-" + d + ".xml.gz"

		err := h.DownloadCVEs(path, fn, cveurl)

		if err != nil {
			return err
		}
	}

	return nil
}

/*
DownloadCVEs takes the give URL and attempts to download the file and store it with the path and
filename specified.
*/
func (h HTTP) DownloadCVEs(path string, filename string, cveurl string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	resp, err := http.Get(cveurl)

	if err != nil {
		// TODO: Improve error routine
		return err
	}

	defer resp.Body.Close()

	f, err := os.Create(strings.Join([]string{path, filename}, "/"))

	if err != nil {
		// TODO: Improve error routine
		return err
	}

	defer f.Close()

	io.Copy(f, resp.Body)

	return nil

}

