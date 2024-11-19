package util

import (
	"fmt"
	"testing"
	"time"
)

func TestZone(t *testing.T) {
	fmt.Println(time.Now().UTC().Format(time.RFC3339))
}
