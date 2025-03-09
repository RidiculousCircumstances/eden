package usecase

import (
	"context"
	"eden/modules/profile/application/usecase/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/minio/minio-go/v7"
	"io"
	"os/exec"
	"time"

	"go.uber.org/zap"
)

type TakeSnapshotConfig struct {
	SnapshotBucket string
	DbUser         string
	DbPassword     string
	DbName         string
	DbHost         string
}

type TakeSnapshot struct {
	logger         loggerIntf.Logger
	storageService interfaces.StorageService
	cfg            *TakeSnapshotConfig
}

func NewTakeSnapshot(logger loggerIntf.Logger, storageService interfaces.StorageService, cfg *TakeSnapshotConfig) *TakeSnapshot {
	return &TakeSnapshot{
		logger:         logger,
		storageService: storageService,
		cfg:            cfg,
	}
}

func (t *TakeSnapshot) Process(ctx context.Context) (string, error) {
	objectName := "eden_" + time.Now().Format("20060102_150405") + ".sql.gz"

	r, w := io.Pipe()

	// Горутина: запускает mysqldump -> gzip -> w
	go func() {
		defer func(w *io.PipeWriter) {
			err := w.Close()
			if err != nil {
				t.logger.Error("error closing pipe", zap.Error(err))
			}
		}(w)

		// 1. Запускаем mysqldump
		dumpCmd := exec.CommandContext(ctx,
			"mysqldump",
			"--host="+t.cfg.DbHost, // Указываем имя контейнера с MySQL
			"--protocol=TCP",       // Используем TCP вместо Unix-сокета
			"-u", t.cfg.DbUser,
			"-p"+t.cfg.DbPassword,
			t.cfg.DbName,
		)
		// Подключаем stdout mysqldump к stdin gzip
		dumpOut, err := dumpCmd.StdoutPipe()
		if err != nil {
			t.logger.Error("Failed to create dump stdout pipe", zap.Error(err))
			_ = w.CloseWithError(err)
			return
		}

		gzipCmd := exec.CommandContext(ctx, "gzip", "-c")
		gzipCmd.Stdin = dumpOut
		gzipCmd.Stdout = w

		// Стартуем mysqldump
		if err := dumpCmd.Start(); err != nil {
			t.logger.Error("Failed to start mysqldump", zap.Error(err))
			_ = w.CloseWithError(err)
			return
		}

		// Стартуем gzip
		if err := gzipCmd.Start(); err != nil {
			t.logger.Error("Failed to start gzip", zap.Error(err))
			_ = w.CloseWithError(err)
			return
		}

		// Ждем окончания mysqldump
		if err := dumpCmd.Wait(); err != nil {
			t.logger.Error("mysqldump exited with error", zap.Error(err))
			_ = w.CloseWithError(err)
			return
		}

		// Ждем окончания gzip
		if err := gzipCmd.Wait(); err != nil {
			t.logger.Error("gzip exited with error", zap.Error(err))
			_ = w.CloseWithError(err)
			return
		}

		t.logger.Info("mysqldump + gzip finished successfully")
	}()

	// Основная горутина: читает r и загружает в MinIO
	// При передаче -1 minio-go работает в "streaming" режиме (без знания размера).
	// Для очень больших файлов возможно лучше вручную делать multi-part upload.
	err := t.storageService.UploadObject(ctx, t.cfg.SnapshotBucket, objectName, -1, r, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		// Если произошла ошибка, попросим io.Pipe() закрыться
		_ = w.CloseWithError(err)
		t.logger.Error("Failed to upload to MinIO", zap.Error(err))
		return "", err
	}

	t.logger.Info("Snapshot streamed and uploaded successfully",
		zap.String("object", objectName))

	// Возвращаем имя заархивированного дампа
	return objectName, nil
}
