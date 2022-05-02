package main

import (
	"math/rand"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	//flag.String returns a pointer to string
	csvFile := flag.String("csv", "problems.csv", "provide a csv file")
	timeDuration := flag.Int("time",30,"set time duration to answer the questions")
	quizShuffle := flag.Bool("shuffle",false,"set true to shuffle order of questions")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()

	fmt.Println("This is a timed test. Press Enter to start")
	fmt.Scanln()

	
	if *quizShuffle == true{
		rand.Shuffle(len(lines),func(i,j int){
			lines[i] , lines[j] = lines[j] , lines[i]
		})
	}
	
	startQuiz(lines,timeDuration)
}

func startQuiz(lines [][]string,timeDuration *int)  {
	timer := time.NewTimer(time.Duration(*timeDuration) * time.Second)
	corrCount := 0
	for i, v := range lines {
		fmt.Printf("#Problem %d: %s = ", i+1, v[0])
		scanCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			scanCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n\n", corrCount, len(lines))
			return
		case ans := <-scanCh:
			if ans == strings.TrimSpace(v[1]) {
				corrCount++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d\n\n", corrCount, len(lines))
}
