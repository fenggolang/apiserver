package errno

// 用来统一存自定义的错误码
// 一个错误类型通常包含两部分: Code部分,用来唯一标识一个错误;
// Message部分,用来展示错误信息，这部分错误信息通常供前端直接展示;
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON Web Token"}

	// user errors
	ErrUserNotFound       = &Errno{Code: 20102, Message: "The user was not found."}
	ErrEncrypt            = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid       = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrrect = &Errno{Code: 20104, Message: "The password was incorrect."}
)
