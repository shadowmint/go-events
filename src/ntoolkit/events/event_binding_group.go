package events

// EventBindingGroup aggregates a set of EventBinding for convenience.
// usage: b := events.NewBindingGroup();
// 				defer b.Release();
//        b.Add(handle.Listen(...)) ...
type EventBindingGroup struct {
	Bindings []*EventBinding
}

func NewEventBindingGroup() *EventBindingGroup {
	return &EventBindingGroup{
		Bindings: make([]*EventBinding, 0),
	}
}

// Release releases and discards all held event bindings
func (group *EventBindingGroup) Release() {
	for i := range group.Bindings {
		group.Bindings[i].Release()
	}
	group.Bindings = make([]*EventBinding, 0)
}

// Add a binding
func (group *EventBindingGroup) Add(binding *EventBinding) {
	group.Bindings = append(group.Bindings, binding)
}

// Join two binding groups together
func (group *EventBindingGroup) Join(other *EventBindingGroup) {
	combo := make([]*EventBinding, len(group.Bindings) + len(other.Bindings))
	for i := 0; i < len(group.Bindings); i++ {
		combo[i] = group.Bindings[i]
	}

	offset := len(group.Bindings)
	for i := 0; i < len(other.Bindings); i++ {
		combo[offset + i] = other.Bindings[i]
	}

	group.Bindings = combo
	other.Bindings = combo
}

