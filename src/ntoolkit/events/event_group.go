package events

import (
	"container/list"
	"ntoolkit/errors"
)

// eventGroup keeps track of all the handlers for a given type
type eventGroup struct {
	handlers *list.List
}

// newEventGroup returns a new group
func newEventGroup() *eventGroup {
	return &eventGroup{
		handlers: list.New()}
}

// Listen adds a new listener
func (group *eventGroup) Listen(handler func(event interface{}), flags int32) *EventBinding {
	binding := &EventBinding{
		group:   group,
		handler: handler,
		once:    (flags & RunOnce) == RunOnce}
	group.handlers.PushBack(binding)
	return binding
}

// Trigger fires an event handler to all the handlers.
func (group *eventGroup) Trigger(event interface{}) {
	count := group.handlers.Len()
	for i := 0; i < count; i++ {
		el := group.handlers.Front()
		binding := el.Value.(*EventBinding)
		group.Execute(event, binding.handler)
		if !binding.once {
			group.handlers.MoveToBack(el)
		} else {
			group.handlers.Remove(el)
		}
	}
}

// Execute safely executes a handler
func (group *eventGroup) Execute(event interface{}, handler func(event interface{})) (err error) {
	defer (func() {
		if r := recover(); r != nil {
			if eval, ok := r.(error); ok {
				err = errors.Fail(ErrHandlerFailed{}, eval, "Failing invoking event handler")
			} else {
				err = errors.Fail(ErrHandlerFailed{}, nil, "Failing invoking event handler")
			}
		}
	})()
	handler(event)
	return nil
}

// Release an event binding.
func (group *eventGroup) Release(binding *EventBinding) {
	count := group.handlers.Len()
	for i := 0; i < count; i++ {
		el := group.handlers.Front()
		bp := el.Value.(*EventBinding)
		if bp == binding {
			group.handlers.Remove(el)
		} else {
			group.handlers.MoveToBack(el)
		}
	}
}
