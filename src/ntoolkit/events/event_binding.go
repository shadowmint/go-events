package events

// EventBinding keeps track of the details of a binding
type EventBinding struct {
	group   *eventGroup
	handler func(event interface{})
	once    bool
}

// Release this binding
func (b *EventBinding) Release() {
	b.group.Release(b)
}
