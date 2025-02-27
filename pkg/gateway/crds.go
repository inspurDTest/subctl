/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package gateway

import (
	"context"

	"github.com/pkg/errors"
	"github.com/submariner-io/submariner-operator/pkg/crd"
	"github.com/submariner-io/submariner-operator/pkg/embeddedyamls"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// Ensure ensures that the required resources are deployed on the target system.
// The resources handled here are the gateway CRDs: Cluster and Endpoint.
func Ensure(ctx context.Context, crdUpdater crd.Updater) error {
	for _, ref := range []struct {
		name string
		crd  string
	}{
		{embeddedyamls.Deploy_submariner_crds_submariner_io_clusters_yaml, "Cluster"},
		{embeddedyamls.Deploy_submariner_crds_submariner_io_endpoints_yaml, "Endpoint"},
		{embeddedyamls.Deploy_submariner_crds_submariner_io_gateways_yaml, "Gateway"},
		{embeddedyamls.Deploy_submariner_crds_submariner_io_clusterglobalegressips_yaml, "ClusterGlobalEgressIP"},
		{embeddedyamls.Deploy_submariner_crds_submariner_io_globalegressips_yaml, "GlobalEgressIP"},
		{embeddedyamls.Deploy_submariner_crds_submariner_io_globalingressips_yaml, "GlobalIngressIP"},
	} {
		_, err := crdUpdater.CreateOrUpdateFromEmbedded(ctx, ref.name)
		if err != nil && !apierrors.IsAlreadyExists(err) {
			return errors.Wrapf(err, "error provisioning the %s CRD", ref.crd)
		}
	}

	return nil
}
