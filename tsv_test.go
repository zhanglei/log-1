package log

import (
	"testing"
)

func TestTSVLogger(t *testing.T) {
	logger := TSVLogger{}

	logger.New().
		TimestampMS().
		Bool(true).
		Bool(false).
		Float64(0.618).
		Int64(123).
		Uint64(456).
		Float32(12.0).
		Int(42).
		Int32(42).
		Int16(42).
		Int8(42).
		Uint32(42).
		Uint16(42).
		Uint8(42).
		Bytes([]byte("\"<,\t>?'")).
		Str("\"<,\t>?'").
		Msg()

}

func TestTSVSeparator(t *testing.T) {
	logger := TSVLogger{
		Separator: '¥',
	}

	logger.New().Msg()

	logger.New().
		TimestampMS().
		Bool(true).
		Bool(false).
		Float64(0.618).
		Int64(123).
		Uint64(456).
		Float32(12.0).
		Int(42).
		Int32(42).
		Int16(42).
		Int8(42).
		Uint32(42).
		Uint16(42).
		Uint8(42).
		Bytes([]byte("\"<,\t>?'")).
		Str("\"<,\t>?'").
		Msg()
}
