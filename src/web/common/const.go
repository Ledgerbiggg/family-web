package common

// 定义常见错误的枚举
const (
	Success = "10000" // 请求成功

	CaptchaError     = "20001" // 验证码错误
	LoginFailed      = "20002" // 用户名或密码错误
	UserIsExist      = "20003" // 用户名或密码错误
	PhoneFormat      = "20004"
	UserNotExist     = "20005"
	RealNameNotMatch = "20006"
	AdminRole        = "20007"
	Unknown          = "20008"

	BadRequest       = "40000" // 请求不合法，参数缺失或无效
	Unauthorized     = "40001" // 未授权访问，需要登录或验证
	Forbidden        = "40002" // 禁止访问，权限不足
	NotFound         = "40004" // 请求的资源未找到
	MethodNotAllowed = "40005" // 不支持的请求方法
	ValidationError  = "40010" // 参数验证失败
	ResourceLocked   = "40011" // 资源被锁定，无法操作\
	Database         = "40012" // 资源被锁定，无法操作\

	SystemServerError     = "50000" // 内部服务器错误
	ServiceUnavailable    = "50001" // 服务不可用
	Timeout               = "50003" // 请求超时
	NotFoundResource      = "50004" // 请求超时
	InviteLinkUsed        = "50005" // 请求超时
	InviteLinkUsedExpired = "50006" // 请求超时
	InviteLinkNotFound    = "50007" // 请求超时
	InviteRegister        = "50008" // 请求超时
	InviteRegisterExpired = "50009" // 请求超时
)

var (
	InviteRegisterExpiredError = NewKnownError(InviteRegisterExpired, "链接已过期")
	InviteRegisterError        = NewKnownError(InviteRegister, "非被邀请用户不允许邀请注册")
	CaptchaGetError            = NewKnownError(CaptchaError, "请获取验证码")
	UnknownError               = NewKnownError(Unknown, "未知异常，请联系站长")
	NotFoundResourceError      = NewKnownError(NotFoundResource, "资源不存在")
	InviteLinkNotFoundError    = NewKnownError(InviteLinkNotFound, "邀请链接不存在")
	InviteLinkUsedError        = NewKnownError(InviteLinkUsed, "邀请链接已经被使用")
	InviteLinkUsedExpiredError = NewKnownError(InviteLinkUsedExpired, "邀请链接已经过期")
	CaptchaErrorError          = NewKnownError(CaptchaError, "验证码错误")
	LoginErrorError            = NewKnownError(LoginFailed, "用户名或密码错误")
	UserIsExistError           = NewKnownError(UserIsExist, "用户已存在")
	UserIsNotExistError        = NewKnownError(UserNotExist, "用户不存在")
	RealNameNotMatchError      = NewKnownError(RealNameNotMatch, "真实姓名不匹配，请重新注册游客用户")
	DatabaseError              = NewKnownError(Database, "数据库执行出错")
	AdminRoleError             = NewKnownError(AdminRole, "用户不是管理员，不允许找回密码，请重新注册")
	PhoneFormatError           = NewKnownError(PhoneFormat, "手机号格式错误")
	BadRequestError            = NewKnownError(BadRequest, "请求不合法，参数缺失或无效")
	UnauthorizedError          = NewKnownError(Unauthorized, "无访问授权，需要登录或联系管理员新增权限")
	ForbiddenError             = NewKnownError(Forbidden, "禁止访问，权限不足")
	NotFoundError              = NewKnownError(NotFound, "请求的资源未找到")
	MethodNotAllowedError      = NewKnownError(MethodNotAllowed, "不支持的请求方法")
	SystemServerErrorError     = NewKnownError(SystemServerError, "内部服务器错误")
	ServiceUnavailableError    = NewKnownError(ServiceUnavailable, "服务不可用")
	ValidationErrorError       = NewKnownError(ValidationError, "参数验证失败")
	ResourceLockedError        = NewKnownError(ResourceLocked, "资源被锁定，无法操作")
	TimeoutError               = NewKnownError(Timeout, "请求超时")
)
