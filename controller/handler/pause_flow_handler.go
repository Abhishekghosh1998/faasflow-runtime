package handler

import (
	"fmt"
	"log"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"
	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
)

func PauseFlowHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {

	fmt.Println("runtime/controller/handler/pause_flow_handler::PauseFlowHandler start")
	log.Printf("Pausing request %s of flow %s\n", request.RequestID, request.FlowName)

	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	err := flowExecutor.Pause(request.RequestID)
	if err != nil {
		return fmt.Errorf("failed to pause request %s, check if request is active", request.RequestID)
	}

	response.Body = []byte("Successfully paused request " + request.RequestID)
	fmt.Println("runtime/controller/handler/pause_flow_handler::PauseFlowHandler end")

	return nil
}
