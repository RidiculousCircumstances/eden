package serializer

import (
	"encoding/json"
	"errors"
	"reflect"
)

// Ошибки, связанные с сериализацией и десериализацией
var (
	ErrUnsupportedType = errors.New("unsupported type for deserialization")
)

// JSONSerializer - конкретная реализация сериализатора для JSON
type JSONSerializer struct {
	types []reflect.Type
}

// NewJSONSerializer создает новый экземпляр JSONSerializer с поддерживаемыми типами
func NewJSONSerializer(supportedTypes ...interface{}) *JSONSerializer {
	types := make([]reflect.Type, len(supportedTypes))
	for i, t := range supportedTypes {
		types[i] = reflect.TypeOf(t)
	}
	return &JSONSerializer{types: types}
}

// Serialize сериализует данные в JSON
func (s *JSONSerializer) Serialize(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// Deserialize десериализует данные из JSON в один из поддерживаемых типов
func (s *JSONSerializer) Deserialize(data []byte, v interface{}) error {
	var errList []error

	for _, t := range s.types {
		// Создаем экземпляр типа
		value := reflect.New(t).Interface()

		// Пытаемся десериализовать
		if err := json.Unmarshal(data, value); err == nil {
			// Успех: записываем результат в v
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(value).Elem())
			return nil
		} else {
			errList = append(errList, err)
		}
	}

	// Если ни один из типов не подошел, возвращаем все ошибки
	return errors.Join(errList...)
}
