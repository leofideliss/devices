package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"github.com/leofideliss/devices/domain"
	"github.com/leofideliss/devices/routes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMain(m *testing.M) {
    err := godotenv.Load("../.env")
    fmt.Println(os.Getenv("DB_HOST"))
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }
    os.Setenv("TEST_ENV", "true")
    m.Run()
}

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
