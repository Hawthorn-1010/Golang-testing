package impl

import "fmt"

type ClassLesson interface {
	LearnC(string) (string, error)
	LearnJAVA(string) (string, error)
	LearnGO(string) (string, error)
}

var _ClassLesson = (*ClassLessonImpl)(nil)

type ClassLessonImpl struct {
}

func NewClassLessonImpl() *ClassLessonImpl {
	return &ClassLessonImpl{}
}

func (c *ClassLessonImpl) LearnC(s string) (string, error) {
	fmt.Println(s, "learn C at class.")
	return "learn C at class", nil
}

func (c *ClassLessonImpl) LearnJAVA(s string) (string, error) {
	fmt.Println(s, "learn JAVA at class.")
	return "learn JAVA at class", nil
}

func (c *ClassLessonImpl) LearnGO(s string) (string, error) {
	fmt.Println(s, "learn GO at class.")
	return "learn GO at class", nil
}
