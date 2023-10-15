package http

import (
	"fmt"
	"net/http"

	"github.com/Abhishekghosh1998/faasflow-runtime/controller/handler"

	runtime "github.com/Abhishekghosh1998/faasflow-runtime"

	"github.com/julienschmidt/httprouter"
)

func router(runtime runtime.Runtime) http.Handler {

	fmt.Println("runtime/controller/http/router::router start")
	router := httprouter.New()
	// router.POST("/flow/execute", newRequestHandlerWrapper(runtime, handler.ExecuteFlowHandler))
	router.POST("/flow/:id/forward", newRequestHandlerWrapper(runtime, handler.PartialExecuteFlowHandler))
	router.POST("/flow/:id/pause", newRequestHandlerWrapper(runtime, handler.PauseFlowHandler))
	router.POST("/flow/:id/resume", newRequestHandlerWrapper(runtime, handler.ResumeFlowHandler))
	router.POST("/flow/:id/stop", newRequestHandlerWrapper(runtime, handler.StopFlowHandler))
	router.GET("/flow/:id/state", newRequestHandlerWrapper(runtime, handler.FlowStateHandler))
	router.POST("/", newRequestHandlerWrapper(runtime, handler.LegacyRequestHandler))
	router.GET("/", newRequestHandlerWrapper(runtime, handler.LegacyRequestHandler))

	fmt.Println("runtime/controller/http/router::router end")
	return router
}
