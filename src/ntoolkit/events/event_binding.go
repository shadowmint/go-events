package events

// EventBinding keeps track of the details of a binding
type EventBinding struct {
	group   *eventGroup
	handler func(event interface{})
	once    bool // Is this a binding for a single invocation event
}

// Release this binding
func (b *EventBinding) Release() {
	if b.group == nil {
		return
	}
	b.group.Release(b)
	b.group = nil
}
