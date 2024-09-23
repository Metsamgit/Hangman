package main

import (
    "fmt"
    "io/ioutil"
)

func play() {
	data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(data)) 
}