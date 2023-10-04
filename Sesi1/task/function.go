package main

import "fmt"

func cetak(arg interface{}) {
	fmt.Println(arg)
}

func getMergeString(car map[string]string)string {
	return "Mobil " + car["name"] + " berwarna " + car["color"]
}

func main() {
    var car = map[string]string{
        "name":  "BMW",
        "color": "Black",
    }

	message := getMergeString(car)
	cetak(message)
}