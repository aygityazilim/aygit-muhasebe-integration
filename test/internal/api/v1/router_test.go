
package v1_test

import (
	v1 "aygit-muhasebe-integration/internal/api/v1"
	"aygit-muhasebe-integration/internal/models"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// AuthMiddlewareMock, testlerde kullanılmak üzere c.Locals alanına sahte bir kullanıcı ekler.
func AuthMiddlewareMock(c *fiber.Ctx) error {
	companyID := uuid.New()
	user := &models.User{
		BaseModel: models.BaseModel{ID: uuid.New()},
		CompanyID: &companyID,
	}
	c.Locals("user", user)
	return c.Next()
}

// setupTestApp, testler için Fiber uygulamasını ve rotaları kurar.
func setupTestApp() *fiber.App {
	app := fiber.New()

	// v1 grubunu sahte auth middleware ile oluştur
	v1Group := app.Group("/v1", AuthMiddlewareMock)

	// Gerekli rota gruplarını ve endpoint'leri tanımla
	management := v1Group.Group("/management")
	management.Get("/identifications", v1.GetIdentifications) // Dışa aktarılmış handler'ı kullan

	return app
}

func TestGetIdentifications(t *testing.T) {
	// Test sunucusunu ve app'i kur
	app := setupTestApp()

	// Test isteğini oluştur
	req := httptest.NewRequest("GET", "/v1/management/identifications", nil)
	req.Header.Set("Content-Type", "application/json")

	// İsteği gerçekleştir
	resp, err := app.Test(req, -1) // -1 timeout'u devre dışı bırakır
	if err != nil {
		t.Fatalf("İstek gönderilirken hata oluştu: %v", err)
	}

	// Yanıtı kontrol et
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status kodu 200 OK olmalı")

	// Yanıt gövdesini oku
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("Yanıt gövdesi okunurken hata: %v", err)
	}

	// Yanıt gövdesini JSON olarak parse et ve doğrula
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	assert.NoError(t, err, "JSON parse edilirken hata olmamalı")

	// Beklenen yanıtı doğrula
	assert.Equal(t, "success", jsonResponse["status"], "'status' alanı 'success' olmalı")
	// data alanının boş bir slice olduğunu kontrol et
	data, ok := jsonResponse["data"].([]interface{})
	assert.True(t, ok, "'data' alanı bir slice olmalı")
	assert.Empty(t, data, "'data' alanı boş olmalı")
}
