package apiserver

//ApiServer
type APIServer struct {
	config *Config 
}

//New
func new(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

//Start
func (s *APIServer) Start() error {
	return nil
}