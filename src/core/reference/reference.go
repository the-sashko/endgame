package reference

import (
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
	"github.com/google/uuid"
)

type coreReference struct {
	value string
}

func (reference *coreReference) Get() string {
	return reference.value
}

func NewReference() interfaces.ICoreReference {
	uuidObject, err := uuid.NewUUID()

	if err != nil {
		doError(err)
	}

	return &coreReference{value: uuidObject.String()}
}

func doError(errorEntity error) {
	logger.LogError("core_reference", errorEntity, true)
}
