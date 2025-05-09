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
	SUCCESS      = 200
	Unauthorized = 401 // Unauthorized

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

func Result(ctx *gin.Context, code int, data interface{}, msg string) {
	// 开始时间
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx *gin.Context) {
	Result(ctx, SUCCESS, map[string]interface{}{}, "操作成功")
}

func OkWithMessage(ctx *gin.Context, message string) {
	Result(ctx, SUCCESS, map[string]interface{}{}, message)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	Result(ctx, SUCCESS, data, "操作成功")
}

func OkWithDetailed(ctx *gin.Context, data interface{}, message string) {
	Result(ctx, SUCCESS, data, message)
}

func Fail(ctx *gin.Context) {
	Result(ctx, ERROR, map[string]interface{}{}, "操作失败")
}

func FailWithMessage(ctx *gin.Context, message string) {
	Result(ctx, ERROR, map[string]interface{}{}, message)
}

func FailWithDetailed(ctx *gin.Context, data interface{}, message string) {
	Result(ctx, ERROR, data, message)
}

func FailWithCode(ctx *gin.Context, code int, msg string) {
	Result(ctx, code, nil, msg)
}

func UnauthorizedWithDetailed(ctx *gin.Context, data interface{}, message string) {
	Result(ctx, Unauthorized, data, message)
}

func ErrorPage(ctx *gin.Context, status int, msg string) {
	ctx.HTML(http.StatusOK, "error.html", gin.H{
		"status":  status,
		"message": msg,
	})
}

func Redirect(ctx *gin.Context, location string) {
	ctx.Redirect(http.StatusFound, location)
}

func FailAsUnauthorized(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, Unauthorized, data, msg)
}

func FailWithNotFound(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, NotFound, data, msg)
}

func FailAorbiddensF(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, Forbidden, data, msg)
}

func FailInvalidParameter(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, InvalidParameter, data, msg)
}

func FailWithMissingParameter(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, MissingParameter, data, msg)
}

func FailWithInternalServerError(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, InternalServerError, data, msg)
}

func FailWithRecordExpired(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, RecordExpired, data, msg)
}

func FailWithDuplicate(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, Duplicate, data, msg)
}

func FailWithFetchResultError(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, FetchResultError, data, msg)
}

func FailWithInvalidRecordError(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, InvalidRecord, data, msg)
}

func FailWithBadParameter(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, BadParameter, data, msg)
}

func FailWithBalanceNotEnough(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, BalanceNotEnough, data, msg)
}

func FailSectionWithData(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, FailSection, data, msg)
}

func BadRequestWithData(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, BadRequest, data, msg)
}

func FailWithDataByCode(ctx *gin.Context, code int, msg string, data interface{}) {
	switch code {
	case Unauthorized:
		UnauthorizedWithDetailed(ctx, data, msg)
	case InvalidParameter:
		FailInvalidParameter(ctx, data, msg)
	case MissingParameter:
		FailWithMissingParameter(ctx, data, msg)
	case FailSection:
		FailSectionWithData(ctx, data, msg)
	case Forbidden:
		FailAorbiddensF(ctx, data, msg)
	case NotFound:
		FailWithNotFound(ctx, data, msg)
	case RecordExpired:
		FailWithRecordExpired(ctx, data, msg)
	case Duplicate:
		FailWithDuplicate(ctx, data, msg)
	case BadParameter:
		FailWithBadParameter(ctx, data, msg)
	case BalanceNotEnough:
		FailWithBalanceNotEnough(ctx, data, msg)
	case InternalServerError:
		FailWithInternalServerError(ctx, data, msg)
	case FetchResultError:
		FailWithFetchResultError(ctx, data, msg)
	case InvalidRecord:
		FailWithInvalidRecordError(ctx, data, msg)
	case BadRequest:
		BadRequestWithData(ctx, data, msg)
	default:
		Result(ctx, code, data, msg)
	}
}

func FailByCode(ctx *gin.Context, code int, msg string) {
	FailWithDataByCode(ctx, code, msg, "")
}
