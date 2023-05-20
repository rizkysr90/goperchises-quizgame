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
	var answer string
	trueAnswers := 0
	for idx, value := range result {
		
		fmt.Printf("Problem #%d : %s = \n", idx + 1, value[0])
		fmt.Scanf("%s\n", &answer)

		if (answer == value[1]) {
			trueAnswers++
		}
	}; 
	fmt.Printf("Anda benar %d dari %d soal \n", trueAnswers, len(result))

}