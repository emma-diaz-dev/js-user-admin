package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/emma-diaz-dev/js-user-admin/cli/script/cfg"
	"github.com/emma-diaz-dev/js-user-admin/cli/script/entities"
	"github.com/emma-diaz-dev/js-user-admin/cli/script/repositories"
)

var (
	inputFile = cfg.GetInput()
	conf      = cfg.GetConfig()
)

func main() {
	tests := []struct {
		name      string
		url       string
		inputUser map[string]io.Reader
		f         func(url string, values map[string]io.Reader) (interface{}, error)
		check     func(a interface{}) bool
		output    interface{}
	}{
		{
			name: "successful create new user case",
			url:  conf.OwnURL + "/user",
			inputUser: func() map[string]io.Reader {
				return map[string]io.Reader{
					"file":      mustOpen(inputFile.FileName),
					"name":      strings.NewReader("Luis"),
					"last_name": strings.NewReader("Diaz"),
					"address":   strings.NewReader("San Jose 1235, Cordoba, Argentina"),
				}
			}(),
			f: repositories.InsertFile,
			check: func(a interface{}) bool {
				v, _ := a.(entities.UserCreateResponse)
				status := v.Documents == 1
				fmt.Println(v, status)
				return status
			},
			output: entities.UserCreateResponse{Documents: 1},
		},
		{
			name:      "successful get users case",
			url:       conf.OwnURL + "/users",
			inputUser: nil,
			f:         repositories.GetUsers,
			check: func(a interface{}) bool {
				v, _ := a.(entities.UserResponse)
				for _, e := range v.Users {
					fmt.Println(e)
				}
				status := v.Documents > 0
				fmt.Println(v.Documents, status)
				return status
			},
			output: entities.UserCreateResponse{Documents: 1},
		},
	}

	for i, test := range tests {
		fmt.Printf("------- test: %s index: %d -------\n", test.name, i)
		result, err := test.f(test.url, test.inputUser)
		if err != nil {
			fmt.Println(err)
			continue
		}
		test.check(result)
		fmt.Printf("--------------- END index: %d ---------------\n", i)
	}
}

func mustOpen(f string) *os.File {
	fmt.Println("aaaaaa", f)
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
