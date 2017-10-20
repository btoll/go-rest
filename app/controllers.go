// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --regen=true
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// EventController is the controller interface for the Event actions.
type EventController interface {
	goa.Muxer
	Create(*CreateEventContext) error
	Delete(*DeleteEventContext) error
	List(*ListEventContext) error
	Show(*ShowEventContext) error
	Update(*UpdateEventContext) error
}

// MountEventController "mounts" a Event resource controller on the given service.
func MountEventController(service *goa.Service, ctrl EventController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/nmg/event/", ctrl.MuxHandler("preflight", handleEventOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/event/:id", ctrl.MuxHandler("preflight", handleEventOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/event/list", ctrl.MuxHandler("preflight", handleEventOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*EventPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("POST", "/nmg/event/", ctrl.MuxHandler("create", h, unmarshalCreateEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Create", "route", "POST /nmg/event/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("DELETE", "/nmg/event/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "Delete", "route", "DELETE /nmg/event/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("GET", "/nmg/event/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "List", "route", "GET /nmg/event/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("GET", "/nmg/event/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Event", "action", "Show", "route", "GET /nmg/event/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateEventContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*EventPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleEventOrigin(h)
	service.Mux.Handle("PUT", "/nmg/event/:id", ctrl.MuxHandler("update", h, unmarshalUpdateEventPayload))
	service.LogInfo("mount", "ctrl", "Event", "action", "Update", "route", "PUT /nmg/event/:id")
}

// handleEventOrigin applies the CORS response headers corresponding to the origin.
func handleEventOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateEventPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateEventPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &eventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateEventPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateEventPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &eventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// GameController is the controller interface for the Game actions.
type GameController interface {
	goa.Muxer
	Create(*CreateGameContext) error
	Delete(*DeleteGameContext) error
	List(*ListGameContext) error
	Show(*ShowGameContext) error
	Update(*UpdateGameContext) error
}

// MountGameController "mounts" a Game resource controller on the given service.
func MountGameController(service *goa.Service, ctrl GameController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/nmg/game/", ctrl.MuxHandler("preflight", handleGameOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/game/:id", ctrl.MuxHandler("preflight", handleGameOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/game/list", ctrl.MuxHandler("preflight", handleGameOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateGameContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*GamePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleGameOrigin(h)
	service.Mux.Handle("POST", "/nmg/game/", ctrl.MuxHandler("create", h, unmarshalCreateGamePayload))
	service.LogInfo("mount", "ctrl", "Game", "action", "Create", "route", "POST /nmg/game/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteGameContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleGameOrigin(h)
	service.Mux.Handle("DELETE", "/nmg/game/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Game", "action", "Delete", "route", "DELETE /nmg/game/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListGameContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleGameOrigin(h)
	service.Mux.Handle("GET", "/nmg/game/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Game", "action", "List", "route", "GET /nmg/game/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowGameContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleGameOrigin(h)
	service.Mux.Handle("GET", "/nmg/game/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Game", "action", "Show", "route", "GET /nmg/game/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateGameContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*GamePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleGameOrigin(h)
	service.Mux.Handle("PUT", "/nmg/game/:id", ctrl.MuxHandler("update", h, unmarshalUpdateGamePayload))
	service.LogInfo("mount", "ctrl", "Game", "action", "Update", "route", "PUT /nmg/game/:id")
}

// handleGameOrigin applies the CORS response headers corresponding to the origin.
func handleGameOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateGamePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateGamePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &gamePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateGamePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateGamePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &gamePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ImageController is the controller interface for the Image actions.
type ImageController interface {
	goa.Muxer
	Upload(*UploadImageContext) error
}

// MountImageController "mounts" a Image resource controller on the given service.
func MountImageController(service *goa.Service, ctrl ImageController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/nmg/image/:entity/:id", ctrl.MuxHandler("preflight", handleImageOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadImageContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Upload(rctx)
	}
	h = handleImageOrigin(h)
	service.Mux.Handle("POST", "/nmg/image/:entity/:id", ctrl.MuxHandler("upload", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Upload", "route", "POST /nmg/image/:entity/:id")
}

// handleImageOrigin applies the CORS response headers corresponding to the origin.
func handleImageOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SportController is the controller interface for the Sport actions.
type SportController interface {
	goa.Muxer
	Create(*CreateSportContext) error
	Delete(*DeleteSportContext) error
	List(*ListSportContext) error
	Show(*ShowSportContext) error
	Update(*UpdateSportContext) error
}

// MountSportController "mounts" a Sport resource controller on the given service.
func MountSportController(service *goa.Service, ctrl SportController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/nmg/sport/", ctrl.MuxHandler("preflight", handleSportOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/sport/:id", ctrl.MuxHandler("preflight", handleSportOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/sport/list", ctrl.MuxHandler("preflight", handleSportOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateSportContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SportPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSportOrigin(h)
	service.Mux.Handle("POST", "/nmg/sport/", ctrl.MuxHandler("create", h, unmarshalCreateSportPayload))
	service.LogInfo("mount", "ctrl", "Sport", "action", "Create", "route", "POST /nmg/sport/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteSportContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSportOrigin(h)
	service.Mux.Handle("DELETE", "/nmg/sport/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Sport", "action", "Delete", "route", "DELETE /nmg/sport/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListSportContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSportOrigin(h)
	service.Mux.Handle("GET", "/nmg/sport/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Sport", "action", "List", "route", "GET /nmg/sport/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowSportContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSportOrigin(h)
	service.Mux.Handle("GET", "/nmg/sport/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Sport", "action", "Show", "route", "GET /nmg/sport/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateSportContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SportPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSportOrigin(h)
	service.Mux.Handle("PUT", "/nmg/sport/:id", ctrl.MuxHandler("update", h, unmarshalUpdateSportPayload))
	service.LogInfo("mount", "ctrl", "Sport", "action", "Update", "route", "PUT /nmg/sport/:id")
}

// handleSportOrigin applies the CORS response headers corresponding to the origin.
func handleSportOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateSportPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateSportPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &sportPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateSportPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateSportPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &sportPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// TeamController is the controller interface for the Team actions.
type TeamController interface {
	goa.Muxer
	Create(*CreateTeamContext) error
	Delete(*DeleteTeamContext) error
	List(*ListTeamContext) error
	Show(*ShowTeamContext) error
	Update(*UpdateTeamContext) error
}

// MountTeamController "mounts" a Team resource controller on the given service.
func MountTeamController(service *goa.Service, ctrl TeamController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/nmg/team/", ctrl.MuxHandler("preflight", handleTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/team/:id", ctrl.MuxHandler("preflight", handleTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/team/list", ctrl.MuxHandler("preflight", handleTeamOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TeamPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleTeamOrigin(h)
	service.Mux.Handle("POST", "/nmg/team/", ctrl.MuxHandler("create", h, unmarshalCreateTeamPayload))
	service.LogInfo("mount", "ctrl", "Team", "action", "Create", "route", "POST /nmg/team/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleTeamOrigin(h)
	service.Mux.Handle("DELETE", "/nmg/team/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Team", "action", "Delete", "route", "DELETE /nmg/team/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleTeamOrigin(h)
	service.Mux.Handle("GET", "/nmg/team/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Team", "action", "List", "route", "GET /nmg/team/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleTeamOrigin(h)
	service.Mux.Handle("GET", "/nmg/team/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Team", "action", "Show", "route", "GET /nmg/team/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TeamPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleTeamOrigin(h)
	service.Mux.Handle("PUT", "/nmg/team/:id", ctrl.MuxHandler("update", h, unmarshalUpdateTeamPayload))
	service.LogInfo("mount", "ctrl", "Team", "action", "Update", "route", "PUT /nmg/team/:id")
}

// handleTeamOrigin applies the CORS response headers corresponding to the origin.
func handleTeamOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateTeamPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTeamPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &teamPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTeamPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTeamPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &teamPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// TeamOpeningConfigController is the controller interface for the TeamOpeningConfig actions.
type TeamOpeningConfigController interface {
	goa.Muxer
	Create(*CreateTeamOpeningConfigContext) error
	Delete(*DeleteTeamOpeningConfigContext) error
	List(*ListTeamOpeningConfigContext) error
	Show(*ShowTeamOpeningConfigContext) error
	Update(*UpdateTeamOpeningConfigContext) error
}

// MountTeamOpeningConfigController "mounts" a TeamOpeningConfig resource controller on the given service.
func MountTeamOpeningConfigController(service *goa.Service, ctrl TeamOpeningConfigController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/nmg/teamOpeningConfig/", ctrl.MuxHandler("preflight", handleTeamOpeningConfigOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/teamOpeningConfig/:id", ctrl.MuxHandler("preflight", handleTeamOpeningConfigOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/nmg/teamOpeningConfig/list", ctrl.MuxHandler("preflight", handleTeamOpeningConfigOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTeamOpeningConfigContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TeamOpeningConfigPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleTeamOpeningConfigOrigin(h)
	service.Mux.Handle("POST", "/nmg/teamOpeningConfig/", ctrl.MuxHandler("create", h, unmarshalCreateTeamOpeningConfigPayload))
	service.LogInfo("mount", "ctrl", "TeamOpeningConfig", "action", "Create", "route", "POST /nmg/teamOpeningConfig/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteTeamOpeningConfigContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleTeamOpeningConfigOrigin(h)
	service.Mux.Handle("DELETE", "/nmg/teamOpeningConfig/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "TeamOpeningConfig", "action", "Delete", "route", "DELETE /nmg/teamOpeningConfig/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTeamOpeningConfigContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleTeamOpeningConfigOrigin(h)
	service.Mux.Handle("GET", "/nmg/teamOpeningConfig/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "TeamOpeningConfig", "action", "List", "route", "GET /nmg/teamOpeningConfig/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTeamOpeningConfigContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleTeamOpeningConfigOrigin(h)
	service.Mux.Handle("GET", "/nmg/teamOpeningConfig/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "TeamOpeningConfig", "action", "Show", "route", "GET /nmg/teamOpeningConfig/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTeamOpeningConfigContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TeamOpeningConfigPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleTeamOpeningConfigOrigin(h)
	service.Mux.Handle("PUT", "/nmg/teamOpeningConfig/:id", ctrl.MuxHandler("update", h, unmarshalUpdateTeamOpeningConfigPayload))
	service.LogInfo("mount", "ctrl", "TeamOpeningConfig", "action", "Update", "route", "PUT /nmg/teamOpeningConfig/:id")
}

// handleTeamOpeningConfigOrigin applies the CORS response headers corresponding to the origin.
func handleTeamOpeningConfigOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateTeamOpeningConfigPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTeamOpeningConfigPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &teamOpeningConfigPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTeamOpeningConfigPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTeamOpeningConfigPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &teamOpeningConfigPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
