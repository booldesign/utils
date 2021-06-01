package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/booldesign/utils/json"
	"github.com/opentracing/opentracing-go"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/5/31 11:46
 * @Desc:
 */

type HttpRequest struct {
	span         *opentracing.Span
	url          string
	method       string
	async        bool
	headers      map[string]string
	cookies      map[string]string
	queries      map[string]string
	contentType  string // form(x-www-form-urlencoded), form-data, json(application/json)
	formBody     map[string]interface{}
	rawBody      []byte
	formDataBody map[string]interface{}
	timeout      time.Duration // 同步请求超时, 默认10秒
}

var (
	httpAsyncPool   chan *HttpRequest // 异步请求池
	defaultPoolSize = 10
	maxPoolSize     = 50
)

// InitAsyncPool 初始化异步请求池
func InitAsyncPool(poolSize int) {
	if poolSize < defaultPoolSize {
		poolSize = defaultPoolSize
	}
	if poolSize > maxPoolSize {
		poolSize = maxPoolSize
	}

	httpAsyncPool = make(chan *HttpRequest, poolSize)

	// 监听消费异步请求
	go handleAsyncRequest()
}

// NewHttpRequest 创建Http请求
func NewHttpRequest(url string) *HttpRequest {
	req := &HttpRequest{
		url:     url,
		timeout: 10 * time.Second,
		headers: make(map[string]string),
		cookies: make(map[string]string),
		queries: make(map[string]string),
	}
	return req
}

// SetTracing 设置追踪
func (req *HttpRequest) SetTracing(span *opentracing.Span) *HttpRequest {
	req.span = span
	return req
}

// SetMethod 设置Http请求方法
func (req *HttpRequest) SetMethod(method string) *HttpRequest {
	req.method = strings.ToUpper(method)
	return req
}

// SetAsync 设置异步
func (req *HttpRequest) SetAsync(on bool) *HttpRequest {
	req.async = on
	return req
}

// SetHeaders 设置Headers
func (req *HttpRequest) SetHeaders(headers map[string]string) *HttpRequest {
	req.headers = headers
	return req
}

// AddHeader 增加一个Header
func (req *HttpRequest) AddHeader(header, val string) *HttpRequest {
	req.headers[header] = val
	return req
}

// SetCookies 设置cookies
func (req *HttpRequest) SetCookies(cookies map[string]string) *HttpRequest {
	req.cookies = cookies
	return req
}

// SetQueries 设置url携带的数据
func (req *HttpRequest) SetQueries(queries map[string]string) *HttpRequest {
	req.queries = queries
	return req
}

// SetContentType 设置ContentType
func (req *HttpRequest) SetContentType(typeStr string) *HttpRequest {
	req.contentType = strings.ToLower(typeStr)
	return req
}

// SetFormBody 设置
func (req *HttpRequest) SetFormBody(body map[string]interface{}) *HttpRequest {
	req.formBody = body
	return req
}

// SetRawBody raw参数
func (req *HttpRequest) SetRawBody(body []byte) *HttpRequest {
	req.rawBody = body
	return req
}

// SetFormDataBody 设置文件上传 文件路径名map/文件二进制内容map
func (req *HttpRequest) SetFormDataBody(body map[string]interface{}) *HttpRequest {
	req.formDataBody = body
	return req
}

// SetTimeout 设置超时
func (req *HttpRequest) SetTimeout(timeout time.Duration) *HttpRequest {
	req.timeout = timeout
	return req
}

