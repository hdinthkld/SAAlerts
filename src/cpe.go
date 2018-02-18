package main

type CPE struct {
	// The WFN used to identify this specific CPE - cpe:2.3:a:microsoft:internet_explorer:8.0.6001:beta:*:*:*:*:*:*
	id      string
	part    string // a = Application, o = Operating System, h = Hardware
	vendor  string // Manufactuer of the product
	product string // The product iteself
	version string // Vendor specific release of the product that can be determined from discoverable data
	update  string // Vendor specific service pack, update or patch of the product
}
