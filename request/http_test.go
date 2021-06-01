package request

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

//  测试get请求（同步）
func TestRequestGet(t *testing.T) {
	url := "http://localhost:8080/frontend/v1/health"
	queries := map[string]string{
		"t":"1622394061000",
		"p":"iOS",
		"v":"1.0.0",
	}
	errRes, err := NewHttpRequest(url).
		SetMethod(http.MethodGet).
		AddHeader("Authorization", "Basic RlJfMDAwMDE6NTM0ZjcyNDU2ZjRhMTM4NWE3ZTQ4NDZiYzdkN2Vj").
		SetQueries(queries).
		Response()
	fmt.Println("err: ", err)
	fmt.Printf("resp: %+v\n", string(errRes))
}

// 测试get请求（同步）
func TestRequestGet2(t *testing.T) {
	url := "http://localhost:8080/frontend/v1/health"
	queries := map[string]string{
		"t":"1622394061000",
		"p":"iOS",
		"v":"1.0.0",
	}
	// 直接获取data的数据指针
	type testBack struct {
		RunMode string `json:"runMode"`
		Host    string `json:"host"`
		Group   string `json:"group"`
	}
	data := &testBack{}
	err := NewHttpRequest(url).SetMethod(http.MethodGet).SetQueries(queries).Unmarshal(data)
	fmt.Println("err: ", err)
	fmt.Printf("resp: %+v\n", data)
}

// 测试post请求（同步）
func TestRequestPOST(t *testing.T)  {
	url := "http://127.0.0.1:8080/frontend/v1/user.login"
	queries := map[string]string{
		"t":"1622394061000",
		"p":"iOS",
		"v":"1.0.0",
	}
	formData := map[string]interface{}{
		"username":  "123uwrr",
		"password": "1234567a",
	}
	result, err := NewHttpRequest(url).
		SetMethod(http.MethodPost).
		SetQueries(queries).
		SetContentType("form").
		SetFormBody(formData).
		Response()
	fmt.Println("err: ", err)
	fmt.Println("resp: ", string(result))
}

// 测试post请求（异步）
func TestRequestPOST3(t *testing.T)  {
	InitAsyncPool(10)
	url := "http://127.0.0.1:8080/frontend/v1/user.register"
	formData := map[string]interface{}{
		"username":"wei3351",
		"password":"1234567a",
	}
	// 只获取success与否 - 返回true/false
	result, err := NewHttpRequest(url).
		SetMethod(http.MethodPost).
		SetContentType("form").
		SetFormBody(formData).
		SetAsync(true).
		Result()
	fmt.Println("err :", err)
	fmt.Println("result = ", result)

	// 等待协程异步执行
	time.Sleep(3 * time.Second)
}


// 测试本地文件上传（同步）
func TestRequestFileUpload(t *testing.T){

}

// 测试文件二进制流上传（异步）
func TestRequestSteamUpload(t *testing.T) {

}
