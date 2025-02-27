package domain

import "time"

type Device struct{
    Id string `json:"id" bson:"_id"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Metadata map[string]any `json:"metadata"`
    Expires_at time.Time `json:"expire_at"`
}

type RequestDevice struct{
    DeviceId string `json:"deviceId"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Metadata map[string]any `json:"metadata"`
    Expires_at time.Time `json:"expire_at"`
}
