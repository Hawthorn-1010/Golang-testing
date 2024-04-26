package impl

import (
	"Golang-testing/projectstructure/inf"
	"fmt"
)

type MyLessonProxy interface {
	LearnGO(string) (string, error)
}

type MyPlan struct {
	learnEvent MyLessonProxy
}

func NewMyPlan(learnEvent inf.Lesson) *MyPlan {
	return &MyPlan{
		learnEvent: learnEvent,
	}
}

func (p *MyPlan) Learn(s string) {
	c, err := p.learnEvent.LearnGO(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
