package presenters

import (
	"github.com/openshift-online/addrewTrex/pkg/api"
	"github.com/openshift-online/addrewTrex/pkg/api/openapi"
	"github.com/openshift-online/addrewTrex/pkg/errors"
)

func ObjectKind(i interface{}) *string {
	result := ""
	switch i.(type) {
	case api.Dinosaur, *api.Dinosaur:
		result = "Dinosaur"
	case errors.ServiceError, *errors.ServiceError:
		result = "Error"
	}

	return openapi.PtrString(result)
}
