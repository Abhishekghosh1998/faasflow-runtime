package handler

import (
	"errors"
	"fmt"
	"log"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"
	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
)

func PartialExecuteFlowHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {

	fmt.Println("runtime/controller/handler/partial_execute_flow_handler::PartialExecuteFlowHandler start")
	log.Printf("Partially executing flow %s, for id %s\n", request.FlowName, request.RequestID)

	var stateOption executor.ExecutionStateOption

	if request.RequestID == "" {
		return errors.New("request ID must be set in partial request")
	}
	partialState, err := executor.DecodePartialReq(request.Body)
	if err != nil {
		return errors.New("failed to decode partial state")
	}
	stateOption = executor.PartialRequest(partialState)

	// Create a flow executor with provided executor
	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	resp, err := flowExecutor.Execute(stateOption)
	if err != nil {
		return fmt.Errorf("failed to execute request. %s", err.Error())
	}

	response.Body = resp
	fmt.Println("runtime/controller/handler/partial_execute_flow_handler::PartialExecuteFlowHandler end")
	return nil
}
