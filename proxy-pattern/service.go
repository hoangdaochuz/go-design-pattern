package proxypattern

type Service struct { // real service

}

// implement the server interface
func (s *Service) HandleRequest(url, methoad string) (int, string) {
	return 200, "OK " + url
}
