package util

import (
	"strings"
	"github.com/google/uuid"
)

func NewUuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
