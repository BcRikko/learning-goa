// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "Task": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/BcRikko/learning-goa/design
// --out=$(GOPATH)/src/github.com/BcRikko/learning-goa
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// TasksController is the controller interface for the Tasks actions.
type TasksController interface {
	goa.Muxer
	Create(*CreateTasksContext) error
	Delete(*DeleteTasksContext) error
	List(*ListTasksContext) error
	Show(*ShowTasksContext) error
	Update(*UpdateTasksContext) error
}

// MountTasksController "mounts" a Tasks resource controller on the given service.
func MountTasksController(service *goa.Service, ctrl TasksController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TaskPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/api/tasks", ctrl.MuxHandler("create", h, unmarshalCreateTasksPayload))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Create", "route", "POST /api/tasks")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/api/tasks/:taskID", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Delete", "route", "DELETE /api/tasks/:taskID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/tasks", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "List", "route", "GET /api/tasks")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/tasks/:taskID", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Show", "route", "GET /api/tasks/:taskID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTasksContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TaskPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("POST", "/api/tasks/:taskID", ctrl.MuxHandler("update", h, unmarshalUpdateTasksPayload))
	service.LogInfo("mount", "ctrl", "Tasks", "action", "Update", "route", "POST /api/tasks/:taskID")
}

// unmarshalCreateTasksPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTasksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &taskPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTasksPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTasksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &taskPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
