/*
Copyright 2019 Red Hat Inc.

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

package controller

import (
	"github.com/konveyor/virt-controller/pkg/controller/host"
	"github.com/konveyor/virt-controller/pkg/controller/map"
	"github.com/konveyor/virt-controller/pkg/controller/migration"
	"github.com/konveyor/virt-controller/pkg/controller/plan"
	"github.com/konveyor/virt-controller/pkg/controller/provider"
	"github.com/konveyor/virt-controller/pkg/settings"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

//
// Function provided by controller packages to add
// them self to the manager.
type AddFunction func(manager.Manager) error

//
// List of main controllers
var MainControllers = []AddFunction{
	migration.Add,
	plan.Add,
	mp.Add,
	host.Add,
}

//
// List of Inventory controllers
var InventoryControllers = []AddFunction{
	provider.Add,
}

//
// Add controllers to the manager based on role.
func AddToManager(m manager.Manager) error {
	err := settings.Settings.Load()
	if err != nil {
		return err
	}
	load := func(functions []AddFunction) error {
		for _, f := range functions {
			if err := f(m); err != nil {
				return err
			}
		}
		return nil
	}
	if settings.Settings.Role.Has(settings.InventoryRole) {
		err := load(InventoryControllers)
		if err != nil {
			return err
		}

	}
	if settings.Settings.Role.Has(settings.MainRole) {
		err := load(MainControllers)
		if err != nil {
			return err
		}

	}

	return nil
}
