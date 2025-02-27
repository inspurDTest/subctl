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

package gather

import (
	submarinerv1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	gatewayPodLabel          = "app=submariner-gateway"
	routeagentPodLabel       = "app=submariner-routeagent"
	globalnetPodLabel        = "app=submariner-globalnet"
	metricsProxyPodLabel     = "app=submariner-metrics-proxy"
	addonPodLabel            = "app=submariner-addon"
	ovnMasterPodLabelOCP     = "app=ovnkube-master"
	ovnMasterPodLabelGeneric = "name=ovnkube-master"
	ovnKubePodLabel          = "app=ovnkube-node"
)

func gatherGatewayPodLogs(info *Info) {
	gatherPodLogs(gatewayPodLabel, info)
}

func gatherMetricsProxyPodLogs(info *Info) {
	gatherPodLogsByContainer(metricsProxyPodLabel, "gateway-metrics-proxy", info)

	if info.Submariner.Spec.GlobalCIDR != "" {
		gatherPodLogsByContainer(metricsProxyPodLabel, "globalnet-metrics-proxy", info)
	}
}

func gatherRouteAgentPodLogs(info *Info) {
	gatherPodLogs(routeagentPodLabel, info)
}

func gatherGlobalnetPodLogs(info *Info) {
	gatherPodLogs(globalnetPodLabel, info)
}

func gatherAddonPodLogs(info *Info) {
	gatherPodLogs(addonPodLabel, info)
}

func gatherEndpoints(info *Info, namespace string) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("endpoints"), namespace, v1.ListOptions{})
}

func gatherClusters(info *Info, namespace string) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("clusters"), namespace, v1.ListOptions{})
}

func gatherGateways(info *Info, namespace string) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("gateways"), namespace, v1.ListOptions{})
}

func gatherRouteAgents(info *Info, namespace string) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("routeagents"), namespace, v1.ListOptions{})
}

func gatherClusterGlobalEgressIPs(info *Info) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("clusterglobalegressips"), corev1.NamespaceAll, v1.ListOptions{})
}

func gatherGlobalEgressIPs(info *Info) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("globalegressips"), corev1.NamespaceAll, v1.ListOptions{})
}

func gatherGlobalIngressIPs(info *Info) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("globalingressips"), corev1.NamespaceAll, v1.ListOptions{})
}

func gatherGatewayRoutes(info *Info) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("gatewayroutes"), corev1.NamespaceAll, v1.ListOptions{})
}

func gatherNonGatewayRoutes(info *Info) {
	ResourcesToYAMLFile(info, submarinerv1.SchemeGroupVersion.WithResource("nongatewayroutes"), corev1.NamespaceAll, v1.ListOptions{})
}
