package main

import (
	"Golang-testing/projectstructure/impl"
)

func main() {
	impl.NewMyPlan(impl.NewClassLessonImpl()).Learn("hzy")
	impl.NewMyPlan(impl.NewRemoteLessonImpl()).Learn("hzy")
}
