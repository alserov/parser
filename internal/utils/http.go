package utils

import (
	"io"
	"net/http"
)

func MakeRequest(method, url string, body io.Reader) (*http.Response, error) {
	var (
		err error
		req *http.Request
		res *http.Response
	)

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, NewError(ErrBadRequest, err.Error())
	}

	cl := &http.Client{}

	res, err = cl.Do(req)
	if err != nil {
		return nil, NewError(ErrBadRequest, err.Error())
	}

	return res, nil
}
