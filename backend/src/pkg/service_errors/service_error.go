package service_errors

type ServiceError struct {
	EndUserMsg string `json:"end_user_message"`
	Err        error  `json:"err"`
}

func (se *ServiceError) Error() string {
	return se.EndUserMsg
}
