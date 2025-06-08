package proxypattern

type NginxProxy struct {
	server            Server
	rateLimiter       map[string]int
	maxAllowedRequest int
}

func NewNginxProxy(server Server, rateLimiter map[string]int, maxAllowedRequest int) Server {
	return &NginxProxy{
		server:            server,
		rateLimiter:       rateLimiter,
		maxAllowedRequest: maxAllowedRequest,
	}
}

func (n *NginxProxy) HandleRequest(url, method string) (int, string) {
	if !n.CheckRateLimit(url) {
		return 429, "Too many requests"
	}
	return n.server.HandleRequest(url, method)
}

func (n *NginxProxy) CheckRateLimit(url string) bool {
	count, ok := n.rateLimiter[url]
	if !ok {
		n.rateLimiter[url] = 1
		return true
	}
	if count >= n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] += 1
	return true
}
