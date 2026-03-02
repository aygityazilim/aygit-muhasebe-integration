package errors

import (
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
