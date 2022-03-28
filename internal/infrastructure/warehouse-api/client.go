package warehouse_api

type Client struct {
	baseUrl string
}

func NewBasicClient(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
	}
}