// doRequest 执行请求
func (req *HttpRequest) doRequest() ([]byte, error) {
	if req.url == "" {
		return nil, errors.New("请求地址不能为空")
	}
	if req.method == "" {
		return nil, errors.New("请求方式不能为空")
	}
	var body io.Reader
	switch req.method {
	case http.MethodGet, http.MethodDelete:
		body = nil
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions:
		if req.contentType == "json" || req.contentType == "application/json" {
			req.AddHeader("Content-Type", "application/json")
			if string(req.rawBody) != "" {
				body = bytes.NewReader(req.rawBody)
			} else if req.formBody != nil {
				jsonData := json.JsonEncode(req.formBody)
				if jsonData != nil {
					return nil, errors.New("formBody 参数错误")
				}
				body = bytes.NewReader(jsonData)
			}
		} else if req.contentType == "form" || req.contentType == "application/x-www-form-urlencoded" {
			req.AddHeader("Content-Type", "application/x-www-form-urlencoded")
			if string(req.rawBody) != "" {
				body = bytes.NewReader(req.rawBody)
			} else if req.formBody != nil {
				postParams := url.Values{}
				for k, v := range req.formBody {
					switch v.(type) {
					case []string:
						for _, vv := range v.([]string) {
							postParams.Set(k, vv)
						}
					case []int:
						for _, vv := range v.([]int) {
							postParams.Set(k, convertStr(vv))
						}
					case []float64:
						for _, vv := range v.([]float64) {
							postParams.Set(k, convertStr(vv))
						}
					case []bool:
						for _, vv := range v.([]bool) {
							postParams.Set(k, convertStr(vv))
						}
					case string:
						postParams.Set(k, v.(string))
					default:
						postParams.Set(k, fmt.Sprintf("%v", v))

					}
				}
				body = ioutil.NopCloser(strings.NewReader(postParams.Encode()))
			} else if req.contentType == "form-data" {
				// TODO
			}
		}
	default:
		return nil, errors.New("无效的请求方式")
	}

	request, err := http.NewRequest(req.method, req.url, body)
	if err != nil {
		return nil, err
	}
	//设置get请求url参数
	q := request.URL.Query()
	for k, v := range req.queries {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
	//设置请求头
	for k, v := range req.headers {
		request.Header.Set(k, v)
	}
	//设置cookie
	for k, v := range req.cookies {
		request.AddCookie(
			&http.Cookie{Name: k, Value: v},
		)
	}
	//Inject Header
	if req.span != nil {
		err := opentracing.GlobalTracer().Inject(
			(*req.span).Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header))
		if err != nil {
			log.Printf("OpenTracing Inject Error:%s", err.Error())
		}
	}

	client := &http.Client{
		Timeout: req.timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//当发生错误时记录详情的请求信息
		if req.span != nil {
			(*req.span).SetTag("curl.error", err)
			(*req.span).SetTag("curl.url", request.URL.RawQuery)
			(*req.span).SetTag("curl.params", req.formBody)
		}
		return []byte(""), err
	}
	if resp.StatusCode != http.StatusOK {
		//当发生错误时记录详情的请求信息
		if req.span != nil {
			(*req.span).SetTag("curl.error", err)
			(*req.span).SetTag("curl.url", request.URL.RawQuery)
			(*req.span).SetTag("curl.params", req.formBody)
			(*req.span).SetTag("status.code", resp.Status)
		}
	}

	return content, nil
}

func convertStr(v interface{}) string {
	var str string
	switch v.(type) {
	case int:
		str = strconv.Itoa(v.(int))
	case int8:
		str = strconv.Itoa(int(v.(int8)))
	case int64:
		str = strconv.Itoa(int(v.(int64)))
	case float64:
		str = strconv.FormatFloat(v.(float64), 'g', -1, 64)
	case bool:
		str = strconv.FormatBool(v.(bool))
	}
	return str
}

// Response 原始http请求（返回byte）
func (req *HttpRequest) Response() ([]byte, error) {
	//发布异步请求任务
	if req.async {
		//未初始化请求池
		if httpAsyncPool == nil {
			return nil, errors.New("async request pool not init")
		}
		httpAsyncPool <- req
		response := &HttpResponse{
			Success: true,
		}
		return json.JsonEncode(response), nil
	}
	//处理同步请求
	return req.doRequest()
}

// Result http请求 返回封装
func (req *HttpRequest) Result() (*HttpResponse, error) {
	rsp, err := req.Response()
	if err != nil {
		return nil, err
	}
	response := &HttpResponse{}
	err = json.JsonDecode(rsp, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Unmarshal HTTP请求（填充data结构体信息）
func (req *HttpRequest) Unmarshal(data interface{}) error {
	rsp, err := req.Response()
	if err != nil {
		return err
	}
	response := &HttpResponse{}
	response.Data = data
	err = json.JsonDecode(rsp, response)
	return err
}

// 消费异步请求任务 TODO：消息队列
func handleAsyncRequest() {
	for {
		select {
		case req := <-httpAsyncPool:
			res, err := req.doRequest()
			if err != nil {
				log.Printf("handleAsyncHttpRequest Error:%s, Req:%#v\n", err.Error(), req)
			} else {
				log.Printf("handleAsyncHttpRequest Req:%#v, res:%s\n", req, string(res))
			}
		}
	}
}
