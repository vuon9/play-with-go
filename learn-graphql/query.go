package main

import (
	"fmt"
	"math/rand"

	"github.com/sirupsen/logrus"
)

type query struct{}

func (*query) Name() (string, error) {
	return "Hello, world!", nil
}

func (*query) School() *schoolQuery {
	return &schoolQuery{}
}

type schoolQuery struct{}

func (*schoolQuery) ID() string {
	return "1"
}

func (*schoolQuery) Name() string {
	logrus.Info("I want to get school name")
	return "1st Elementary School"
}

func (*schoolQuery) Address() string {
	logrus.Info("I want to get school address")
	return fmt.Sprintf("%d Bridge, Somewhere, Daknong", rand.Intn(100))
}

func (*schoolQuery) Students() []*studentQuery {
	return []*studentQuery{
		{id: 1},
		{id: 2},
	}
}

type Student struct {
	ID           int32
	Name         string
	TotalCourses int32
}

var studentsDB = map[int32]Student{
	1: {ID: 1, Name: "Anna", TotalCourses: 5},
	2: {ID: 2, Name: "Brian", TotalCourses: 10},
}

type studentQuery struct {
	id int32
}

func (q *studentQuery) ID() int32 {
	return q.id
}

func (q *studentQuery) Name() string {
	return studentsDB[q.id].Name
}

func (q *studentQuery) TotalCourses() int32 {
	return studentsDB[q.id].TotalCourses
}
