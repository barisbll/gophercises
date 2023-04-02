package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readCsvToSlice(slice [][]string)[][]string {
	csvfile, err := os.Open("quiz/problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		slice = append(slice, record)
	}

	return slice
}

func giveResults(slice [][]string) string {
	trueCounter := 0
	falseCounter := 0
	
	for i := 0; i < len(slice) ; i++  {
		if (slice[i][2] == "t") {
			trueCounter++
			continue
		}
		falseCounter++
	}

	
	return fmt.Sprintf("Correct answers: %d, Wrong answers: %d", trueCounter, falseCounter)
}

func Quiz() {
	var slice [][]string
	slice = readCsvToSlice(slice)


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Quiz Program")
	fmt.Println("---------------------")
	  
	for i := 0; i < len(slice) ; i++  {
	  
	  fmt.Printf("What is %s \n", slice[i][0])
	  fmt.Print("-> ")
	  text, _ := reader.ReadString('\n')
	  text = strings.Replace(text, "\n", "", -1)
	  
	  if strings.Compare(slice[i][1], text) == 0 {
		slice[i] = append(slice[i], "t")
		} else {
		  slice[i] = append(slice[i], "f")
	  }
	}

	fmt.Println(giveResults(slice))
}