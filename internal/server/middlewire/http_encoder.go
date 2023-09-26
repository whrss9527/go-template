package middlewire

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("ResponseError: %d", e.Code)
}

func newResponseError(code string, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

// fmtError try to convert an error to *ResponseError.
func fmtError(err error) *ResponseError {
	if err == nil {
		return nil
	}
	if se := new(errors.Error); errors.As(err, &se) {
		return newResponseError(se.Reason, se.Message)
	}
	return &ResponseError{Code: "ServerErr", Message: err.Error()}
}

// ErrorEncoder 自定义返回error
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	if se := new(errors.Error); !errors.As(err, &se) {
		w.WriteHeader(500)
	}
	se := fmtError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(body)
}

// ResponseEncoder 自定义返回数据结构
func ResponseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	header := w.Header()
	//在响应头添加分块传输的头字段Transfer-Encoding: chunked
	header.Set("Transfer-Encoding", "chunked")
	header.Set("Content-Type", "text/html")
	w.WriteHeader(200)

	//Flush()方法，好比服务端在往一个文件中写了数据，浏览器会看见此文件的内容在不断地增加。
	w.Write([]byte(`
            <html>
                    <body>
        `))
	w.(http.Flusher).Flush()

	for i := 0; i < 10; i++ {
		w.Write([]byte(fmt.Sprintf(`
                <h1>%d</h1>
            `, i)))
		w.(http.Flusher).Flush()
		time.Sleep(time.Duration(1) * time.Second)
	}

	w.Write([]byte(`
                    </body>
            </html>
        `))
	w.(http.Flusher).Flush()
	//http.Flusher.Flush()

	return nil
}
