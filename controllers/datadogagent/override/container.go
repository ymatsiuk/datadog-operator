// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package override

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/DataDog/datadog-operator/apis/datadoghq/v2alpha1"
	"github.com/DataDog/datadog-operator/controllers/datadogagent/feature"
)

// Container use to override a corev1.Container with a 2alpha1.DatadogAgentGenericContainer.
func Container(manager feature.PodTemplateManagers, override *v2alpha1.DatadogAgentGenericContainer) (*corev1.Container, error) {
	// TODO(operator-ga): implement OverrideContainer
	return nil, nil
}
