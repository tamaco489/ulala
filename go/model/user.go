package model

type User struct {
	UID      uint32 `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	AuthUUID string `json:"auth_uuid"`
}

type CreateUserResponse struct{} // TODO 要件が固まるまで空で返却、UIDをUUIDに変えて変更しても良いかもしれない

type GetUserResponse struct {
	UID   uint32 `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
