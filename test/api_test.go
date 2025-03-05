package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/leofideliss/devices/domain"
	"github.com/leofideliss/devices/routes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegisterDevice(t *testing.T){   
    r:=routes.SetupRouter()
	data := domain.RequestDevice{
		Owner:    faker.Name(),
		DeviceId: faker.Word(),
		Title:    "laptop work",
		Metadata: map[string]interface{}{
			"created_by": faker.Username(),
			"created_at": faker.Date(),
			"updated_by": faker.Username(),
			"updated_at": faker.Date(),
			"tags":       []string{"example", "metadata", "mongodb"},
		},
		Expires_at: primitive.NewDateTimeFromTime(time.Now()),
	}
    jsonData , _ := json.Marshal(data)
    req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
    r.ServeHTTP(w,req)    
    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetDevice(t *testing.T){
    owner := faker.Name()
    deviceId := faker.Word()
    r:=routes.SetupRouter()
    createStatusCode := createDevice(r,owner,deviceId)
    if createStatusCode == 201{
        req, _ := http.NewRequest("GET", "/" + deviceId + "?owner=" + owner,nil)
        w := httptest.NewRecorder()
        r.ServeHTTP(w,req)
        assert.Equal(t,http.StatusOK,w.Code)
    } else {
        t.Error()
    }
}

func TestDeleteDevice(t *testing.T){
    owner := faker.Name()
    deviceId := faker.Word()
    r:=routes.SetupRouter()
    createStatusCode := createDevice(r,owner,deviceId)
    if createStatusCode == 201{
        req, _ := http.NewRequest("DELETE", "/" + deviceId + "?owner=" + owner,nil)
        w := httptest.NewRecorder()
        r.ServeHTTP(w,req)
        assert.Equal(t,http.StatusOK,w.Code)
    } else {
        t.Error()
    }
}

func TestUpdateDevice(t *testing.T){
    owner := faker.Name()
    deviceId := faker.Word()
    r:=routes.SetupRouter()
    createStatusCode := createDevice(r,owner,deviceId)
    if createStatusCode == 201{
        data := domain.RequestDevice{
            Owner:    faker.Name(),
            DeviceId: faker.Word(),
            Title:    "laptop work",
            Metadata: map[string]interface{}{
                "created_by": faker.Username(),
                "created_at": faker.Date(),
                "updated_by": faker.Username(),
                "updated_at": faker.Date(),
                "tags":       []string{"example", "metadata", "mongodb"},
            },
            Expires_at: primitive.NewDateTimeFromTime(time.Now()),
        }
        jsonData , _ := json.Marshal(data)
        req, _ := http.NewRequest("PATCH", "/" + deviceId + "?owner=" + owner, bytes.NewBuffer(jsonData))
        w := httptest.NewRecorder()
        r.ServeHTTP(w,req)
        assert.Equal(t,http.StatusOK,w.Code)
    } else {
        t.Error()
    }
}

func TestListDevice(t *testing.T){
    owner := faker.Name()
    r := routes.SetupRouter()
    req, _ := http.NewRequest("GET", "/list?owner=" + owner + "&limit=1" + "&page=1",nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w,req)
    assert.Equal(t,http.StatusOK,w.Code)
}

func createDevice(r *gin.Engine, owner , deviceId string) int{
    data := domain.RequestDevice{
		Owner:    owner,
		DeviceId: deviceId,
		Title:    "laptop work",
		Metadata: map[string]interface{}{
			"created_by": faker.Username(),
			"created_at": faker.Date(),
			"updated_by": faker.Username(),
			"updated_at": faker.Date(),
			"tags":       []string{"example", "metadata", "mongodb"},
		},
		Expires_at: primitive.NewDateTimeFromTime(time.Now()),
	}
    jsonData , _ := json.Marshal(data)
    reqCreate, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	responseCreate := httptest.NewRecorder()
    r.ServeHTTP(responseCreate,reqCreate)
    return responseCreate.Code
}
