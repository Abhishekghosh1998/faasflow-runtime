package handler

import (
	"fmt"
	"log"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"
	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
	"github.com/Abhishekghosh1998/faasflow-sdk/exporter"
)

func GetDagHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {

	fmt.Println("runtime/controller/handler/get_dag_handler::GetDagHandler start")
	log.Printf("Exporting DAG of flow: %s\n", request.FlowName)

	flowExporter := exporter.CreateFlowExporter(ex)
	resp, err := flowExporter.Export()
	if err != nil {
		return fmt.Errorf("failed to export dag, error %v", err)
	}

	response.Body = resp
	fmt.Println("runtime/controller/handler/get_dag_handler::GetDagHandler end")
	return nil
}
