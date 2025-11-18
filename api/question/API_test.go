package question

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
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("Can't to connect to SQLite: %v", err)
    }

    err = db.AutoMigrate(&model.Question{}, &model.Answer{})
    if err != nil {
        t.Fatalf("Failed to migrate: %v", err)
    }

    return db
}

func TestPostQuestion_BadJSON(t *testing.T) {
    db := setupTestDB(t)
    handler := &questionHandler{DB: db}

    req := httptest.NewRequest("POST", "/question", bytes.NewBufferString(`{invalid json}`))
    w := httptest.NewRecorder()

    http.HandlerFunc(handler.PostQuestion).ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", w.Code)
    }
}

func TestGETQuestions(t *testing.T) {
    db := setupTestDB(t)

    db.Create(&model.Question{Text: "Q1"})
    db.Create(&model.Question{Text: "Q2"})

    handler := &questionHandler{DB: db}
    req := httptest.NewRequest("GET", "/questions", nil)
    w := httptest.NewRecorder()

    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handler.GETQuestions(w, r, "")
    }).ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }

    var arr []model.Question
    json.Unmarshal(w.Body.Bytes(), &arr)

    if len(arr) != 2 {
        t.Fatalf("expected 2 questions, got %d", len(arr))
    }
}

func TestDELETEquestion(t *testing.T) {
    db := setupTestDB(t)

    q := model.Question{Text: "Delete me"}
    db.Create(&q)

    db.Create(&model.Answer{Text: "Answer1", Question_id: q.Id})
    db.Create(&model.Answer{Text: "Answer2", Question_id: q.Id})

    handler := &questionHandler{DB: db}

    req := httptest.NewRequest("DELETE", "/question/1", nil)
    w := httptest.NewRecorder()

    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handler.DELETEquestion(w, r, "1")
    }).ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }

    var count int64

    db.Model(&model.Question{}).Where("id = ?", q.Id).Count(&count)
    if count != 0 {
        t.Fatalf("question was not deleted")
    }

    db.Model(&model.Answer{}).Where("question_id = ?", q.Id).Count(&count)
    if count != 0 {
        t.Fatalf("answers were not deleted")
    }
}
