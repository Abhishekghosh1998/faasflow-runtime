package handler

import (
	"fmt"
	"log"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"

	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
)

func FlowStateHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {

	fmt.Println("runtime/controller/handler/flow_state_handler::FlowStateHandler start")
	log.Printf("Getting state of flow %s for request: %s\n", request.FlowName, request.RequestID)

	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	state, err := flowExecutor.GetState(request.RequestID)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("failed to get request state for %s, check if request is active", request.RequestID)
	}

	response.Body = []byte(state)
	fmt.Println("runtime/controller/handler/flow_state_handler::FlowStateHandler end")
	return nil
}
