package util

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	CallbackUrlHeader   = "X-Faas-Flow-Callback-Url"
	RequestIdHeader     = "X-Faas-Flow-Reqid"
	AuthSignatureHeader = "X-Hub-Signature"
)

// IsDagExportRequest check if dag export request
func IsDagExportRequest(query string) bool {

	fmt.Println("runtime/controller/util/util::IsDagExportRequest start")
	values, err := url.ParseQuery(query)
	if err != nil {
		return false
	}

	if strings.ToUpper(values.Get("export-dag")) == "TRUE" {
		return true
	}
	fmt.Println("runtime/controller/util/util::IsDagExportRequest end")
	return false
}

// GetStateRequestID check if state request and return the requestID
func GetStateRequestID(query string) string {

	fmt.Println("runtime/controller/util/util::GetStateRequestID start")
	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}

	fmt.Println("runtime/controller/util/util::GetStateRequestID end")
	return values.Get("state")
}

// GetStopRequestID check if stop request and return the requestID
func GetStopRequestID(query string) string {

	fmt.Println("runtime/controller/util/util::GetStopRequestID start")
	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}

	fmt.Println("runtime/controller/util/util::GetStopRequestID end")
	return values.Get("stop-flow")
}

// GetPauseRequestID check if pause request and return the requestID
func GetPauseRequestID(query string) string {

	fmt.Println("runtime/controller/util/util::GetPauseRequestID start")
	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}
	fmt.Println("runtime/controller/util/util::GetPauseRequestID end")
	return values.Get("pause-flow")
}

// GetResumeRequestID check if resume request and return the requestID
func GetResumeRequestID(query string) string {

	fmt.Println("runtime/controller/util/util::GetResumeRequestID start")
	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}
	fmt.Println("runtime/controller/util/util::GetResumeRequestID end")
	return values.Get("resume-flow")
}
