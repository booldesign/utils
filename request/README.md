# HttpRequest 

### Usage

#### 1.异步http request时，需要初始化异步请求池，写在main里
```go
// 需要异步http时，poolSize值10~100
request.InitAsyncRequestPool(poolSize) 
```
