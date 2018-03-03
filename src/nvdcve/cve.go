package nvdcve

type cve struct {
	ref        string
	vendorRefs []string
	title      string
	severity   string
	cpes       []string
	feature    string
}
