package internal

import (
	"github.com/stretchr/testify/assert"
	warehouse_api "go-test-2/internal/infrastructure/warehouse-api"
	"testing"
)

func TestTheMain(t *testing.T) {

	main := NewMain(warehouse_api.NewBasicClient("localhost"))
	assert.Equal(t, true, main.ReturnTrue())
}
