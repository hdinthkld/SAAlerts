package main

type device struct {
	id           string
	hostname     string
	software_ver string
	hw_model     string
	cpes         []string
}

func (d device) append() {

}
