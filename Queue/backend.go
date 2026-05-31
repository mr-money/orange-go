package Queue

import (
	"github.com/RichardKnop/machinery/v1/backends/iface"
	"github.com/RichardKnop/machinery/v1/tasks"
)

// LocalTimeBackend wraps a Backend and converts CreatedAt from UTC to local timezone
type LocalTimeBackend struct {
	iface.Backend
}

// NewLocalTimeBackend creates a new LocalTimeBackend wrapper
func NewLocalTimeBackend(backend iface.Backend) *LocalTimeBackend {
	return &LocalTimeBackend{Backend: backend}
}

// GetState overrides to convert CreatedAt to local timezone
func (b *LocalTimeBackend) GetState(taskUUID string) (*tasks.TaskState, error) {
	state, err := b.Backend.GetState(taskUUID)
	if err != nil {
		return nil, err
	}
	if state != nil {
		state.CreatedAt = state.CreatedAt.Local()
	}
	return state, nil
}

// GroupTaskStates overrides to convert CreatedAt to local timezone
func (b *LocalTimeBackend) GroupTaskStates(groupUUID string, groupTaskCount int) ([]*tasks.TaskState, error) {
	states, err := b.Backend.GroupTaskStates(groupUUID, groupTaskCount)
	if err != nil {
		return nil, err
	}
	for _, state := range states {
		if state != nil {
			state.CreatedAt = state.CreatedAt.Local()
		}
	}
	return states, nil
}
