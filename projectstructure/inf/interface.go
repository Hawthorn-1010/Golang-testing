package inf

type Lesson interface {
	LearnC(string) (string, error)
	LearnJAVA(string) (string, error)
	LearnGO(string) (string, error)
}
