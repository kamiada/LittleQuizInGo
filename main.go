package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in format, question - answer")
	flag.Parse()
	_ = csvFilename
}
