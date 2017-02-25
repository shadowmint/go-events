package events_test

import (
	"testing"

	"ntoolkit/assert"
	"ntoolkit/events"
)

type EventType struct {
}

func TestNew(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		handler := events.New()
		T.Assert(handler != nil)
	})
}

func TestTrigger(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		handler := events.New()

		resolved := 0
		handler.Listen(EventType{}, func(ep interface{}) {
			resolved += 1
		})

		T.Assert(resolved == 0)
		for i := 0; i < 10; i++ {
			handler.Trigger(EventType{})
			T.Assert(resolved == i + 1)
		}
	})
}

func TestTriggerOnce(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		handler := events.New()

		resolved := 0
		handler.Listen(EventType{}, func(ep interface{}) {
			resolved += 1
		}, events.RunOnce)

		T.Assert(resolved == 0)
		for i := 0; i < 10; i++ {
			handler.Trigger(EventType{})
		}

		T.Assert(resolved == 1)
	})
}

func TestReleaseBinding(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		handler := events.New()

		resolved := 0
		var binding *events.EventBinding
		binding = handler.Listen(EventType{}, func(ep interface{}) {
			resolved += 1
			if resolved == 2 {
				handler.Release(binding)
			}
		})

		handler.Trigger(EventType{})
		handler.Trigger(EventType{})

		T.Assert(resolved == 2)

		handler.Trigger(EventType{})
		handler.Trigger(EventType{})

		T.Assert(resolved == 2)
	})
}