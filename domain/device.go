package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct{
    Id string `json:"id" bson:"_id"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Metadata map[string]any `json:"metadata"`
    Expires_at primitive.DateTime `json:"expire_at"`
}

type RequestDevice struct{
    DeviceId string `json:"deviceId" validate:"required,min=1,max=100"`
    Owner string `json:"owner" validate:"required,min=5,max=65"`
    Title string `json:"title" validate:"required,min=1,max=100"`
    Metadata map[string]any `json:"metadata"`
    Expires_at primitive.DateTime `json:"expire_at"`
}
