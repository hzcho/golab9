package models

type Account struct {
	Id       int64  `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"pass_hash"`
}

type RegisterReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	AccountId uint64 `json:"account_id"`
}
