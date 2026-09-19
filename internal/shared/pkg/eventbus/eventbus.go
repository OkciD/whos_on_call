package eventbus

import "sync"

type EventType any
type EventPayload any

type Event struct {
	Type    EventType
	Payload EventPayload
}

type EventChan chan Event

type EventBus interface {
	Emit(event Event)
	On(eventType EventType) EventChan
	Off(eventType EventType, ch EventChan)
}

type eventBus struct {
	mu          sync.RWMutex
	subscribers map[EventType][]EventChan
}

func NewEventBus() EventBus {
	return &eventBus{
		subscribers: make(map[EventType][]EventChan),
	}
}

func (eb *eventBus) On(eventType EventType) EventChan {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	ch := make(EventChan, 1)
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)

	return ch
}

func (eb *eventBus) Emit(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	chans := append([]EventChan{}, eb.subscribers[event.Type]...)
	go func() {
		for _, ch := range chans {
			select {
			case ch <- event:
			default:
			}
		}
	}()
}

func (eb *eventBus) Off(eventType EventType, ch EventChan) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	if subscribers, ok := eb.subscribers[eventType]; ok {
		for i, subscriber := range subscribers {
			if ch == subscriber {
				eb.subscribers[eventType] = append(subscribers[:i], subscribers[i+1:]...)
				return
			}
		}
	}
}
