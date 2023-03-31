package db

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/deepsquare-io/the-grid/grid-logger/server/crypto"
	"github.com/nxadm/tail"
	"go.uber.org/zap"
)

type File struct {
	storagePath string
	key         []byte
}

func NewFileDB(storagePath string, key []byte) *File {
	if err := os.MkdirAll(storagePath, 0o700); err != nil {
		logger.I.Panic("failed to mkdir storage path", zap.Error(err))
	}
	return &File{
		storagePath: storagePath,
		key:         key,
	}
}

func (db *File) Append(logName string, user string, content []byte) (n int, err error) {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", db.storagePath, user), 0o700); err != nil {
		logger.I.Error("failed to mkdir storage path", zap.Error(err))
	}
	logPath := fmt.Sprintf("%s/%s/%s.log", db.storagePath, user, logName)
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	encrypted, err := crypto.Encrypt(db.key, bytes.TrimRight(content, "\n\r"))
	if err != nil {
		return 0, err
	}
	str := base64.StdEncoding.EncodeToString(encrypted)
	n, err = file.WriteString(str + "\n")
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (db *File) ReadAndWatch(
	ctx context.Context,
	address string,
	logName string,
	out chan<- string,
) error {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", db.storagePath, address), 0o700); err != nil {
		logger.I.Error("failed to mkdir storage path", zap.Error(err))
	}
	logPath := fmt.Sprintf("%s/%s/%s.log", db.storagePath, address, logName)

	t, err := tail.TailFile(logPath, tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekStart,
		},
		Logger: tail.DiscardingLogger,
	})
	if err != nil {
		return err
	}
	defer func() {
		if err := t.Stop(); err != nil {
			logger.I.Error("tail failed to close with err", zap.Error(err))
		}
	}()
	for {
		select {
		case l, ok := <-t.Lines:
			if !ok {
				return nil
			}
			data, err := base64.StdEncoding.DecodeString(l.Text)
			if err != nil {
				return err
			}
			decrypted, err := crypto.Decrypt(db.key, data)
			if err != nil {
				return err
			}

			out <- string(decrypted)
		case <-ctx.Done():
			return nil
		}
	}
}

func (db *File) ListAndWatch(
	ctx context.Context,
	address string,
	out chan<- []string,
) error {
	ticker := time.NewTicker(10 * time.Second)

	logDir := fmt.Sprintf("%s/%s", db.storagePath, address)
	if err := os.MkdirAll(logDir, 0o700); err != nil {
		logger.I.Error("failed to mkdir storage path", zap.Error(err))
	}

	for {
		files, err := os.ReadDir(logDir)
		if err != nil {
			return err
		}
		names := make([]string, 0, len(files))
		for _, file := range files {
			if !file.IsDir() {
				names = append(names, strings.TrimRight(file.Name(), ".log"))
			}
		}
		out <- names

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
	}
}
