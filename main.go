package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	openedFile, err := os.Open("problems.csv")
	if (err != nil) {
		fmt.Println(err)
	}
	r := csv.NewReader(openedFile)
	result, _ := r.ReadAll()

	fmt.Println(result)

}