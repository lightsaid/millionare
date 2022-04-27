package errcodes

// ErrorInt 定义错误数据类型
type ErrorInt int

// 公共的错误码
const (
	// ErrSuccess 200: OK
	ErrSuccess ErrorInt = iota + 100000

	// ErrBindBody 400: 请求体绑定struct错误
	ErrBindBody // 100001

	// ErrValidation 400: 验证不通过
	ErrValidation // 100002

	// ErrTokenInvalid 401: token 无效
	ErrTokenInvalid // 100003

	// ErrTokenExpired 401: token 过期
	ErrTokenExpired // 100004

	// ErrUnknown 500: 服务内部错误
	ErrUnknown // 100005

	// ErrMissingArg 400: 缺少请求参数
	ErrMissingArg // 100006

	// ErrInvalidJSON 500: 不是有效的json （内部）
	ErrInvalidJSON // 100007

	// ErrEncodingJSON 500: json 编码错误（内部）
	ErrEncodingJSON // 100008

	// ErrDecodingJSON 500: json 解码错误（内部）
	ErrDecodingJSON // 100009

	// ErrAlreadyExist 400: 已经存在
	ErrAlreadyExist // 100010

	// ErrNotFound 404: 没找到
	ErrNotFound

	// ErrNotAllowed 405: 请求方法不对
	ErrNotAllowed
)

// 数据库错误公共码
const (
	ErrDatabase ErrorInt = iota + 100100
)
