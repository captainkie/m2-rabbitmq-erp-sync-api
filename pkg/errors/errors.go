package errors

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "VALIDATION_ERROR"
	// ErrorTypeNotFound represents not found errors
	ErrorTypeNotFound ErrorType = "NOT_FOUND"
	// ErrorTypeUnauthorized represents unauthorized errors
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	// ErrorTypeForbidden represents forbidden errors
	ErrorTypeForbidden ErrorType = "FORBIDDEN"
	// ErrorTypeInternal represents internal server errors
	ErrorTypeInternal ErrorType = "INTERNAL_ERROR"
	// ErrorTypeBadRequest represents bad request errors
	ErrorTypeBadRequest ErrorType = "BAD_REQUEST"
	// ErrorTypeConflict represents conflict errors
	ErrorTypeConflict ErrorType = "CONFLICT"
	// ErrorTypeRateLimit represents rate limit errors
	ErrorTypeRateLimit ErrorType = "RATE_LIMIT"
)

// AppError represents an application error
type AppError struct {
	Type      ErrorType   `json:"type"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Timestamp time.Time   `json:"timestamp"`
	File      string      `json:"file,omitempty"`
	Line      int         `json:"line,omitempty"`
	Function  string      `json:"function,omitempty"`
	Details   interface{} `json:"details,omitempty"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

// New creates a new AppError
func New(errType ErrorType, message string, code int, details interface{}) *AppError {
	_, file, line, _ := runtime.Caller(1)
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return &AppError{
		Type:      errType,
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
		File:      file,
		Line:      line,
		Function:  fn.Name(),
		Details:   details,
	}
}

// ErrorHandler is a centralized error handling middleware
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			switch e := err.(type) {
			case *AppError:
				c.JSON(e.Code, e)
			default:
				// Convert unknown errors to internal server error
				appErr := New(
					ErrorTypeInternal,
					"An unexpected error occurred",
					http.StatusInternalServerError,
					nil,
				)
				c.JSON(appErr.Code, appErr)
			}
		}
	}
}

// Common error constructors
func NewValidationError(message string, details interface{}) *AppError {
	return New(ErrorTypeValidation, message, http.StatusBadRequest, details)
}

func NewNotFoundError(message string) *AppError {
	return New(ErrorTypeNotFound, message, http.StatusNotFound, nil)
}

func NewUnauthorizedError(message string) *AppError {
	return New(ErrorTypeUnauthorized, message, http.StatusUnauthorized, nil)
}

func NewForbiddenError(message string) *AppError {
	return New(ErrorTypeForbidden, message, http.StatusForbidden, nil)
}

func NewInternalError(message string) *AppError {
	return New(ErrorTypeInternal, message, http.StatusInternalServerError, nil)
}

func NewBadRequestError(message string, details interface{}) *AppError {
	return New(ErrorTypeBadRequest, message, http.StatusBadRequest, details)
}

func NewConflictError(message string) *AppError {
	return New(ErrorTypeConflict, message, http.StatusConflict, nil)
}

func NewRateLimitError(message string, retryAfter int) *AppError {
	return New(ErrorTypeRateLimit, message, http.StatusTooManyRequests, map[string]int{
		"retry_after": retryAfter,
	})
}
