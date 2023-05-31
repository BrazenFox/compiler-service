package service

import (
	"errors"
	"github.com/BrazenFox/compiler-service/internal/app/service/runners"
	"github.com/BrazenFox/compiler-service/pkg/entity"
	"github.com/sirupsen/logrus"
	"os"
)

type ProgramRunner interface {
	RunProgram() (string, error)
	GetPath() string
	GetFileName() string
	GetCode() string
}

type ProgramHandlerService struct {
}

func NewProgramHandlerService() *ProgramHandlerService {
	return &ProgramHandlerService{}
}

func (s *ProgramHandlerService) HandleProgram(program entity.Program) (string, error) {
	var runner ProgramRunner
	switch lang := program.Language; lang {
	case "java":
		runner = runners.NewJavaRunnerService(program)
	case "python":
		runner = runners.NewPythonRunnerService(program)
	case "go":
		runner = runners.NewGoRunnerService(program)
	default:
		logrus.Error("this programm language doesn't supported")
		return "", errors.New("this programm language doesn't supported")
	}
	CreateFolders(runner.GetPath(), runner.GetFileName(), runner.GetCode())

	output, err := runner.RunProgram()

	if err != nil {
		logrus.Error("Unable to write data", err)
		return "", err
	}

	logrus.Info("Command Successfully Executed")

	return output, nil
}

func CreateFolders(path string, filename string, code string) (string, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		logrus.Error("filepath can't be created")
	}
	file, err := os.Create(path + "/" + filename)
	if err != nil {
		logrus.Error("Unable to open file", err)
		return "", err
	}

	logrus.Info("file created", file.Name(), file)

	_, err = file.WriteString(code)
	if err != nil {
		logrus.Error("Unable to write data", err)
		return "", err
	}
	file.Close()

	return "", errors.New("")
}

func DeleteFolders() (str string) {
	return str
}

/*{
    "language": "go",
    "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello from external program!!!\")\n}",
    "user_id": "userGo",
    "name": "test"
}*/

/*{
    "language": "python",
    "code": "print(\"hello\")",
    "user_id": "userPython",
    "name": "test"
}*/

/*{
    "language": "java",
    "code": "public class Main {\n    public static void main(String[] args) {\n        System.out.println(\"Hello\");\n    }\n}",
    "user_id": "userJava",
    "name": "Main"
}*/
