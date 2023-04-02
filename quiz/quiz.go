package quiz

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func readCsvToSlice(slice [][]string) [][]string {
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

	for i := 0; i < len(slice); i++ {
		lastIndex := len(slice[i]) - 1
		if slice[i][lastIndex] == "t" {
			trueCounter++
			continue
		}
		falseCounter++
	}

	return fmt.Sprintf("Correct answers: %d, Wrong answers: %d", trueCounter, falseCounter)
}

func timeout(channel chan<- bool, seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
	channel <- true
}

func Quiz() {
	timeoutFlag := flag.Int64("timeout", 15, "timeout to race against in seconds")
	flag.Parse()

	var slice [][]string
	timeoutChannel := make(chan bool)
	slice = readCsvToSlice(slice)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Quiz Program")
	fmt.Println("---------------------")
	fmt.Printf("You have %ds to answer questions press <enter> to start\n", *timeoutFlag)
	reader.ReadString('\n')

consoleLoop:
	for i := 0; i < len(slice); i++ {

		go timeout(timeoutChannel, *timeoutFlag)

		select {
		case timeoutMessage := <-timeoutChannel:
			if timeoutMessage {
				break consoleLoop
			}
		default:

		}

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
