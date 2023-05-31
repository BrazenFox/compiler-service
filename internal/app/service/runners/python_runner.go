package runners

import (
	"github.com/BrazenFox/compiler-service/pkg/entity"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
)

type PythonRunnerService struct {
	Path     string
	Filename string
	Code     string
}

func NewPythonRunnerService(program entity.Program) *PythonRunnerService {
	return &PythonRunnerService{
		Path:     "program/" + program.Language + "/" + program.UserId,
		Filename: program.Name,
		Code:     program.Code,
	}
}

func (s *PythonRunnerService) RunProgram() (string, error) {
	var command = "python"
	if runtime.GOOS != "windows" {
		command = command + "3"
	}

	out, err := exec.Command(command, s.GetPath()+s.GetFileName()).Output()
	logrus.Info(command, s.Path+s.Filename+".py")
	if err != nil {
		logrus.Error("Fail in python compile/runtime")
		return "", err
	}

	return string(out[:]), nil
}

func (s *PythonRunnerService) GetPath() string {
	return s.Path + "/"
}

func (s *PythonRunnerService) GetFileName() string {
	return s.Filename + ".py"
}

func (s *PythonRunnerService) GetCode() string {
	return s.Code
}
