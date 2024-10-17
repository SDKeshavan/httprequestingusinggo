package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type brecords struct {
	time string
	lat  string
	long string
}

type igcStruct struct {
	manf       string
	date       string
	gliderName string
	gliderid   string
	route      []brecords
}

func main() {

	reader, _ := os.Open("../../IGCFiles/4A9B3591.IGC")

	scan := bufio.NewScanner(reader)

	var igcfile igcStruct

	for scan.Scan() {

		line := scan.Text()

		fmt.Println(line)

		if strings.Compare(line[:1], "A") == 0 {
			igcfile.manf = line[1:4]
		} else if strings.Compare(line[:1], "B") == 0 {

			igcfile.route = append(igcfile.route, addtoRoute(line[1:]))
		} else if strings.Compare(line[:1], "H") == 0 {
			startInd := strings.Index(line, ":")
			if strings.Contains(line, "GLIDERID") {
				igcfile.gliderid = line[startInd+1:]
			} else if strings.Contains(line, "HFPLT") {
				igcfile.gliderName = line[startInd+1:]
			}
		}
	}

	printInfo(igcfile)

	// for i := 0; i < len(igcfile.route); i++ {
	// 	fmt.Println(igcfile.route[i].lat)
	// }

}

func addtoRoute(data string) brecords {
	time := data[:6]
	lat := data[6:14]
	long := data[14:22]
	var brecord brecords
	brecord.lat = lat
	brecord.long = long
	brecord.time = time

	return brecord
}

func printInfo(fileDet igcStruct) {

	fmt.Println("Manufacturer :" + fileDet.manf)
	fmt.Println("Glider Name :" + fileDet.gliderName)
	fmt.Println("Glider ID :" + fileDet.gliderid)

}
