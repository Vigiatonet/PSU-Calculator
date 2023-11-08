package constants

const (
	AdminRoleName           = "admin"
	DefaultRoleName         = "default"
	DefaultUserName         = "admin"
	DefaultRedisKey         = "otp"
	AuthenTicationHeaderKey = "Authorization"

	// claims
	UserIdKey   = "UserId"
	FullNameKey = "FullName"
	UserNameKey = "UserName"
	PhoneKey    = "Phone"
	EmailKey    = "Email"
	RolesKey    = "Roles"
	ExpKey      = "Exp"
	AccessType  = "access_token"
	RefreshType = "refresh_token"
)
