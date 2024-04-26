package impl

import "fmt"

type RemoteLesson interface {
	LearnC(string) (string, error)
	LearnJAVA(string) (string, error)
	LearnGO(string) (string, error)
}

var _RemoteLesson = (*RemoteLessonImpl)(nil)

type RemoteLessonImpl struct {
}

func NewRemoteLessonImpl() *RemoteLessonImpl {
	return &RemoteLessonImpl{}
}

func (c *RemoteLessonImpl) LearnC(s string) (string, error) {
	fmt.Println(s, "learn C remotely.")
	return "learn C remotely", nil
}

func (c *RemoteLessonImpl) LearnJAVA(s string) (string, error) {
	fmt.Println(s, "learn JAVA remotely.")
	return "learn JAVA remotely", nil
}

func (c *RemoteLessonImpl) LearnGO(s string) (string, error) {
	fmt.Println(s, "learn GO remotely.")
	return "learn GO remotely", nil
}
