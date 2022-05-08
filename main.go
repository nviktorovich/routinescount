package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

var (
	ArgsError     = errors.New("ошибка. Количество аргументов командной строки не должно превышать 3\n")
	ArgsGetError  = errors.New("ошибка. Не удалось получить аргументы командной строки\n")
	PathError     = errors.New("ошибка. Некоректный путь к файлу/имя файла\n")
	FuncNameError = errors.New("ошибка. Недопустимое имя функции\n")
	ParseError    = errors.New("ошибка. Не удалось произвести парсинг файла по указанному пути\n")
	AnaliseError  = errors.New("ошибка. Не удалось проанализировать переданную функцию\n")
)

func main() {

	//	1 получить аргументы os.Stdin
	userData := os.Args
	filePath, fnName, err := GetDataFromStdIn(userData)
	if err != nil {
		log.Fatal(ArgsGetError, err)
	}

	// 2 вызвать функцию для подсчета количества вызовов асинхронных функций и ошибку, передать, полученные в 1 параметры
	asyncCnt, err := AsyncFuncCallsCounter(filePath, fnName)
	if err != nil {
		log.Fatal(AnaliseError, err)
	}
	fmt.Printf("\nУказанный файл: \t\t`%s`\n"+
		"Указанное имя функции: \t\t`%s`\n"+
		"Вызвано асинхронных функций: \t%d\n\n",
		filePath,
		fnName,
		asyncCnt,
	)

}

// GetDataFromStdIn для возврата пути файла, имени функции, ошибки
func GetDataFromStdIn(args []string) (string, string, error) {

	if len(args) != 3 {
		err := ArgsError
		return "", "", err
	}

	// file path путь до файла
	fP := args[1]
	if fP == "" || fP == " " {
		return "", "", PathError
	}

	// function name название функции
	fN := args[2]
	if fN == "" || fN == " " {
		return "", "", FuncNameError
	}

	return fP, fN, nil
}

// AsyncFuncCallsCounter для возврата количества вызовов асинхронных функций.
// Возвращает количество вызрвов асинхронных функций (горутин)
func AsyncFuncCallsCounter(path, fnName string) (int, error) {
	var GoFuncCounter []*ast.GoStmt
	set := token.NewFileSet()
	if file, err := parser.ParseFile(set, path, nil, 0); err == nil {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if ok {
				if fn.Name.Name == fnName {
					ast.Inspect(decl, func(node ast.Node) bool {
						fun, yes := node.(*ast.GoStmt)
						if yes {
							GoFuncCounter = append(GoFuncCounter, fun)
						}
						return true
					})
					return len(GoFuncCounter), nil
				}
			}
		}
		return 0, FuncNameError
	}
	return 0, ParseError

}
