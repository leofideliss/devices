package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct{
    Id string `json:"id" bson:"_id"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Metadata map[string]any `json:"metadata"`
    Expires_at primitive.DateTime `json:"expire_at" example:"2024-12-03T12:15:20Z"`
}

type RequestDevice struct{
    DeviceId string `json:"deviceId" validate:"required,min=1,max=100"`
    Owner string `json:"owner" validate:"required,min=5,max=65"`
    Title string `json:"title" validate:"required,min=1,max=100"`
    Metadata map[string]any `json:"metadata"`
    Expires_at primitive.DateTime `json:"expire_at" example:"2024-12-03T12:15:20Z"`
}

// Modelo para usar no swagger

// @Model DeviceDetails
// @Schema definition: DeviceDetails
// @Description Modelo que representa um dispositivo
type RequestDeviceSwagger struct{
    DeviceId string `json:"deviceId" bson:"_id"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Metadata map[string]any `json:"metadata"`
    Expires_at string `json:"expire_at" example:"2024-12-03T12:15:20Z"`
}
