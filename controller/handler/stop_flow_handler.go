package handler

import (
	"fmt"
	"log"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"
	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
)

func StopFlowHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {

	log.Printf("runtime/controller/handler/stop_flow_handler::StopFlowHandler start")
	log.Printf("Pausing request %s for flow %s\n", request.FlowName, request.RequestID)

	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	err := flowExecutor.Stop(request.RequestID)
	if err != nil {
		return fmt.Errorf("failed to stop request %s, check if request is active", request.RequestID)
	}

	response.Body = []byte("Successfully stopped request " + request.RequestID)
	log.Printf("runtime/controller/handler/stop_flow_handler::StopFlowHandler end")

	return nil
}
