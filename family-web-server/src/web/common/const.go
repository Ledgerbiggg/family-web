package common

// 定义常见错误的枚举
const (
	Success = "10000" // 请求成功

	CaptchaError = "20001" // 验证码错误

	BadRequest       = "40000" // 请求不合法，参数缺失或无效
	Unauthorized     = "40001" // 未授权访问，需要登录或验证
	Forbidden        = "40002" // 禁止访问，权限不足
	NotFound         = "40004" // 请求的资源未找到
	MethodNotAllowed = "40005" // 不支持的请求方法
	ValidationError  = "40010" // 参数验证失败
	ResourceLocked   = "40011" // 资源被锁定，无法操作\

	SystemServerError  = "50000" // 内部服务器错误
	ServiceUnavailable = "50001" // 服务不可用
	Timeout            = "50003" // 请求超时
)

var (
	CaptchaErrorError       = NewKnownError(CaptchaError, "验证码错误")
	BadRequestError         = NewKnownError(BadRequest, "请求不合法，参数缺失或无效")
	UnauthorizedError       = NewKnownError(Unauthorized, "未授权访问，需要登录或验证")
	ForbiddenError          = NewKnownError(Forbidden, "禁止访问，权限不足")
	NotFoundError           = NewKnownError(NotFound, "请求的资源未找到")
	MethodNotAllowedError   = NewKnownError(MethodNotAllowed, "不支持的请求方法")
	SystemServerErrorError  = NewKnownError(SystemServerError, "内部服务器错误")
	ServiceUnavailableError = NewKnownError(ServiceUnavailable, "服务不可用")
	ValidationErrorError    = NewKnownError(ValidationError, "参数验证失败")
	ResourceLockedError     = NewKnownError(ResourceLocked, "资源被锁定，无法操作")
	TimeoutError            = NewKnownError(Timeout, "请求超时")
)
