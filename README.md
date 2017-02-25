# Events

Events is a simple event handling system.

# Usage

    npm install shadowmint/go-events --save

Then use the event handler to receive and trigger events as callbacks:

    handler := events.New()

    handler.Listen(EventType{}, func(ep interface{}) {
        e := ep.(ConcreteEventType)
        ...
    })

    handler.Trigger(ConcreteEventType{})

Event handlers are bound until something clears them, but you can also
supply 'one time' event handlers using flags on Listen():

    handler := events.New()

    handler.Listen(EventType{}, func(ep interface{}) {
        e := ep.(ConcreteEventType)
        ...
    }, events.RunOnce)

    handler.Trigger(ConcreteEventType{})
