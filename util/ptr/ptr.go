package ptr

import "time"

type IPtrUtil interface {
	DurationPtr(d time.Duration) *time.Duration
	TimePtr(t time.Time) *time.Time
	StrPtr(s string) *string
	IntPtr(i int) *int
	BoolPtr(b bool) *bool
	Float64Ptr(f float64) *float64
	Float32Ptr(f float32) *float32
	Int64Ptr(i int64) *int64
	Int32Ptr(i int32) *int32
	Int16Ptr(i int16) *int16
	Int8Ptr(i int8) *int8
	Uint64Ptr(i uint64) *uint64
	Uint32Ptr(i uint32) *uint32
	Uint16Ptr(i uint16) *uint16
	Uint8Ptr(i uint8) *uint8
	BytePtr(b byte) *byte
	RunePtr(r rune) *rune
}

type PtrUtil struct{}

func (p *PtrUtil) DurationPtr(d time.Duration) *time.Duration {
	return &d
}

func (p *PtrUtil) TimePtr(t time.Time) *time.Time {
	return &t
}

func (p *PtrUtil) StrPtr(s string) *string {
	return &s
}

func (p *PtrUtil) IntPtr(i int) *int {
	return &i
}

func (p *PtrUtil) BoolPtr(b bool) *bool {
	return &b
}

func (p *PtrUtil) Float64Ptr(f float64) *float64 {
	return &f
}

func (p *PtrUtil) Float32Ptr(f float32) *float32 {
	return &f
}

func (p *PtrUtil) Int64Ptr(i int64) *int64 {
	return &i
}

func (p *PtrUtil) Int32Ptr(i int32) *int32 {
	return &i
}

func (p *PtrUtil) Int16Ptr(i int16) *int16 {
	return &i
}

func (p *PtrUtil) Int8Ptr(i int8) *int8 {
	return &i
}

func (p *PtrUtil) Uint64Ptr(i uint64) *uint64 {
	return &i
}

func (p *PtrUtil) Uint32Ptr(i uint32) *uint32 {
	return &i
}

func (p *PtrUtil) Uint16Ptr(i uint16) *uint16 {
	return &i
}

func (p *PtrUtil) Uint8Ptr(i uint8) *uint8 {
	return &i
}

func (p *PtrUtil) BytePtr(b byte) *byte {
	return &b
}

func (p *PtrUtil) RunePtr(r rune) *rune {
	return &r
}
