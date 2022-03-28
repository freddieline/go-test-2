package internal

import warehouse_api "go-test-2/internal/infrastructure/warehouse-api"

type Main struct {
	warehouseApi *warehouse_api.Client
}

func NewMain(warehouseApi *warehouse_api.Client) Main {
	return Main{
		warehouseApi: warehouseApi,
	}
}

func (m Main) ReturnTrue() bool {
	return true
}
