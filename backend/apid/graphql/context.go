package graphql

import (
	"context"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/backend/apid/graphql/globalid"
)

var clusterKey key

// ContextWithCluster returns new context with given cluster name added to the
// context.
func ContextWithCluster(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, clusterKey, name)
}

// ClusterFromContext return cluster name stored in the given context; returns
// empty string if missing.
func ClusterFromContext(ctx context.Context) string {
	if val := ctx.Value(clusterKey); val != nil {
		return val.(string)
	}
	return ""
}

// setContextFromComponents takes a context and global id components, adds the
// namespace to the context, and returns the updated context.
func setContextFromComponents(ctx context.Context, c globalid.Components) context.Context {
	return contextWithNamespace(ctx, c.Namespace())
}

// contextWithNamespace takes a context and a name adds it to the context, and
// returns the updated context.
func contextWithNamespace(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, corev2.NamespaceKey, name)
}
