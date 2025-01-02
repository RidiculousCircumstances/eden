package logger

import (
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

type ZapLoggerAdapter struct {
	logger loggerIntf.Logger
	fields []zap.Field
}

func NewZapLoggerAdapter(baseLogger loggerIntf.Logger) *ZapLoggerAdapter {
	return &ZapLoggerAdapter{
		logger: baseLogger,
		fields: []zap.Field{},
	}
}

func (z *ZapLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	allFields := append(z.fields, zap.Error(err))
	allFields = append(allFields, mapToZapFields(fields)...)
	z.logger.Error(msg, allFields...)
}

func (z *ZapLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	allFields := append(z.fields, mapToZapFields(fields)...)
	z.logger.Info(msg, allFields...)
}

func (z *ZapLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	allFields := append(z.fields, mapToZapFields(fields)...)
	z.logger.Debug(msg, allFields...)
}

func (z *ZapLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	allFields := append(z.fields, mapToZapFields(fields)...)
	z.logger.Debug(msg, append(allFields, zap.String("level", "trace"))...)
}

func (z *ZapLoggerAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	newFields := append(z.fields, mapToZapFields(fields)...)
	return &ZapLoggerAdapter{
		logger: z.logger.With(newFields...),
		fields: newFields,
	}
}

// mapToZapFields конвертирует LogFields в []zap.Field
func mapToZapFields(fields watermill.LogFields) []zap.Field {
	if fields == nil {
		return nil
	}
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return zapFields
}
