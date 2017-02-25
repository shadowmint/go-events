package events

import (
	"reflect"
)

// EventHandler is a generic thread safe event handling system.
type EventHandler struct {
	groups map[reflect.Type]*eventGroup
}

// New returns a new event handler instance.
func New() *EventHandler {
	return &EventHandler{
		make(map[reflect.Type]*eventGroup)}
}

// Trigger pushes an event to all the handlers in the appropriate event group.
func (h *EventHandler) Trigger(event interface{}) {
	group := h.group(reflect.TypeOf(event))
	group.Trigger(event)
}

// Listen adds an event listener to the event handler group for the given event type.
func (h *EventHandler) Listen(event interface{}, handler func(event interface{}), flags ...int32) *EventBinding {
	var flag int32 = 0
	if len(flags) > 0 {
		for i := range flags {
			flag |= flags[i]
		}
	}
	group := h.group(reflect.TypeOf(event))
	return group.Listen(handler, flag)
}

// Clear removes all event handlers from the given event type.
func (h *EventHandler) Clear(event interface{}) {
	delete(h.groups, reflect.TypeOf(event))
}

// Release a specific event binding.
func (h *EventHandler) Release(binding *EventBinding) {
	binding.Release()
}

// group returns the event group for the given type, creating it if missing.
func (h *EventHandler) group(T reflect.Type) *eventGroup {
	if _, ok := h.groups[T]; !ok {
		h.groups[T] = newEventGroup()
	}
	return h.groups[T]
}