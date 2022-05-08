package main

import (
	"fmt"
	"testing"
)

type OsIncorrectInputSet struct {
	err  error
	fPfn []string
}

type OsCorrectInputSet struct {
	fP   string
	fN   string
	fPfn []string
}

var OsErrorTestingSets = []OsIncorrectInputSet{
	{ArgsError, []string{"1", "1", "1", "1"}},
	{ArgsError, []string{"1"}},
	{PathError, []string{"", "", "1"}},
	{PathError, []string{"", " ", "1"}},
	{FuncNameError, []string{"", "123", ""}},
	{FuncNameError, []string{"", "123", " "}},
}

var OsCorrectTestingSets = []OsCorrectInputSet{
	{"makers.go", "foo", []string{"", "makers.go", "foo"}},
	{"new.go", "boo", []string{"", "new.go", "boo"}},
	{"main.go", "zoo", []string{"sd", "main.go", "zoo"}},
}

// 1. тест без аргументов в командной строке

func TestErrorGetDataFromStdIn(t *testing.T) {
	for _, set := range OsErrorTestingSets {
		_, _, err := GetDataFromStdIn(set.fPfn)
		if err != set.err {
			fmt.Printf("набор данных: %v\nожидалась ошибка: %s\nполучена ошика: %s\n", set.fPfn, set.err, err)
		}
	}
}

// 2. тест с валидными данными

func TestCorrectGetDataFromStdIn(t *testing.T) {
	for _, set := range OsCorrectTestingSets {
		fP, fN, err := GetDataFromStdIn(set.fPfn)
		if err != nil {
			fmt.Printf("набор данных: %v\nожидалась ошибка: %v\nполучена ошика: %s\n", set.fPfn, nil, err)
		}
		if fP != set.fP || fN != set.fN {
			fmt.Printf("тестовый набор: %v\nожидалось fP: %s, fN: %s\nполучено fP: %s, fN: %s\n",
				set.fPfn,
				set.fP,
				set.fN,
				fP,
				fN,
			)
		}
	}
}

type PathNamePair struct {
	AsyncFuncCnt int
	err          error
	FilePath     string
	FunctionName string
}

var IncorrecCounterSets = []PathNamePair{
	{0, ParseError, "noExistFilename.go", "noExistFunc"},
	{0, FuncNameError, "tfunction.go", "main"},
}

var CorrectCounterSets = []PathNamePair{
	{3, nil, "tfunction.go", "foo"},
	{0, nil, "tfunction.go", "zoo"},
	{1, nil, "tfunction.go", "boo"},
}

func TestIncorrectAsyncFuncCallsCounter(t *testing.T) {
	for _, set := range IncorrecCounterSets {
		AsyncFuncCnt, err := AsyncFuncCallsCounter(set.FilePath, set.FunctionName)

		if err != set.err || err == nil || AsyncFuncCnt != set.AsyncFuncCnt {
			fmt.Printf("набор данных: %s, %s\nожидалась ошибка: %s\nполучена ошика: %s\n",
				set.FilePath,
				set.FunctionName,
				set.err,
				err,
			)
		}
	}
}

func TestCorrectAsyncFuncCallsCounter(t *testing.T) {
	for _, set := range CorrectCounterSets {
		AsyncFuncCnt, err := AsyncFuncCallsCounter(set.FilePath, set.FunctionName)
		if err != nil {
			fmt.Printf("набор данных: %s, %s\nожидалась ошибка: %s\nполучена ошика: %s\n",
				set.FilePath,
				set.FunctionName,
				set.err,
				err,
			)
		}
		if AsyncFuncCnt != set.AsyncFuncCnt {
			fmt.Printf("набор данных: %s, %s\nожидался результат: %d\nполучен результат: %d\n",
				set.FilePath,
				set.FunctionName,
				set.AsyncFuncCnt,
				AsyncFuncCnt,
			)
		}

	}
}
