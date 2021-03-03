package services

// HTTPService is a representation of the HTTP/2 REST API
type HTTPService interface {
	// Start will initialize the server
	Start()
	// Routes will register the endpoints
	Routes()
}
