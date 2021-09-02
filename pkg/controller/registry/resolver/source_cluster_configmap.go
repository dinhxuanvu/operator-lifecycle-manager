package resolver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry/resolver/cache"
	v1 "k8s.io/client-go/listers/core/v1"
)

type globalConfigMapSource struct {
	lister    v1.ConfigMapLister
	name      string
	namespace string
}

func (s *globalConfigMapSource) Snapshot(ctx context.Context) (*cache.Snapshot, error) {
	cm, err := s.lister.ConfigMaps(s.namespace).Get(s.name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch configmap %q: %v", s.name, err)
	}

	entry := &cache.Operator{
		Name: "cluster",
		SourceInfo: &cache.OperatorSourceInfo{
			Catalog: cache.SourceKey{
				Name:      "@cluster",
				Namespace: s.namespace,
			},
		},
	}
	if err := json.Unmarshal([]byte(cm.Data["properties"]), &entry.Properties); err != nil {
		return nil, fmt.Errorf("failed to unmarshal properties from configmap %q: %v", s.name, err)
	}

	fmt.Printf("@@@ loaded cm with properties: %#v\n", entry.Properties)

	return &cache.Snapshot{
		Entries: []*cache.Operator{entry},
	}, nil
}
