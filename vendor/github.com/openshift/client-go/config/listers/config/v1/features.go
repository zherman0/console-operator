// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FeaturesLister helps list Features.
type FeaturesLister interface {
	// List lists all Features in the indexer.
	List(selector labels.Selector) (ret []*v1.Features, err error)
	// Get retrieves the Features from the index for a given name.
	Get(name string) (*v1.Features, error)
	FeaturesListerExpansion
}

// featuresLister implements the FeaturesLister interface.
type featuresLister struct {
	indexer cache.Indexer
}

// NewFeaturesLister returns a new FeaturesLister.
func NewFeaturesLister(indexer cache.Indexer) FeaturesLister {
	return &featuresLister{indexer: indexer}
}

// List lists all Features in the indexer.
func (s *featuresLister) List(selector labels.Selector) (ret []*v1.Features, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Features))
	})
	return ret, err
}

// Get retrieves the Features from the index for a given name.
func (s *featuresLister) Get(name string) (*v1.Features, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("features"), name)
	}
	return obj.(*v1.Features), nil
}