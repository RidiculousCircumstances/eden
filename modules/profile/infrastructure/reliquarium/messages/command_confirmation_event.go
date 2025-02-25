package messages

type ServiceName string

const (
	Eden ServiceName = "service_name"
)

type CommandType string

const (
	Pause         CommandType = "pause"
	TakeSnapshots CommandType = "takeSnapshots"
	Resume        CommandType = "resume"
)

type CommandConfirmationEvent struct {
	ServiceName      ServiceName          `json:"service_name"`
	Command          CommandType          `json:"command"`
	TakeSnapshotData *TakeSnapshotPayload `json:"takeSnapshotData,omitempty"`
}
type TakeSnapshotPayload struct {
	SnapshotStorageKey string `json:"snapshotStorageKey,omitempty"`
}

func NewCommandConfirmationEvent(serviceName ServiceName, command CommandType, snapshotStorageKey string) *CommandConfirmationEvent {
	event := &CommandConfirmationEvent{
		ServiceName: serviceName,
		Command:     command,
	}

	if command == TakeSnapshots {
		event.TakeSnapshotData = &TakeSnapshotPayload{SnapshotStorageKey: snapshotStorageKey}
	}

	return event
}
