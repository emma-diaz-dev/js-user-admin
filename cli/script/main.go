package main

import (
	"encoding/json"
	"errors"
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
	var (
		values = map[string]io.Reader{
			"file":      mustOpen(inputFile.FileName),
			"name":      strings.NewReader("Luis"),
			"last_name": strings.NewReader("Diaz"),
			"address":   strings.NewReader("San Jose 1235, Cordoba, Argentina"),
		}
		url = conf.OwnURL + "/user"
	)

	result, err := repositories.InsertUser(url, values)
	if err != nil {
		fmt.Println(err)
		return
	}
	newUser, _ := result.(entities.UserCreateResponse)
	if newUser.Documents <= 0 {
		fmt.Println(errors.New("error in external resource"))
		return
	}

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
			f: repositories.InsertUser,
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
		{
			name:      "successful get user by id case",
			url:       conf.OwnURL + "/user/" + newUser.ID,
			inputUser: nil,
			f:         repositories.GetUser,
			check: func(a interface{}) bool {
				v, _ := a.(entities.User)
				status := v.Name != ""
				b, _ := json.Marshal(&v)
				fmt.Println(string(b), status)
				return status
			},
			output: entities.UserCreateResponse{Documents: 1},
		},
		{
			name: "successful update user name by id case",
			url:  conf.OwnURL + "/user/" + newUser.ID,
			inputUser: func() map[string]io.Reader {
				return map[string]io.Reader{
					"name":    strings.NewReader("Emmanuel123"),
					"address": strings.NewReader("San Jose 1235, Cordoba, Argentina"),
				}
			}(),
			f: repositories.UpdateUser,
			check: func(a interface{}) bool {
				v, _ := a.(entities.UserResponse)
				status := v.Documents == 1
				b, _ := json.Marshal(&v)
				fmt.Println(string(b), status)
				return status
			},
			output: entities.UserCreateResponse{Documents: 1},
		},
	}

	for i, test := range tests {
		fmt.Printf("----- START Nro.: %d  %s -----\n", i+1, test.name)
		result, err := test.f(test.url, test.inputUser)
		if err != nil {
			fmt.Println(err)
			continue
		}
		test.check(result)
		fmt.Printf("--------------- END index: %d ---------------\n", i+1)
		fmt.Println()
		fmt.Println()
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
