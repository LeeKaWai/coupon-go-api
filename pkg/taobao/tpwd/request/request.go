package request

type TpwdQueryRequest struct {
	/*
	   淘口令     */
	PasswordContent *string `json:"password_content" required:"true" `
}

func (s *TpwdQueryRequest) SetPasswordContent(v string) *TpwdQueryRequest {
	s.PasswordContent = &v
	return s
}

func (req *TpwdQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.PasswordContent != nil {
		paramMap["password_content"] = *req.PasswordContent
	}
	return paramMap
}

func (req *TpwdQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
