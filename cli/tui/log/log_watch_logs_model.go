package log

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"time"
	"unicode"

	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"go.uber.org/zap"
)

type logMsg struct {
	timestamp time.Time
	message   string
}

func makeWatchLogsModel(
	ctx context.Context,
	jobID [32]byte,
	loggerDialer logger.Dialer,
) channel.Model[logMsg] {
	return channel.Model[logMsg]{
		Channel: make(chan logMsg, 100),
		OnInit: func(c chan logMsg) func() error {
			l, conn, err := loggerDialer.DialContext(ctx)
			if err != nil {
				log.I.Error("failed to get logs", zap.Error(err))
				return nil
			}
			stream, err := l.WatchLogs(ctx, jobID)
			if err != nil {
				log.I.Error("failed to get logs", zap.Error(err))
				return nil
			}

			go func() {
				defer conn.Close()
				defer func() {
					_ = stream.CloseSend()
				}()
				for {
					req, err := stream.Recv()
					if err == io.EOF || errors.Is(err, context.Canceled) {
						c <- logMsg{
							timestamp: time.Now(),
							message:   "---Connection to logging server closed---",
						}
						return
					}
					if err != nil {
						c <- logMsg{
							timestamp: time.Now(),
							message:   fmt.Sprintf("---Connection to logging server closed unexpectedly---\n%s", err),
						}
						return
					}
					clean := bytes.Map(func(r rune) rune {
						if unicode.IsPrint(r) {
							return r
						}
						return -1
					}, req.GetData())
					select {
					case c <- logMsg{
						timestamp: time.Unix(0, req.GetTimestamp()),
						message:   string(clean),
					}:
					case <-ctx.Done():
						return
					}
				}
			}()

			return func() error {
				_ = conn.Close()
				_ = stream.CloseSend()
				return nil
			}
		},
	}
}
