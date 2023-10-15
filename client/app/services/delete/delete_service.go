package delete

import (
	"github.com/youtrolledhahaha/youtrolledhahahaXxxclient/app/services"
	"os"
)

type Service struct {
}

func NewService() services.Delete {
	return &Service{}
}

func (d Service) DeleteFile(filepath string) error {
	return os.Remove(filepath)
}
