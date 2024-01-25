package eventmanager

import (
	"testing"
)

// TestEventManager is a test function for the EventManager.
func TestEventManager(t *testing.T) {
	// Create a new EventManager instance.
	eventManager := NewEventManager()

	// Initialize a counter to be used for testing.
	counter := 0

	// Add a new event handler for the "example_event" event.
	handlerID := eventManager.AddHandler("example_event", func(data interface{}) {
		// Increment the counter when the event is triggered.
		*data.(*int)++
	})

	// Trigger the "example_event" event with the initialized counter.
	eventManager.TriggerEvent("example_event", &counter)

	// Check if the counter has been incremented as expected.
	if counter != 1 {
		t.Errorf("Expected counter to be 1, got %d", counter)
	}

	// Remove the previously added event handler.
	eventManager.RemoveHandler("example_event", handlerID)

	// Trigger the "example_event" event again after removing the handler.
	eventManager.TriggerEvent("example_event", &counter)

	// Check if the counter remains unchanged after removing the handler.
	if counter != 1 {
		t.Errorf("Expected counter to be 1, got %d", counter)
	}
}
