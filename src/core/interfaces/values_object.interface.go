package interfaces

type ICoreValuesObject interface {
	Get(name string) interface{}
	Set(name string, value interface{})
	Has(name string) bool
	Delete(name string)
}
