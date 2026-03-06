package errors

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AppError represents a custom application error
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Predefined error codes
const (
	ErrCodeUnauthorized      = "AUTH_001"
	ErrCodeInvalidRequest    = "REQ_001"
	ErrCodeNotFound          = "RES_001"
	ErrCodeInternalServer    = "SYS_001"
	ErrCodeIntegrationFailed = "INT_001"
	ErrCodeDatabaseError     = "DB_001"
)

// NewError creates a new AppError
func NewError(status int, code string, message string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

// ErrorHandler is the global error handler for Fiber
func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*AppError); ok {
		return c.Status(e.Status).JSON(e)
	}

	// Default internal server error
	return c.Status(fiber.StatusInternalServerError).JSON(AppError{
		Code:    ErrCodeInternalServer,
		Message: "An unexpected error occurred",
	})
}

// TranslateHTTPStatus HTTP durum kodlarını Türkçe ve anlamlı açıklamalara çevirir
func TranslateHTTPStatus(status int) string {
	switch status {
	case http.StatusUnauthorized: // 401
		return "Yetkilendirme Hatası: API Anahtarı (API Key) hatalı, geçersiz veya bu işlem için yetkiniz yok. Lütfen .env dosyasındaki veya veritabanındaki NES_API_KEY bilgisini kontrol edin."
	case http.StatusForbidden: // 403
		return "Erişim Engellendi: Bu kaynağa erişim yetkiniz bulunmuyor. Hesabınızın aktif olduğundan ve IP kısıtlaması olmadığından emin olun."
	case http.StatusNotFound: // 404
		return "Kaynak Bulunamadı: İstek gönderilen API uç noktası (endpoint) mevcut değil veya yanlış."
	case http.StatusMethodNotAllowed: // 405
		return "Metod İzin Verilmedi: Bu uç nokta için kullanılan HTTP metodu (GET/POST/PUT vb.) geçersiz."
	case http.StatusUnsupportedMediaType: // 415
		return "Desteklenmeyen Medya Türü: Gönderilen veri formatı (Content-Type) hatalı."
	case http.StatusUnprocessableEntity: // 422
		return "Doğrulama Hatası: Gönderilen veriler şemaya uygun ancak iş mantığına göre geçersiz (örn: zorunlu alan eksik)."
	case http.StatusInternalServerError: // 500
		return "Sunucu Hatası: Özel Entegratör (NES) tarafında sistemsel bir hata oluştu."
	case http.StatusServiceUnavailable: // 503
		return "Servis Devre Dışı: Özel Entegratör sistemi şu anda bakımda veya yoğunluk nedeniyle kapalı."
	default:
		return fmt.Sprintf("Bilinmeyen HTTP Hatası (%d)", status)
	}
}

// FormatNESError NES API'sinden gelen hataları kullanıcı dostu bir formatta birleştirir
func FormatNESError(status int, body []byte) error {
	translation := TranslateHTTPStatus(status)
	if len(body) > 0 {
		return fmt.Errorf("NES API Hatası [%d]: %s | Detay: %s", status, translation, string(body))
	}
	return fmt.Errorf("NES API Hatası [%d]: %s", status, translation)
}
