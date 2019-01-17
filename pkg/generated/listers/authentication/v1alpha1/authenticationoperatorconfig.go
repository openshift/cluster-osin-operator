// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/cluster-authentication-operator/pkg/apis/authentication/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthenticationOperatorConfigLister helps list AuthenticationOperatorConfigs.
type AuthenticationOperatorConfigLister interface {
	// List lists all AuthenticationOperatorConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AuthenticationOperatorConfig, err error)
	// Get retrieves the AuthenticationOperatorConfig from the index for a given name.
	Get(name string) (*v1alpha1.AuthenticationOperatorConfig, error)
	AuthenticationOperatorConfigListerExpansion
}

// authenticationOperatorConfigLister implements the AuthenticationOperatorConfigLister interface.
type authenticationOperatorConfigLister struct {
	indexer cache.Indexer
}

// NewAuthenticationOperatorConfigLister returns a new AuthenticationOperatorConfigLister.
func NewAuthenticationOperatorConfigLister(indexer cache.Indexer) AuthenticationOperatorConfigLister {
	return &authenticationOperatorConfigLister{indexer: indexer}
}

// List lists all AuthenticationOperatorConfigs in the indexer.
func (s *authenticationOperatorConfigLister) List(selector labels.Selector) (ret []*v1alpha1.AuthenticationOperatorConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthenticationOperatorConfig))
	})
	return ret, err
}

// Get retrieves the AuthenticationOperatorConfig from the index for a given name.
func (s *authenticationOperatorConfigLister) Get(name string) (*v1alpha1.AuthenticationOperatorConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("authenticationoperatorconfig"), name)
	}
	return obj.(*v1alpha1.AuthenticationOperatorConfig), nil
}
