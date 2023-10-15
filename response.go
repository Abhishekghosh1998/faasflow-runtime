package runtime

import "fmt"

type Response struct {
	RequestID string
	Header    map[string][]string
	Body      []byte
}

func (response *Response) SetHeader(header string, value string) {
	fmt.Println("runtime/response::SetHeader start")
	response.Header[header] = []string{value}
	fmt.Println("runtime/response::SetHeader end")
}
