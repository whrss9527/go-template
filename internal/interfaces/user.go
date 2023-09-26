package interfaces

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/gin-gonic/gin"

	"go-template/internal/biz"
)

type GinService struct {
	uc  *biz.TemplateUsecase
	log *log.Helper
}

func NewGinService(uc *biz.TemplateUsecase, logger log.Logger) *GinService {
	return &GinService{uc: uc, log: log.NewHelper(logger)}
}

func (u *GinService) SayHi(ctx *gin.Context) {

	// 模拟流式返回
	w := ctx.Writer
	header := w.Header()
	//在响应头添加分块传输的头字段Transfer-Encoding: chunked
	header.Set("Transfer-Encoding", "chunked")
	header.Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	//Flush()方法，好比服务端在往一个文件中写了数据，浏览器会看见此文件的内容在不断地增加。
	w.Write([]byte(`
            <html>
                    <body>
        `))
	w.(http.Flusher).Flush()

	for i := 0; i < 1000; i++ {
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

	ctx.JSON(200, gin.H{
		"msg": "hello world",
	})
}
