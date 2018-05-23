package main

import (
	"time"

	"github.com/BcRikko/learning-goa/app"
	"github.com/goadesign/goa"
)

// TasksController implements the Tasks resource.
type TasksController struct {
	*goa.Controller
}

// NewTasksController creates a Tasks controller.
func NewTasksController(service *goa.Service) *TasksController {
	return &TasksController{Controller: service.NewController("TasksController")}
}

// Create runs the create action.
func (c *TasksController) Create(ctx *app.CreateTasksContext) error {
	// TasksController_Create: start_implement

	// Put your logic here

	return nil
	// TasksController_Create: end_implement
}

// Delete runs the delete action.
func (c *TasksController) Delete(ctx *app.DeleteTasksContext) error {
	// TasksController_Delete: start_implement

	// Put your logic here

	return nil
	// TasksController_Delete: end_implement
}

// List runs the list action.
func (c *TasksController) List(ctx *app.ListTasksContext) error {
	// TasksController_List: start_implement

	// Put your logic here

	res := app.XLearningGoaCollection{}
	return ctx.OK(res)
	// TasksController_List: end_implement
}

// Show runs the show action.
func (c *TasksController) Show(ctx *app.ShowTasksContext) error {
	// TasksController_Show: start_implement

	if ctx.TaskID == 0 {
		return ctx.NotFound()
	}

	res := &app.XLearningGoa{
		ID:        ctx.TaskID,
		Title:     "example task title",
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return ctx.OK(res)

	// TasksController_Show: end_implement
}

// Update runs the update action.
func (c *TasksController) Update(ctx *app.UpdateTasksContext) error {
	// TasksController_Update: start_implement

	// Put your logic here

	return nil
	// TasksController_Update: end_implement
}
