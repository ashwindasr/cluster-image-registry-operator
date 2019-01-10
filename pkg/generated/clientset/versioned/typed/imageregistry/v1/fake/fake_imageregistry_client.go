// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/cluster-image-registry-operator/pkg/generated/clientset/versioned/typed/imageregistry/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeImageregistryV1 struct {
	*testing.Fake
}

func (c *FakeImageregistryV1) Configs() v1.ConfigInterface {
	return &FakeConfigs{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeImageregistryV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
