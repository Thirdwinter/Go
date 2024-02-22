package middleware

import (
	"fmt"
	"strconv"

	log "github.com/ThirdWinter/Go/mylog"
	"github.com/gin-gonic/gin"
)

func write(c *gin.Context) string {
    // 获取请求头信息
    headers := c.Request.Header
    headersStr := ""
    for key, values := range headers {
        headersStr += fmt.Sprintf("[%s: %s] ", key, values)
    }
    
    // 获取请求体信息
    requestBody, _ := c.GetRawData()
    
    return fmt.Sprintf("[METHOD: %s] [URL: %s] [HOST: %s] [HEADERS: %s] [BODY: %s]", c.Request.Method, c.Request.URL.Path, c.Request.Host, headersStr, string(requestBody))
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		// 在这里编写你的日志逻辑
		//log.Debug(write(c))
		// 继续处理请求
		c.Next()
		// 获取响应状态码
        status := c.Writer.Status()

        // 根据不同的状态码选择不同的日志级别
        switch {
        case status >= 500:
            defer func() {
				if err := recover(); err != nil {
					log.Error("%s: Panic recovered: %s", write(c),err)
					// 可以在这里记录错误信息到日志中
					c.AbortWithStatus(500) // 返回500错误状态码
				}
			}()
        case status >= 400:
            log.Warning("HTTP Status %s:%s ", strconv.Itoa(status), write(c))
        case status >= 300:
            log.Info("HTTP Status %s:%s ", strconv.Itoa(status), write(c))
        default:
            log.Trace("HTTP Status %s:%s ", strconv.Itoa(status), write(c))
        }
}
}