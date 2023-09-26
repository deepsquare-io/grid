package watcher

import (
	"context"
	"errors"
	"strings"

	"github.com/deepsquare-io/grid/ldap-connector/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/grid/ldap-connector/ldap"
	"github.com/deepsquare-io/grid/ldap-connector/logger"
	"github.com/deepsquare-io/grid/ldap-connector/validate"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"
)

type Watcher struct {
	metascheduler *metascheduler.MetaScheduler
	ldap          *ldap.DataSource
}

func New(
	metascheduler *metascheduler.MetaScheduler,
	ldap *ldap.DataSource,
) *Watcher {
	if metascheduler == nil {
		logger.I.Panic("metascheduler is nil")
	}
	if ldap == nil {
		logger.I.Panic("ldap is nil")
	}
	return &Watcher{
		metascheduler: metascheduler,
		ldap:          ldap,
	}
}

func (w *Watcher) Watch(parent context.Context) error {
	events := make(chan *metascheduler.MetaSchedulerNewJobRequestEvent)

	sub, err := w.metascheduler.WatchNewJobRequestEvent(&bind.WatchOpts{
		Context: parent,
	}, events)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	logger.I.Info("Watching events...")

	for {
		select {
		case <-parent.Done():
			logger.I.Warn("watch context is done")
			return nil
		case err := <-sub.Err():
			logger.I.Error("watch thrown an error", zap.Error(err))
			return err
		case event := <-events:
			logger.I.Info("Received event", zap.Any("event", event))
			if err := w.handleEvent(parent, event); err != nil {
				return err
			}
		}
	}
}

func (w *Watcher) handleEvent(
	parent context.Context,
	event *metascheduler.MetaSchedulerNewJobRequestEvent,
) error {
	user := strings.ToLower(event.CustomerAddr.Hex())
	if errMsg := validate.LDAPUserIsValid(user); errMsg != "" {
		logger.I.Error("user is invalid", zap.Error(errors.New(errMsg)))
		return nil
	}

	if err := w.ldap.CreateUser(parent, user); err != nil {
		return err
	}
	logger.I.Info("Created user", zap.String("user", user))

	if err := w.ldap.AddUserToGroup(parent, user); err != nil {
		return err
	}
	logger.I.Info("Added user to group", zap.String("user", user))

	return nil
}
