package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountInfo struct {
	ID           primitive.ObjectID `json:"_id,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	AccountAuth  *AccountAuth       `json:"account_auth,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone        int                `json:"phone,omitempty" bson:"phone,omitempty"`
	ReserveEmail string             `json:"reserve_email,omitempty" bson:"reserve_email,omitempty"`
	Owner        string             `json:"owner,omitempty" bson:"owner,omitempty"`
	Notice       string             `json:"notice,omitempty" bson:"_id,omitempty"`
}

type AccountAuth struct {
	Login    string `json:"login,omitempty" bson:"login,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
