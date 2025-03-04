package test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/bxcodec/faker/v3"
    "github.com/leofideliss/devices/domain"
    "github.com/leofideliss/devices/routes"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func BenchmarkRegisterDevice(b *testing.B){
    r:=routes.SetupRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        data := domain.RequestDevice{
            Owner: faker.Name(),
            DeviceId: faker.Word(),
            Title: "laptop work",
            Metadata: map[string]interface{}{
                "created_by": faker.Username(),
                "created_at": faker.Date(),
                "updated_by": faker.Username(),
                "updated_at": faker.Date(),
                "tags": []string{"example", "metadata", "mongodb"},
            },
            Expires_at: primitive.NewDateTimeFromTime(time.Now()),
        }
        jsonData , err := json.Marshal(data)
        if err != nil {
            b.Fatalf("Erro ao fazer Marshal: %v", err)
        }
        req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
        w := httptest.NewRecorder()
        r.ServeHTTP(w, req)
        if w.Code != http.StatusCreated {
            b.Errorf("Esperado status 201 Created, mas recebeu %s", w.Body)
        }
    }
}
