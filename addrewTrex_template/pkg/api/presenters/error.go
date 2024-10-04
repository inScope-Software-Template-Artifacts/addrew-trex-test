package presenters

import (
	"github.com/openshift-online/addrewTrex/pkg/api/openapi"
	"github.com/openshift-online/addrewTrex/pkg/errors"
)

func PresentError(err *errors.ServiceError) openapi.Error {
	return err.AsOpenapiError("")
}
