package answer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"questionAnswer/common/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
  db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
  if err != nil {
    t.Fatalf("failed to connect to sqlite: %v", err)
  }

  err = db.AutoMigrate(&model.Question{}, &model.Answer{})
  if err != nil {
    t.Fatalf("failed to migrate: %v", err)
  }

  return db
}

func TestPostAnswer(t *testing.T) {
  db := setupTestDB(t)
  h := &answerHandler{DB: db}

  question := model.Question{Text: "Test question"}
  db.Create(&question)

  body := `{"text":"my test answer"}`
  req := httptest.NewRequest(http.MethodPost, "/questions/1/answers/", bytes.NewBufferString(body))
  rr := httptest.NewRecorder()

  h.PostAnswer(rr, req)

  if rr.Code != http.StatusOK {
    t.Fatalf("expected status 200, got %d", rr.Code)
  }

  var resp map[string]interface{}
  json.Unmarshal(rr.Body.Bytes(), &resp)

  if resp["message"] != "Answer was added" {
    t.Fatalf("unexpected response message: %v", resp)
  }

  var answers []model.Answer
  db.Find(&answers)
  if len(answers) != 1 {
    t.Fatalf("expected 1 answer in DB, got %d", len(answers))
  }
}