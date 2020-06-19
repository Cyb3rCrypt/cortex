/*
Copyright 2020 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	"fmt"

	"github.com/cortexlabs/cortex/pkg/lib/errors"
	"github.com/cortexlabs/cortex/pkg/types/userconfig"
)

const (
	ErrKindNotSupported              = "resources.kind_not_supported"
	ErrAPIUpdating                   = "resources.api_updating"
	ErrAPINotDeployed                = "resources.api_not_deployed"
	ErrCannotChangeTypeOfDeployedAPI = "resources.cannot_change_kind_of_deployed_api"
)

func ErrorKindNotSupported(kind userconfig.Kind) error {
	return errors.WithStack(&errors.Error{
		Kind:    ErrKindNotSupported,
		Message: fmt.Sprintf("%s kind is not supported", kind.String()),
	})
}

func ErrorAPIUpdating(apiName string) error {
	return errors.WithStack(&errors.Error{
		Kind:    ErrAPIUpdating,
		Message: fmt.Sprintf("%s is updating (override with --force)", apiName),
	})
}

func ErrorAPINotDeployed(apiName string) error {
	return errors.WithStack(&errors.Error{
		Kind:    ErrAPINotDeployed,
		Message: fmt.Sprintf("%s is not deployed", apiName), // note: if modifying this string, search the codebase for it and change all occurrences
	})
}

func ErrorCannotChangeKindOfDeployedAPI(name string, newKind, prevKind userconfig.Kind) error {
	return errors.WithStack(&errors.Error{
		Kind:    ErrCannotChangeTypeOfDeployedAPI,
		Message: fmt.Sprintf("cannot change the kind of %s to %s because it has already been deployed with kind %s; please delete it with `cortex delete %s` and redeploy to change the kind", name, newKind.String(), prevKind.String(), name),
	})
}
