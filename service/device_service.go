package service

import (
	"context"

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

func (d *DeviceService) GetDevice (ctx context.Context , id string ) (*domain.Device , error){
    return d.repo.FindById(ctx,id)
}

func (d *DeviceService) RegisterDevice (ctx context.Context , device *domain.Device) (*mongo.InsertOneResult , error){
    return d.repo.Insert(ctx,device)
}
