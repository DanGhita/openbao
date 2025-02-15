// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"context"

	"github.com/openbao/openbao/helper/namespace"
	"github.com/openbao/openbao/sdk/logical"
)

type entPolicyStore struct{}

func (ps *PolicyStore) extraInit() {
}

func (ps *PolicyStore) loadNamespacePolicies(context.Context, *Core) error { return nil }

func (ps *PolicyStore) getACLView(*namespace.Namespace) *BarrierView {
	return ps.aclView
}

func (ps *PolicyStore) getRGPView(ns *namespace.Namespace) *BarrierView {
	return ps.rgpView
}

func (ps *PolicyStore) getEGPView(ns *namespace.Namespace) *BarrierView {
	return ps.egpView
}

func (ps *PolicyStore) getBarrierView(ns *namespace.Namespace, _ PolicyType) *BarrierView {
	return ps.getACLView(ns)
}

func (ps *PolicyStore) handleSentinelPolicy(context.Context, *Policy, *BarrierView, *logical.StorageEntry) error {
	return nil
}

func (ps *PolicyStore) parseEGPPaths(*Policy) error { return nil }

func (ps *PolicyStore) invalidateEGPTreePath(string) {}

func (ps *PolicyStore) pathsToEGPPaths(*Policy) ([]*egpPath, error) { return nil, nil }

func (ps *PolicyStore) loadACLPolicyNamespaces(ctx context.Context, policyName, policyText string) error {
	return ps.loadACLPolicyInternal(namespace.RootContext(ctx), policyName, policyText)
}
