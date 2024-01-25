package eventmanager

import (
	"sync"
)

// EventHandler represents a function that handles events.
type EventHandler func(interface{})

// EventManager manages events and their handlers.
type EventManager struct {
	mutex    sync.RWMutex
	handlers map[string]map[int64]EventHandler
	nextID   int64
}

// NewEventManager creates a new instance of EventManager.
func NewEventManager() *EventManager {
	return &EventManager{
		handlers: make(map[string]map[int64]EventHandler),
		nextID:   1,
	}
}

// AddHandler adds an event handler for the specified event and returns the handler ID.
func (em *EventManager) AddHandler(event string, handler EventHandler) int64 {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	handlerID := em.nextID
	em.nextID++

	if em.handlers[event] == nil {
		em.handlers[event] = make(map[int64]EventHandler)
	}
	em.handlers[event][handlerID] = handler

	return handlerID
}

// RemoveHandler removes an event handler for the specified event and handler ID.
func (em *EventManager) RemoveHandler(event string, handlerID int64) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	handlers, ok := em.handlers[event]
	if !ok {
		return
	}

	delete(handlers, handlerID)

	// Remove the map entry if there are no more handlers for the event
	if len(handlers) == 0 {
		delete(em.handlers, event)
	}
}

// TriggerEvent triggers the specified event with the given data.
func (em *EventManager) TriggerEvent(event string, data interface{}) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	handlers, ok := em.handlers[event]
	if !ok {
		return
	}

	// Execute each handler
	for _, handler := range handlers {
		handler(data)
	}
}
