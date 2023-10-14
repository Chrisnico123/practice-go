package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func main() {
	dataCh, err := readFileConcurren("../data.json")

	if err != nil {
		panic(err)
	}
	//Worker
	changeDollarToIdrConcurren(dataCh)
	changeDollarToIdrConcurren(dataCh)
	changeDollarToIdrConcurren(dataCh)
	changeDollarToIdrConcurren(dataCh)
}

func writeFile(data <-chan User, done chan bool) {
	wg := sync.WaitGroup{}

	for user := range data {
		wg.Add(1)
		go func(user User) {
			userByte, _ := json.Marshal(user)
			err := ioutil.WriteFile("users/"+user.Name+".json", userByte, 0666)
			if err != nil {
				log.Println("error when try to write file", err)
			}
			wg.Done()
		}(user)
	}

	go func() {
		wg.Wait()
		done <- true
	}()
}

func changeDollarToIdrConcurren(dataCh <-chan User) {
	now := time.Now()
	outCh := make(chan User)

	go func ()  {
		for data := range dataCh {
			time.Sleep(10 * time.Millisecond)
			newData := data
			newData.Salary = newData.Salary * 10_000

			outCh <- newData
		}
		
		close(outCh)
	}()

	log.Println("success change usd to idr in", time.Since(now).Nanoseconds(), "ns")
	done := make(chan bool)
	writeFile(outCh , done)
}

func readFileConcurren(filename string) (<-chan User , error){
	now := time.Now()
	outCh := make(chan User)
	data , err := os.ReadFile(filename)
	if err != nil {
		return outCh , err
	}

	users := []User{}

	err = json.Unmarshal(data , &users)
	if err != nil {
		return outCh , err
	}

	go func ()  {
		for _, user := range users {
			outCh <- user
		}

		close(outCh)
	}()

	log.Println("success read data in", time.Since(now).Seconds(), "s")

	return outCh , nil
}