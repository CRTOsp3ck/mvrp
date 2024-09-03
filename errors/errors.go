package errors

import (
	"errors"
	"fmt"
)

// ============================
// General Errors
// ============================

var (
	ErrTypeUnknown        = errors.New("unknown error")
	ErrTypeInternal       = errors.New("internal error")
	ErrTypeInvalid        = errors.New("invalid error")
	ErrTypeTimeout        = errors.New("timeout error")
	ErrTypeCanceled       = errors.New("operation canceled")
	ErrTypeConflict       = errors.New("conflict error")
	ErrTypeUnavailable    = errors.New("service unavailable error")
	ErrTypeNotImplemented = errors.New("not implemented error")
	ErrTypeBadRequest     = errors.New("bad request error")
	ErrTypeUnauthorized   = errors.New("unauthorized error")
	ErrTypeForbidden      = errors.New("forbidden error")
	ErrTypeNotFound       = errors.New("not found error")
	ErrTypeConversion     = errors.New("conversion error")
	ErrTypeService        = errors.New("service error")
)

// ============================
// Type Errors
// ============================

var (
	ErrTypeAssertion       = errors.New("type assertion error")
	ErrTypeNil             = errors.New("nil value error")
	ErrTypeInvalidType     = errors.New("invalid type error")
	ErrTypeUnsupportedType = errors.New("unsupported type error")
)

// ============================
// File System Errors
// ============================

var (
	ErrTypeFileNotFound       = errors.New("file not found error")
	ErrTypeFileExists         = errors.New("file already exists error")
	ErrTypeFilePermission     = errors.New("file permission denied error")
	ErrTypeFileCorrupted      = errors.New("file corrupted error")
	ErrTypeDirectoryNotFound  = errors.New("directory not found error")
	ErrTypeInvalidPath        = errors.New("invalid file path error")
	ErrTypeDiskFull           = errors.New("disk full error")
	ErrTypeFileLocked         = errors.New("file is locked error")
	ErrTypeReadOnlyFilesystem = errors.New("read-only filesystem error")
)

// ============================
// I/O Errors
// ============================

var (
	ErrTypeIO              = errors.New("I/O error")
	ErrTypeEOF             = errors.New("EOF error")
	ErrTypeBrokenPipe      = errors.New("broken pipe error")
	ErrTypeConnectionReset = errors.New("connection reset error")
	ErrTypeWriteTimeout    = errors.New("write timeout error")
	ErrTypeReadTimeout     = errors.New("read timeout error")
	ErrTypeShortWrite      = errors.New("short write error")
	ErrTypeShortRead       = errors.New("short read error")
	ErrTypeBufferOverflow  = errors.New("buffer overflow error")
	ErrTypeBufferUnderflow = errors.New("buffer underflow error")
	ErrTypeStreamCorrupted = errors.New("stream corrupted error")
	ErrTypeStreamClosed    = errors.New("stream closed error")
)

// ============================
// Network Errors
// ============================

var (
	ErrTypeNetwork           = errors.New("network error")
	ErrTypeDNSResolution     = errors.New("DNS resolution error")
	ErrTypeConnectionFailed  = errors.New("connection failed error")
	ErrTypeHostUnreachable   = errors.New("host unreachable error")
	ErrTypeProtocolError     = errors.New("protocol error")
	ErrTypeSSLHandshake      = errors.New("SSL handshake error")
	ErrTypeProxyError        = errors.New("proxy error")
	ErrTypeRateLimited       = errors.New("rate limited error")
	ErrTypeTLSVerification   = errors.New("TLS verification error")
	ErrTypeSocketError       = errors.New("socket error")
	ErrTypeConnectionTimeout = errors.New("connection timeout error")
	ErrTypeKeepAliveFailed   = errors.New("keep-alive failed error")
)

// ============================
// Database Errors
// ============================

var (
	ErrTypeDatabase         = errors.New("database error")
	ErrTypeRecordNotFound   = errors.New("record not found error")
	ErrTypeDuplicateEntry   = errors.New("duplicate entry error")
	ErrTypeForeignKey       = errors.New("foreign key constraint error")
	ErrTypeUniqueConstraint = errors.New("unique constraint error")
	ErrTypeTransaction      = errors.New("transaction error")
	ErrTypeDBDeadlock       = errors.New("database deadlock error")
	ErrTypeSyntax           = errors.New("database syntax error")
	ErrTypeConnectionLost   = errors.New("database connection lost error")
	ErrTypeTimeoutExceeded  = errors.New("database timeout exceeded error")
	ErrTypeInvalidQuery     = errors.New("invalid database query error")
	ErrTypeNoRowsAffected   = errors.New("no rows affected error")
	ErrTypeMigration        = errors.New("database migration error")
)

// ============================
// JSON/XML Errors
// ============================

