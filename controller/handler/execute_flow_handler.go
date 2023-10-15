package handler

import (
	"fmt"
	"log"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"
	"github.com/Abhishekghosh1998/faasflow-runtime/controller/util"

	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
)

func ExecuteFlowHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {

	fmt.Println("runtime/controller/handler/execute_flow_handler::ExecuteFlowHandler start")
	log.Printf("Executing flow %s\n", request.FlowName)

	var stateOption executor.ExecutionStateOption

	callbackURL := request.GetHeader(util.CallbackUrlHeader)
	rawRequest := &executor.RawRequest{}
	rawRequest.Data = request.Body
	rawRequest.Query = request.RawQuery
	rawRequest.AuthSignature = request.GetHeader(util.AuthSignatureHeader)
	if request.RequestID != "" {
		rawRequest.RequestId = request.RequestID
	}
	stateOption = executor.NewRequest(rawRequest)

	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	resp, err := flowExecutor.Execute(stateOption)
	if err != nil {
		return fmt.Errorf("failed to execute request. %s", err.Error())
	}

	response.RequestID = flowExecutor.GetReqId()
	response.SetHeader(util.RequestIdHeader, response.RequestID)
	response.SetHeader(util.CallbackUrlHeader, callbackURL)
	response.Body = resp

	fmt.Println("runtime/controller/handler/execute_flow_handler::ExecuteFlowHandler end")
	return nil
}
