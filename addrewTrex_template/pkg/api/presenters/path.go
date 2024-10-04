package presenters

import (
	"fmt"

	"github.com/openshift-online/addrewTrex/pkg/api/openapi"

	"github.com/openshift-online/addrewTrex/pkg/api"
	"github.com/openshift-online/addrewTrex/pkg/errors"
)

const (
	BasePath = "/api/addrewTrex/v1"
)

func ObjectPath(id string, obj interface{}) *string {
	return openapi.PtrString(fmt.Sprintf("%s/%s/%s", BasePath, path(obj), id))
}

func path(i interface{}) string {
	switch i.(type) {
	case api.Dinosaur, *api.Dinosaur:
		return "dinosaurs"
	case errors.ServiceError, *errors.ServiceError:
		return "errors"
	default:
		return ""
	}
}
