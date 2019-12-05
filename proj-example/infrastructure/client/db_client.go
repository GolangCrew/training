package client

// интерфейс для клиента
type Client interface {
	Open()
	Close()
	// etc.
}
