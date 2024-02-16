# Event Manager for Go - Simple and Flexible Event Handling

## Overview

This is a simple event manager module written in Go. It allows you to manage events and their handlers, providing functionalities to add, remove, and trigger events.

## Installation

To use this module in your Go project, you can install it using the following command:

```bash
go get -u github.com/NIR3X/eventmanager
```

## Usage

Here's a basic example demonstrating how to use the event manager:

```go
package main

import (
	"fmt"
	"github.com/NIR3X/eventmanager"
)

func main() {
	// Create a new EventManager instance.
	em := eventmanager.NewEventManager()

	// Add an event handler for "example_event".
	handlerId := em.AddHandler("example_event", func(data interface{}) {
		fmt.Printf("Received event: example_event, Data: %v\n", data)
	})

	// Trigger the "example_event" event.
	em.TriggerEvent("example_event", "Hello, world!")

	// Remove the event handler using the handler Id.
	em.RemoveHandler("example_event", handlerId)

	// Trigger the event again to verify removal.
	em.TriggerEvent("example_event", "Hello again!")
}
```

## License

[![GNU AGPLv3 Image](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.html)

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
