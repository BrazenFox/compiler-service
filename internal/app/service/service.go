package service

import "github.com/BrazenFox/compiler-service/pkg/entity"

type ProgramRunner interface {
	RunProgram(program entity.Program) (string, error)
}

type Service struct {
	ProgramRunner
}

func NewService() *Service {
	return &Service{
		ProgramRunner: NewRunProgramService(),
	}

}
