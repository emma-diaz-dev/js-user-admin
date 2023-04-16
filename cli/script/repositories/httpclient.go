package repositories

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/emma-diaz-dev/js-user-admin/cli/script/entities"
)

func InsertFile(url string, values map[string]io.Reader) (interface{}, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var (
			fw  io.Writer
			err error
		)
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else if fw, err = w.CreateFormField(key); err != nil {
			return nil, err
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}

	}
	w.Close()

	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := entities.UserCreateResponse{}
	_ = json.Unmarshal(body, &result)
	return result, nil
}

func UpdateFile(url string, values map[string]io.Reader) (result interface{}, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else if fw, err = w.CreateFormField(key); err != nil {
			return
		}
		if _, err = io.Copy(fw, r); err != nil {
			return
		}

	}
	w.Close()

	req, err := http.NewRequest(http.MethodPut, url, &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	_ = json.Unmarshal(body, &result)
	return result, nil
}

func GetUser(url string, values map[string]io.Reader) (interface{}, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := entities.User{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetUsers(url string, values map[string]io.Reader) (interface{}, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := entities.UserResponse{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