var (
	ErrTypeMarshal           = errors.New("marshal error")
	ErrTypeUnmarshal         = errors.New("unmarshal error")
	ErrTypeInvalidFormat     = errors.New("invalid format error")
	ErrTypeEncoding          = errors.New("encoding error")
	ErrTypeDecoding          = errors.New("decoding error")
	ErrTypeInvalidJSON       = errors.New("invalid JSON error")
	ErrTypeInvalidXML        = errors.New("invalid XML error")
	ErrTypeSchemaMismatch    = errors.New("schema mismatch error")
	ErrTypeUnsupportedFormat = errors.New("unsupported format error")
)

// ============================
// Validation Errors
// ============================

var (
	ErrTypeValidation       = errors.New("validation error")
	ErrTypeMissingField     = errors.New("missing field error")
	ErrTypeInvalidValue     = errors.New("invalid value error")
	ErrTypeConstraintFailed = errors.New("constraint failed error")
	ErrTypeFormatMismatch   = errors.New("format mismatch error")
	ErrTypeLengthExceeded   = errors.New("length exceeded error")
	ErrTypePatternMismatch  = errors.New("pattern mismatch error")
	ErrTypeRangeExceeded    = errors.New("range exceeded error")
	ErrTypeTypeMismatch     = errors.New("type mismatch error")
	ErrTypeDependencyFailed = errors.New("dependency failed error")
)

// ============================
// Authentication & Authorization Errors
// ============================

var (
	ErrTypeAuth           = errors.New("authentication error")
	ErrTypeAuthFailed     = errors.New("authentication failed error")
	ErrTypeAuthExpired    = errors.New("authentication expired error")
	ErrTypeAuthInvalid    = errors.New("authentication invalid error")
	ErrTypeAuthMissing    = errors.New("authentication missing error")
	ErrTypeAuthToken      = errors.New("authentication token error")
	ErrTypeAuthScope      = errors.New("authentication scope error")
	ErrTypeAuthPermission = errors.New("authentication permission error")
	ErrTypeAuthRole       = errors.New("authentication role error")
)

// ============================
// HTTP Errors
// ============================

var (
	ErrTypeHTTPBadRequest           = errors.New("HTTP bad request error")
	ErrTypeHTTPUnauthorized         = errors.New("HTTP unauthorized error")
	ErrTypeHTTPForbidden            = errors.New("HTTP forbidden error")
	ErrTypeHTTPNotFound             = errors.New("HTTP not found error")
	ErrTypeHTTPMethodNotAllowed     = errors.New("HTTP method not allowed error")
	ErrTypeHTTPNotAcceptable        = errors.New("HTTP not acceptable error")
	ErrTypeHTTPConflict             = errors.New("HTTP conflict error")
	ErrTypeHTTPInternalServerError  = errors.New("HTTP internal server error")
	ErrTypeHTTPServiceUnavailable   = errors.New("HTTP service unavailable error")
	ErrTypeHTTPGatewayTimeout       = errors.New("HTTP gateway timeout error")
	ErrTypeHTTPBadGateway           = errors.New("HTTP bad gateway error")
	ErrTypeHTTPUnsupportedMediaType = errors.New("HTTP unsupported media type error")
	ErrTypeHTTPTooManyRequests      = errors.New("HTTP too many requests error")
	ErrTypeHTTPTeapot               = errors.New("HTTP I'm a teapot error")
)

// ============================
// Concurrency Errors
// ============================

var (
	ErrTypeRaceCondition   = errors.New("race condition error")
	ErrTypeDeadlock        = errors.New("deadlock error")
	ErrTypeGoroutineLeak   = errors.New("goroutine leak error")
	ErrTypeChannelClosed   = errors.New("channel closed error")
	ErrTypeMutexLock       = errors.New("mutex lock error")
	ErrTypeMutexUnlock     = errors.New("mutex unlock error")
	ErrTypeWaitGroup       = errors.New("wait group error")
	ErrTypeContextCanceled = errors.New("context canceled error")
	ErrTypeContextDeadline = errors.New("context deadline exceeded error")
	ErrTypeAtomicOperation = errors.New("atomic operation error")
	ErrTypeSemaphore       = errors.New("semaphore error")
	ErrTypeBarrier         = errors.New("barrier error")
	ErrTypeCondition       = errors.New("condition variable error")
)

// ============================
// Custom Errors
// ============================

var (
	ErrTypeCustom = errors.New("custom error")
	// Add more custom error types here as needed
)

// ============================
// Error Handling Utilities
// ============================

// AnnotateError annotates an error with a specific message
func AnnotateError(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}

// WrapError wraps an error with additional context
func WrapError(err error, context string) error {
	return fmt.Errorf("%s: %w", context, err)
}

// Is checks if an error matches a specific error type
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Unwrap unwraps an error to get the underlying error
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// New creates a new error with a specific message
func New(msg string) error {
	return errors.New(msg)
}
