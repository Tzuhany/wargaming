package errno

const (
	// For api-gateway

	StatusSuccessCode = 0
	StatusSuccessMsg  = "ok"

	// For microservices

	SuccessCode = 10000
	SuccessMsg  = "ok"

	// Error

	ServiceErrorCode           = 10001 // 未知微服务错误
	ParamErrorCode             = 10002 // 参数错误
	AuthorizationFailedErrCode = 10003 // 鉴权失败
	UnexpectedTypeErrorCode    = 10004 // 未知类型

	MatchTimeoutErrorCode = 20001
)

var (
	// Success

	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(int64(AuthorizationFailedErrCode), "authorization failed")

	UserExistedError  = NewErrNo(ParamErrorCode, "user existed")
	UserNotFoundError = NewErrNo(ParamErrorCode, "user not found")

	// Websocket

	MatchTimeoutError = NewErrNo(MatchTimeoutErrorCode, "match timeout")
)
