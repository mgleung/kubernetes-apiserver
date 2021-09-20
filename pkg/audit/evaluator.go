/*
Copyright 2021 The Kubernetes Authors.

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

package audit

import (
	"k8s.io/apiserver/pkg/apis/audit"
	"k8s.io/apiserver/pkg/authorization/authorizer"
)

// AuditContext is a pair of the audit configuration object that applies to
// a given request and the audit Event object that is being captured.
// It's a convenient placeholder to store both these objects in the request context.
type AuditContext struct {
	// RequestAuditConfig is the audit configuration that applies to the request
	RequestAuditConfig RequestAuditConfig

	// Event is the audit Event object that is being captured to be written in
	// the API audit log. It is set to nil when the request is not being audited.
	Event *audit.Event
}

// RequestAuditConfig is the evaluated audit configuration that is applicable to
// a given request. PolicyRuleEvaluator evaluates the audit policy against the
// authorizer attributes and returns a RequestAuditConfig that applies to the request.
type RequestAuditConfig struct{}

// PolicyRuleEvaluator exposes methods for evaluating the policy rules.
type PolicyRuleEvaluator interface {
	// Check the audit level for a request with the given authorizer attributes.
	LevelAndStages(authorizer.Attributes) (audit.Level, []audit.Stage)
}
