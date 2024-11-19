package model

type UserFirebaseAuthentication struct {
	FirebaseUID string `json:"firebase_uid"`
	UUID        string `json:"uuid"`
}

type SignUpRequest struct {
	FirebaseToken string `json:"firebase_token"`
	Email         string `json:"email"`
}

type SignUpResponse struct {
	FirebaseUID string `json:"firebase_uid"`
}

type SignInRequest struct {
	FirebaseToken string `json:"firebase_token"`
	Email         string `json:"email"`
}

type SignInResponse struct {
	UID uint32 `json:"uid"`
}

type SignOutRequest struct {
	FirebaseToken string `json:"firebase_token"`
}
