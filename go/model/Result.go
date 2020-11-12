package model

type Result struct {
	Success bool        `json:"success"`
	Obj     interface{} `json:"obj"`
}
