package usecase

type Usecase struct {
	*Parser
}

func NewUsecase() *Usecase {
	return &Usecase{}
}
