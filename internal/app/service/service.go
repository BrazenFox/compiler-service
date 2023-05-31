package service

import "github.com/BrazenFox/compiler-service/pkg/entity"

type ProgramHandler interface {
	HandleProgram(program entity.Program) (string, error)
}

type Service struct {
	ProgramHandler
}

func NewService() *Service {
	return &Service{
		ProgramHandler: NewProgramHandlerService(),
	}

}
