package utils

import (
	"time"
)

// SafeString safely dereferences *string to string
func SafeString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// SafeInt safely dereferences *int to int
func SafeInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

// SafeInt32 safely dereferences *int32 to int32
func SafeInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

// SafeInt64 safely dereferences *int64 to int64
func SafeInt64(i *int64) int64 {
	if i != nil {
		return *i
	}
	return 0
}

// SafeFloat32 safely dereferences *float32 to float32
func SafeFloat32(f *float32) float32 {
	if f != nil {
		return *f
	}
	return 0
}

// SafeFloat64 safely dereferences *float64 to float64
func SafeFloat64(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}

// SafeBool safely dereferences *bool to bool
func SafeBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// SafeTime safely dereferences *time.Time to time.Time
func SafeTime(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{}
}

// SafeTimeString formats *time.Time to string using layout
func SafeTimeString(t *time.Time, layout string) string {
	if t != nil {
		return FormatInLocation(layout, *t)
	}
	return ""
}
