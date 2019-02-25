package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Emmett struct {
	age   int
	place string
}

type Mike struct {
	dateOfBirth string
	job         string
}

type Person interface {
	Age() int
	Profession() string
}

func (e Emmett) Age() int {
	return e.age
}

func (e Emmett) Profession() string {
	if e.place == "clifton park" {
		return "Kitware"
	}
	return ""
}

func (m Mike) Age() int {
	mdy := strings.Split(m.dateOfBirth, "/")
	// m := int(mdy[0])
	// d := int(mdy[1])
	y, _ := strconv.Atoi(mdy[2])
	time := 2019
	return time - y
}

func (m Mike) Profession() string {
	return m.job
}

func GetResumeSpecs(p Person) {
	fmt.Println(p.Age())
	fmt.Println(p.Profession())
}

func main() {
	emmett := Emmett{
		age:   24,
		place: "clifton park",
	}
	mike := Mike{
		dateOfBirth: "1/1/1988",
		job:         "rnd",
	}
	GetResumeSpecs(emmett)
	GetResumeSpecs(mike)
}
