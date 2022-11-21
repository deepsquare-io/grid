package watcher

import (
	"context"

	"github.com/deepsquare-io/the-grid/ldap-connector/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/ldap-connector/ldap"
	"github.com/deepsquare-io/the-grid/ldap-connector/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"
)

type Watcher struct {
	ms   *metascheduler.MetaScheduler
	ldap *ldap.DataSource
}

func New(
	ms *metascheduler.MetaScheduler,
	ldap *ldap.DataSource,
) *Watcher {
	if ms == nil {
		logger.I.Panic("ms is nil")
	}
	if ldap == nil {
		logger.I.Panic("ldap is nil")
	}
	return &Watcher{
		ms:   ms,
		ldap: ldap,
	}
}

func (w *Watcher) Watch(parent context.Context) error {
	events := make(chan *metascheduler.MetaSchedulerClaimNextJobEvent)

	sub, err := w.ms.WatchClaimNextJobEvent(&bind.WatchOpts{
		Context: parent,
	}, events)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-parent.Done():
			logger.I.Warn("watch context is done")
			return nil
		case err := <-sub.Err():
			logger.I.Error("watch thrown an error", zap.Error(err))
			return err
		case event := <-events:
			if err := w.handleEvent(parent, event); err != nil {
				return err
			}
		}
	}
}

func (w *Watcher) handleEvent(parent context.Context, event *metascheduler.MetaSchedulerClaimNextJobEvent) error {
	user := event.CustomerAddr.Hex()
	if err := w.ldap.CreateUser(parent, user); err != nil {
		return err
	}

	if err := w.ldap.AddUserToGroup(parent, user); err != nil {
		return err
	}

	return nil
}
