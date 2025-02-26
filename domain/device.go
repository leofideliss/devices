package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct{
    Id primitive.ObjectID `json:"id"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Metadata map[string]string `json:"metadata"`
    Expires_At time.Time `json:"expire_at"`
}
