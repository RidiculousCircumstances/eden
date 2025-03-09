package storage

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
)

var (
	ErrWrongObjectOptions = errors.New("wrong object options")
	ErrWrongPutOptions    = errors.New("wrong put options")
)

type MinioClient struct {
	client *minio.Client
}

func NewMinioClient(client *minio.Client) (*MinioClient, error) {
	return &MinioClient{client: client}, nil
}

func (mc *MinioClient) GetObject(ctx context.Context, bucketName, objectName string, opts interface{}) (*minio.Object, error) {
	getOptions, ok := opts.(minio.GetObjectOptions)
	if !ok {
		return nil, ErrWrongObjectOptions
	}
	return mc.client.GetObject(ctx, bucketName, objectName, getOptions)
}

// DeleteObject удаляет объект (файл) из хранилища
func (mc *MinioClient) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	err := mc.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

// приватный метод для создания бакета, если он не существует
func (mc *MinioClient) createBucketIfNotExists(ctx context.Context, bucketName string) error {
	// Проверяем наличие бакета
	exists, err := mc.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	// Если бакет не существует, создаем его
	if !exists {
		err = mc.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Printf("Failed to create bucket: %v", err)
			return err
		}
		log.Printf("Bucket created successfully: %s", bucketName)
	}

	return nil
}

// UploadObject загружает объект в хранилище. Если бакет не существует, он будет создан.
func (mc *MinioClient) UploadObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts interface{}) error {
	//Вызываем приватный метод для создания бакета, если нужно
	err := mc.createBucketIfNotExists(ctx, bucketName)
	if err != nil {
		return err
	}

	// Проверка и загрузка объекта
	putOptions, ok := opts.(minio.PutObjectOptions)
	if !ok {
		return ErrWrongPutOptions
	}

	_, err = mc.client.PutObject(ctx, bucketName, objectName, reader, objectSize, putOptions)
	if err != nil {
		return err
	}
	return nil
}
