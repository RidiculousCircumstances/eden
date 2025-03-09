package usecase

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	usecaseIntf "eden/modules/profile/application/usecase/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/modules/profile/infrastructure/reliquarium/messages"
	loggerIntf "eden/shared/logger/interfaces"
	"errors"
)

var ErrUnknownCommand = errors.New("unknown command")

type ManageSnapshotLifecycle struct {
	stateManager          interfaces.AppStateManager
	confirmationPublisher interfaces.ServiceCommandConfirmationPublisher
	takeSnapshot          usecaseIntf.TakeSnapshot
	logger                loggerIntf.Logger
}

func NewManageSnapshotLifecycle(
	stateManager interfaces.AppStateManager,
	confirmationPublisher interfaces.ServiceCommandConfirmationPublisher,
	takeSnapshot usecaseIntf.TakeSnapshot,
	logger loggerIntf.Logger,
) *ManageSnapshotLifecycle {
	return &ManageSnapshotLifecycle{
		stateManager:          stateManager,
		confirmationPublisher: confirmationPublisher,
		takeSnapshot:          takeSnapshot,
		logger:                logger,
	}
}

func (u *ManageSnapshotLifecycle) Process(ctx context.Context, msg *message.ServiceControlCommand) error {
	switch msg.Command {
	case message.Pause:
		u.stateManager.Pause()
		u.logger.Info("[ManageSnapshotLifecycle] pause")
		return u.confirmationPublisher.Publish(ctx, messages.NewCommandConfirmationEvent(messages.Eden, messages.Pause, ""))
	case message.TakeSnapshots:
		u.logger.Info("[ManageSnapshotLifecycle] taking snapshot")
		dumpStorageKey, err := u.takeSnapshot.Process(ctx)
		if err != nil {
			return err
		}
		err = u.confirmationPublisher.Publish(ctx, messages.NewCommandConfirmationEvent(messages.Eden, messages.TakeSnapshots, dumpStorageKey))
		if err != nil {
			return err
		}
		return nil
	case message.Resume:
		u.stateManager.Resume()
		u.logger.Info("[ManageSnapshotLifecycle] resume")
		return nil
	default:
		return ErrUnknownCommand
	}
}
