package service

import (
	"github.com/leofideliss/devices/domain"
	"github.com/leofideliss/devices/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceService struct {
    repo *repository.MongoRepository[domain.Device]
}

func NewDeviceService(db *mongo.Database) *DeviceService{
    return &DeviceService{
        repo: repository.NewMongoRepository[domain.Device](db,"devices"),
    }
}

