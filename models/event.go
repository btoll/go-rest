package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type EventModel struct{}

func (m *EventModel) AllocateID(ctx *Context) error {
	return AllocateSequentialID(ctx)
}

func (m *EventModel) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateEventContext:
		return &Context{
			Kind:    constants.DB_EVENT,
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteEventContext:
		return &Context{
			Kind:   constants.DB_EVENT,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListEventContext:
		return &Context{
			Kind:   constants.DB_EVENT,
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowEventContext:
		return &Context{
			Kind:   constants.DB_EVENT,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateEventContext:
		return &Context{
			Kind:    constants.DB_EVENT,
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *EventModel) GetModelInstance() interface{} {
	return &app.EventMedia{}
}

func (m *EventModel) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := app.EventMediaCollection{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *EventModel) SetModel(ctx *Context) error {
	rec := &app.EventMedia{}

	copier.Copy(rec, ctx.Payload)
	rec.ID = ctx.ID
	_, err := datastore.Put(ctx.GaeCtx, GetKey(ctx), rec)

	return err
}
