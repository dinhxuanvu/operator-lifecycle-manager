package resolver

import (
	"github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry/resolver/cache"
)

type mergedSourceProvider []cache.SourceProvider

func (sps mergedSourceProvider) Sources(namespaces ...string) map[cache.SourceKey]cache.Source {
	sources := make(map[cache.SourceKey]cache.Source)

	for _, sp := range sps {
		for key, source := range sp.Sources(namespaces...) {
			sources[key] = source
		}
	}

	return sources
}
