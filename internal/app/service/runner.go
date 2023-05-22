package service

import (
	"errors"
	"github.com/BrazenFox/compiler-service/pkg/entity"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

type RunProgramService struct {
}

type Runner struct {
	Path     string
	Filename string
	Content  string
	Commands []string
}

func NewRunProgramService() *RunProgramService {
	return &RunProgramService{}
}

func (s *RunProgramService) RunProgram(program entity.Program) (string, error) {
	//var result = "result_value:" + program.Language + program.Code

	runner := new(Runner)
	runner.Path = program.Language + "/" + program.UserId + "/"
	runner.Filename = program.Name + "." + program.Language
	runner.Content = program.Code

	if err := os.MkdirAll(runner.Path, os.ModePerm); err != nil {
		logrus.Error("filepath can't be created")
	}

	file, err := os.Create(runner.Path + runner.Filename)
	if err != nil {
		logrus.Error("Unable to open file", err)
		return "", err
	}

	logrus.Info("file created", file.Name(), file)

	_, err = file.WriteString(runner.Content)
	if err != nil {
		logrus.Error("Unable to write data", err)
		return "", err
	}
	file.Close()

	var output string

	switch lang := program.Language; lang {
	case "java":
		runner.Commands = append(runner.Commands, "javac", "java")
		_, err := exec.Command(runner.Commands[0], runner.Path+runner.Filename).Output()
		if err != nil {
			logrus.Error("Fail compile java")
			return "", err
		}

		out, err := exec.Command(runner.Commands[1], runner.Path+program.Name+".class").Output()
		if err != nil {
			logrus.Error("Fail compile java")
			return "", err
		}

		output = string(out[:])

	case "python":
		runner.Commands = append(runner.Commands, "python3")

		out, err := exec.Command(runner.Commands[0], runner.Path+program.Name+".py").Output()
		logrus.Info(runner.Commands[0], runner.Path+program.Name+".py")
		if err != nil {
			logrus.Error("Fail in python compile/runtime")
			return "", err
		}

		output = string(out[:])
	case "go":
		runner.Commands = append(runner.Commands, "go", "run")

		out, err := exec.Command(runner.Commands[0], runner.Commands[1], runner.Path+runner.Filename).Output()
		logrus.Info(runner.Commands[0], runner.Commands[1], runner.Path+runner.Filename)
		if err != nil {
			logrus.Error("Fail in go compile/runtime")
			return "", err
		}
		output = string(out[:])
	default:
		logrus.Error("this programm language doesn't supported")
		return "", errors.New("this programm language doesn't supported")
	}

	logrus.Info("Command Successfully Executed")

	return output, nil
}

//{
//    "language": "go",
//    "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello from external program!!!\")\n}",
//    "user_id": "user",
//    "name": "test"
//}
