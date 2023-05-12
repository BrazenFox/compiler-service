package service

type ProgramRunner interface{

}

type Service struct {
	ProgramRunner
}

func NewService() *Service {
	return &Service{}
}