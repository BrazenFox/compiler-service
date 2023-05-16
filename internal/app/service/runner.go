package service

import "github.com/BrazenFox/compiler-service/pkg/entity"

type RunProgramService struct {
}

func NewRunProgramService() *RunProgramService {
	return &RunProgramService{}
}

func (s *RunProgramService) RunProgram(program entity.Program) (string, error) {
	var result = "result_value:" + program.Language + program.Code

	return result, nil
}
