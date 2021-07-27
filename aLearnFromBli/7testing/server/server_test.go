package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// go test -v 所有的testcase
// go test -v -run Double 跑指定的testcase

// 单元测试
// 表组测试 for range

//以Testxxx开头
func TestDoubleHandler(t *testing.T) {
// 包含输入
//正常返回信息，返回码，错误返回信息
	testCases := []struct {
		name   string
		input  string
		result int
		status int
		err    string
	}{
		{name: "double of two", input: "2", result: 4, status: http.StatusOK, err: ""},
		{name: "double of nine", input: "9", result: 18, status: http.StatusOK, err: ""},
		{name: "double of nil", input: "", status: http.StatusBadRequest, err: "missing value"},
	}

	for _, testCase := range testCases {
		// 注册subtest集
		t.Run(testCase.name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, "localhost:4000/double?v="+testCase.input, nil)
			if err != nil {
				t.Fatalf("could not create a new request: %v, err: %v", req, err)
			}
//不需要真的起server， 而是把结果打桩
//ResponseRecorder is an implementation of http.ResponseWriter 接口
			rec := httptest.NewRecorder()
			// 调用handler，返回值
			doubleHandler(rec, req)
			res := rec.Result() //获取打桩的结果

			if res.StatusCode != testCase.status {
				t.Errorf("received status code %d, expect %d", res.StatusCode, testCase.status)
				return
			}

			respBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("cannot read all from the response body, err: %v", err)
			}
			defer res.Body.Close()
			trimedResult := strings.TrimSpace(string(respBytes))

			if res.StatusCode != http.StatusOK {
				// check the error message
				if trimedResult != testCase.err {
					t.Errorf("received error message %s, expect %s", trimedResult, testCase.err)
				}
				return
			}

			// compare the returned value
			doubleVal, err := strconv.Atoi(trimedResult)
			if err != nil {
				t.Errorf("cannot convert response body to int, err: %v", err)
				return
			}

			if doubleVal != testCase.result {
				t.Errorf("received result %d, expected %d", doubleVal, testCase.result)
			}
		})
	}
}