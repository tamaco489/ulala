package model

const DEFAULT_RESPONSE_202_MESSAGE string = "Request has been successfully accepted." // 202: リクエストを受け付けて処理が非同期で行われる場合

type DefaultResponse struct {
	Message string `json:"message"`
}
