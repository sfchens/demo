package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR        = 7
	SUCCESS      = 0
	Unauthorized = 401 // Unauthorized

	OkSuccess = 200

	InvalidParameter = 40001 // invalid argument
	MissingParameter = 40002 // 缺失参数
	BadRequest       = 40003 // 请求不能处理等
	BadMethod        = 40005 // 请求方法错误
	FailSection      = 40006 // 部分失败

	Forbidden = 40403 // 拒绝访问
	NotFound  = 40404 // 没找到记录等

	RecordExpired    = 20001 // 记录过期
	Duplicate        = 20002 // 重复操作
	BadParameter     = 20003 // 参数值错误，例如：参数过大等
	BalanceNotEnough = 20004 // 余额不足
	RecordDisabled   = 20005 // 不可用

	InternalServerError      = 50000 // 内部错误
	FetchResultError         = 50001 // 数据库查询记录错误
	FetchResultValidateError = 50002 // 数据库查询记录错误
	InvalidRecord            = 50101 // 有问题的数据
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithCode(code int, msg string, c *gin.Context) {
	Result(code, nil, msg, c)
}

func UnauthorizedWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Unauthorized, data, message, c)
}

func ErrorPage(status int, msg string, c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", gin.H{
		"status":  status,
		"message": msg,
	})
}

func OkSuccessWithData(data interface{}, c *gin.Context) {
	Result(OkSuccess, data, "操作成功", c)
}

func Redirect(location string, c *gin.Context) {
	c.Redirect(http.StatusFound, location)
}

func FailAsUnauthorized(data interface{}, msg string, c *gin.Context) {
	Result(Unauthorized, data, msg, c)
}

func FailWithNotFound(data interface{}, msg string, c *gin.Context) {
	Result(NotFound, data, msg, c)
}

func FailAorbiddensF(data interface{}, msg string, c *gin.Context) {
	Result(Forbidden, data, msg, c)
}

func FailInvalidParameter(data interface{}, msg string, c *gin.Context) {
	Result(InvalidParameter, data, msg, c)
}

func FailWithMissingParameter(data interface{}, msg string, c *gin.Context) {
	Result(MissingParameter, data, msg, c)
}

func FailWithInternalServerError(data interface{}, msg string, c *gin.Context) {
	Result(InternalServerError, data, msg, c)
}

func FailWithRecordExpired(data interface{}, msg string, c *gin.Context) {
	Result(RecordExpired, data, msg, c)
}

func FailWithDuplicate(data interface{}, msg string, c *gin.Context) {
	Result(Duplicate, data, msg, c)
}

func FailWithFetchResultError(data interface{}, msg string, c *gin.Context) {
	Result(FetchResultError, data, msg, c)
}

func FailWithInvalidRecordError(data interface{}, msg string, c *gin.Context) {
	Result(InvalidRecord, data, msg, c)
}

func FailWithBadParameter(data interface{}, msg string, c *gin.Context) {
	Result(BadParameter, data, msg, c)
}

func FailWithBalanceNotEnough(data interface{}, msg string, c *gin.Context) {
	Result(BalanceNotEnough, data, msg, c)
}

func FailSectionWithData(data interface{}, msg string, c *gin.Context) {
	Result(FailSection, data, msg, c)
}

func BadRequestWithData(data interface{}, msg string, c *gin.Context) {
	Result(BadRequest, data, msg, c)
}

func FailWithDataByCode(code int, msg string, data interface{}, c *gin.Context) {
	switch code {
	case Unauthorized:
		UnauthorizedWithDetailed(data, msg, c)
	case InvalidParameter:
		FailInvalidParameter(data, msg, c)
	case MissingParameter:
		FailWithMissingParameter(data, msg, c)
	case FailSection:
		FailSectionWithData(data, msg, c)
	case Forbidden:
		FailAorbiddensF(data, msg, c)
	case NotFound:
		FailWithNotFound(data, msg, c)
	case RecordExpired:
		FailWithRecordExpired(data, msg, c)
	case Duplicate:
		FailWithDuplicate(data, msg, c)
	case BadParameter:
		FailWithBadParameter(data, msg, c)
	case BalanceNotEnough:
		FailWithBalanceNotEnough(data, msg, c)
	case InternalServerError:
		FailWithInternalServerError(data, msg, c)
	case FetchResultError:
		FailWithFetchResultError(data, msg, c)
	case InvalidRecord:
		FailWithInvalidRecordError(data, msg, c)
	case BadRequest:
		BadRequestWithData(data, msg, c)
	default:
		Result(code, data, msg, c)
	}
}

func FailByCode(code int, msg string, c *gin.Context) {
	FailWithDataByCode(code, msg, "", c)
}
