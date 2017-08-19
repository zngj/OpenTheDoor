package util

import (
	"github.com/google/uuid"
	"strings"
)

func NewUuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
