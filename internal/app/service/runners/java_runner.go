package runners

import (
	"github.com/BrazenFox/compiler-service/pkg/entity"
	"github.com/sirupsen/logrus"
	"os/exec"
)

type JavaRunnerService struct {
	Path     string
	Filename string
	Code     string
}

func NewJavaRunnerService(program entity.Program) *JavaRunnerService {
	return &JavaRunnerService{
		Path:     "program/" + program.Language + "/" + program.UserId,
		Filename: program.Name,
		Code:     program.Code,
	}
}

func (s *JavaRunnerService) RunProgram() (string, error) {
	_, err := exec.Command("javac", s.GetPath()+s.GetFileName()).Output()
	if err != nil {
		logrus.Error("Fail compile java")
		return "", err
	}

	out, err := exec.Command("java", "-cp", s.Path, s.Filename).Output()
	if err != nil {
		logrus.Error("Fail run java")
		return "", err
	}

	return string(out[:]), nil
}

func (s *JavaRunnerService) GetPath() string {
	return s.Path + "/"
}

func (s *JavaRunnerService) GetFileName() string {
	return s.Filename + ".java"
}

func (s *JavaRunnerService) GetCode() string {
	return s.Code
}
