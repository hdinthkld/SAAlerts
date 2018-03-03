package main

type device struct {
	id              string
	hostname        string
	softwareVersion string
	hwModel         string
	cpes            []string
	features        []string
}

func (d device) append() {

}
