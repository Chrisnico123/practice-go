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

func writeFile(data <-chan User , done chan bool) {
	wg := sync.WaitGroup{}

	for user := range data {
		wg.Add(1)
		go func(user User) {
			userByte , _ := json.Marshal(user)
			err := ioutil.WriteFile("users/"+user.Name+".json" , userByte , 0666)
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

func changeDollarToIdr(users []User) (newUser []User) {
	now := time.Now()
	for _, user := range users {
		user.Salary *= 100
		newUser = append(newUser, user)
	}

	log.Println("Succes Change Salary Data to idr " , time.Since(now).Seconds() , "s")

	return newUser
}

func mergeData(dataCh ...<-chan User) <-chan User {
	outCh := make(chan User)

	wg := sync.WaitGroup{}

	for _, data := range dataCh {
		wg.Add(1)
		go func(data <-chan User) {
			for d := range data {
				outCh <- d
			}
			wg.Done()
		}(data)
	}

	go func() {
		wg.Wait()	
		close(outCh)
	}()

	return outCh
}

func changeDollarToIdrConcurren(dataCh <-chan User) <-chan User {
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

	return outCh
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

func readFile(filename string) (users []User , err error) {
	data , err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	users = []User{}

	err = json.Unmarshal(data , &users)

	if err != nil {
		return nil , err
	}

	return users , nil
}

func asynchornous() {
	dataCh , err := readFileConcurren("data.json")

	if err != nil {
		panic(err)
	}

	
	users := mergeData(
		// Worker
		changeDollarToIdrConcurren(dataCh),
		changeDollarToIdrConcurren(dataCh),
		changeDollarToIdrConcurren(dataCh),
		changeDollarToIdrConcurren(dataCh),
	)
	done := make(chan bool)

	writeFile(users , done)
	if <-done {
		log.Println("Donee")
	}
}

func synchronous() {
	user,_ := readFile("./data.json")
	result := changeDollarToIdr(user)
	log.Println(result[0], cap(result))
}

func main() {
	now := time.Now()
	asynchornous()
	log.Println("Done in", time.Since(now).Seconds())
}