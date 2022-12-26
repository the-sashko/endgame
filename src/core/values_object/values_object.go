package values_object

import "endgame/src/core/interfaces"

type coreValuesObject struct {
	values map[string]interface{}
}

func (valuesObject *coreValuesObject) Get(name string) interface{} {
	if !valuesObject.Has(name) {
		return nil
	}

	return valuesObject.values[name]
}

func (valuesObject *coreValuesObject) Set(name string, value interface{}) {
	valuesObject.values[name] = value
}

func (valuesObject *coreValuesObject) Has(name string) bool {
	_, isExist := valuesObject.values[name]

	return isExist
}

func (valuesObject *coreValuesObject) Delete(name string) {
	if !valuesObject.Has(name) {
		return
	}

	delete(valuesObject.values, name)
}

func NewValuesObject(values map[string]interface{}) interfaces.ICoreValuesObject {
	return &coreValuesObject{
		values: values,
	}
}
