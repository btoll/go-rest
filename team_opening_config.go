package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
)

// TeamOpeningConfigController implements the TeamOpeningConfig resource.
type TeamOpeningConfigController struct {
	*goa.Controller
}

// NewTeamOpeningConfigController creates a TeamOpeningConfig controller.
func NewTeamOpeningConfigController(service *goa.Service) *TeamOpeningConfigController {
	return &TeamOpeningConfigController{Controller: service.NewController("TeamOpeningConfigController")}
}

// Create runs the create action.
func (c *TeamOpeningConfigController) Create(ctx *app.CreateTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Create: start_implement

	id, err := models.Create(constants.DB_TEAM_OPENING_CONFIG, ctx)

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "create")
	}

	return ctx.OKTiny(&app.TeamOpeningConfigMediaTiny{id})

	// TeamOpeningConfigController_Create: end_implement
}

// Delete runs the delete action.
func (c *TeamOpeningConfigController) Delete(ctx *app.DeleteTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Delete: start_implement

	if err := models.Delete(constants.DB_TEAM_OPENING_CONFIG, ctx); err != nil {
		return goa.ErrInternal(err, "endpoint", "delete")
	}

	return ctx.NoContent()

	// TeamOpeningConfigController_Delete: end_implement
}

// List runs the list action.
func (c *TeamOpeningConfigController) List(ctx *app.ListTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_List: start_implement

	coll, err := models.List(constants.DB_TEAM_OPENING_CONFIG, ctx)

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "list")
	}

	return ctx.OK(coll.(app.TeamOpeningConfigMediaCollection))

	// TeamOpeningConfigController_List: end_implement
}

// Show runs the show action.
func (c *TeamOpeningConfigController) Show(ctx *app.ShowTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Show: start_implement

	model, err := models.Read(constants.DB_TEAM_OPENING_CONFIG, ctx)

	if err != nil {
		return goa.ErrBadRequest(err, "endpoint", "show")
	}

	return ctx.OK(model.(*app.TeamOpeningConfigMedia))

	// TeamOpeningConfigController_Show: end_implement
}

// Update runs the update action.
func (c *TeamOpeningConfigController) Update(ctx *app.UpdateTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Update: start_implement

	if err := models.Update(constants.DB_TEAM_OPENING_CONFIG, ctx); err != nil {
		return goa.ErrInternal(err, "endpoint", "update")
	}

	return ctx.NoContent()

	// TeamOpeningConfigController_Update: end_implement
}
