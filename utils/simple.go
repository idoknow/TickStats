package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CurrentTime() string {
	// Get current time
	return fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
}
func CurrentTimeStampSecond() int64 {
	// Get current time
	return time.Now().Unix()
}

func GenerateUUID(prefix string) string {
	// Generate a new UUID
	return fmt.Sprintf("%v", prefix+uuid.New().String()[:8])
}
