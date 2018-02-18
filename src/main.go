package main

import (
	"fmt"
	"os"
)

var cves []cve
var devices []device

func populateData() {
	var temp_device device
	temp_device = device{
		id:           "CI0001",
		hostname:     "CUST01-DEV01-LOC01",
		software_ver: "9.1.7(22)",
		hw_model:     "Cisco ASA 5525",
		cpes: []string{
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.22:*:*:*:*:*:*:*",
			"cpe:2.3:h:cisco:Adaptive Security Appliance 5525-X:*:*:*:*:*:*:*:*",
		},
	}

	devices = append(devices, temp_device)

	temp_device = device{
		id:           "CI10101",
		hostname:     "CUST02-DEV02-LOC02",
		software_ver: "8.2.1(1)",
		hw_model:     "Cisco ASA 5510",
		cpes: []string{
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.1:*:*:*:*:*:*:*",
			"cpe:2.3:h:cisco:Adaptive Security Appliance 5510:*:*:*:*:*:*:*:*",
		},
	}

	devices = append(devices, temp_device)

	temp_device = device{
		id:           "CI255",
		hostname:     "CUST01-DEV02-LOC01",
		software_ver: "LTM 11.1",
		hw_model:     "1400 Series",
		cpes: []string{
			"cpe:2.3:o:f5:Local Traffic Monitor:11.1:*:*:*:*:*:*:*",
			"cpe:2.3:h:f5:1400 Series Appliance:*:*:*:*:*:*:*:*",
		},
	}

	devices = append(devices, temp_device)

	var temp_cve cve
	temp_cve = cve{
		ref: "CVE-2018-0101",
		vendor_refs: []string{
			"cisco-sa-20180129-asa1",
		},
		title:    "Cisco Adaptive Security Appliance Remote Code Execution and Denial of Service Vulnerability",
		severity: "Critical",
		cpes: []string{
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.1:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.2:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.3:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.4:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.5:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.6:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.7:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.8:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.9:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.10:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:8.2.11:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.1:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.2:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.3:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.4:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.5:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.6:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.7:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.8:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.9:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.10:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.11:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.12:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.13:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.14:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.15:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.16:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.17:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.18:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.19:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.20:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.21:*:*:*:*:*:*:*",
			"cpe:2.3:o:cisco:Adaptive Security Appliance Software:9.1.7.22:*:*:*:*:*:*:*",
		},
	}

	cves = append(cves, temp_cve)
}

func main() {
	populateData()

	var cve_ref string

	fmt.Print("Enter a CVE reference: ")
	fmt.Scanln(&cve_ref)

	var found_cve cve
	found := false

	for _, cve := range cves {
		if cve.ref == cve_ref {
			found = true
			found_cve = cve
			break
		}
	}

	if !found {
		fmt.Println("Unknown CVE Reference: ", cve_ref)
		os.Exit(1)
	}

	// Print out number of affected CPEs
	fmt.Printf("Found %s. There are %d vulnerable products listed\n\n", cve_ref, len(found_cve.cpes))

	// Search for devices of affected CPEs

	var affected_devices []int

	for di, device := range devices {
		for _, cpe := range device.cpes {
			for _, affected_cpe := range found_cve.cpes {
				// fmt.Println(affected_cpe, cpe)
				if affected_cpe == cpe {
					affected_devices = append(affected_devices, di)
				}
			}
		}
	}

	// Print out devices affected
	fmt.Println("The following devices could be vulnerable:")
	for di := range affected_devices {
		fmt.Println(devices[di].hostname)
	}

	fmt.Println("\nDevices may only be vulnerable if they are running affected features.\nExamine configuration to be sure")
}
