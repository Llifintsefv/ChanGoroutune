package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) 

	ChanString := make(chan string, 10)
	ChanInt := make(chan string, 10)
	ChanFloat := make(chan string, 10)

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup
	var writeWG sync.WaitGroup 

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go GenerateString(&wg, ChanString)
	}

	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go GenerateInt(&wg2, ChanInt)
	}

	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go GenerateIntFloat(&wg3, ChanFloat)
	}


	writeWG.Add(3) 
	go WriteString(ChanString, &writeWG)
	go WriteInt(ChanInt, &writeWG)
	go WriteFloat(ChanFloat, &writeWG)

	wg.Wait()
	close(ChanString)
	wg2.Wait()
	close(ChanInt)
	wg3.Wait()
	close(ChanFloat)

	writeWG.Wait() 
	fmt.Println("aboba")
}

func GenerateString(wg *sync.WaitGroup, ChanString chan<- string) {
	defer wg.Done()
	stringArray := []string{"apple ", "banana ", "cherry ", "date "}
	var generatedString string
	randomIteration := rand.Intn(5) + 1
	for i := 0; i < randomIteration; i++ {
		generatedString += stringArray[rand.Intn(len(stringArray))]
	}
	ChanString <- generatedString
}

func GenerateInt(wg *sync.WaitGroup, ChanInt chan<- string) {
	defer wg.Done()
	intArray := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var generatedInt string
	randomIteration := rand.Intn(5) + 1
	for i := 0; i < randomIteration; i++ {
		generatedInt += intArray[rand.Intn(len(intArray))]
	}
	ChanInt <- generatedInt
}

func GenerateIntFloat(wg *sync.WaitGroup, ChanFloat chan<- string) {
	defer wg.Done()
	floatArray := []string{"1.1", "2.2", "3.3", "4.4", "5.5", "6.6", "7.7", "8.8", "9.9", "10.10"}
	var generatedFloat string
	randomIteration := rand.Intn(5) + 1
	for i := 0; i < randomIteration; i++ {
		generatedFloat += floatArray[rand.Intn(len(floatArray))]
	}
	ChanFloat <- generatedFloat
}

func WriteString(ChanString <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, _ := os.OpenFile("text.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for genString := range ChanString {
		_, _ = file.WriteString(genString)
	}
}

func WriteInt(ChanInt <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, _ := os.OpenFile("int.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for genInt := range ChanInt {
		_, _ = file.WriteString(genInt)
	}
}

func WriteFloat(ChanFloat <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, _ := os.OpenFile("float.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for genFloat := range ChanFloat {
		_, _ = file.WriteString(genFloat)
	}
}
