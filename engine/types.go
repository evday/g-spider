package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //url 对应 解析器
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
