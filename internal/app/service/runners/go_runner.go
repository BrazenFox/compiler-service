package runners

import (
	"github.com/BrazenFox/compiler-service/pkg/entity"
	"github.com/sirupsen/logrus"
	"os/exec"
)

type GoRunnerService struct {
	Path     string
	Filename string
	Code     string
}

func NewGoRunnerService(program entity.Program) *GoRunnerService {
	return &GoRunnerService{
		Path:     "program/" + program.Language + "/" + program.UserId,
		Filename: program.Name,
		Code:     program.Code,
	}
}

func (s *GoRunnerService) RunProgram() (string, error) {
	out, err := exec.Command("go", "run", s.GetPath()+s.GetFileName()).Output()
	logrus.Info("go", "run", s.Path+s.Filename)
	if err != nil {
		logrus.Error("Fail in go compile/runtime")
		return "", err
	}
	return string(out[:]), nil
}

func (s *GoRunnerService) GetPath() string {
	return s.Path + "/"
}

func (s *GoRunnerService) GetFileName() string {
	return s.Filename + ".go"
}

func (s *GoRunnerService) GetCode() string {
	return s.Code
}
