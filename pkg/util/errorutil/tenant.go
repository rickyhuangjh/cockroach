// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package errorutil

import "github.com/cockroachdb/cockroach/pkg/util/errorutil/unimplemented"

// UnsupportedUnderClusterVirtualizationMessage is the message used by UnsupportedUnderClusterVirtualization error.
const UnsupportedUnderClusterVirtualizationMessage = "operation is unsupported within a virtual cluster"

// UnsupportedUnderClusterVirtualization returns an error suitable for
// returning when an operation could not be carried out due to the SQL
// server running inside a virtual cluster. In that mode, Gossip and
// other components of the KV layer are not available.
func UnsupportedUnderClusterVirtualization(issue int) error {
	return unimplemented.NewWithIssue(issue, UnsupportedUnderClusterVirtualizationMessage)
}

// FeatureNotAvailableToNonSystemTenantsIssue is to be used with the
// Optional and related error interfaces when a feature is simply not
// available to non-system tenants (i.e. we're not planning to change
// this).
// For all other multitenancy errors where there is a plan to
// improve the situation, a specific issue should be created instead.
const FeatureNotAvailableToNonSystemTenantsIssue = 54252
