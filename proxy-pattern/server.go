package proxypattern

type Server interface {
	HandleRequest(string, string) (int, string)
}
