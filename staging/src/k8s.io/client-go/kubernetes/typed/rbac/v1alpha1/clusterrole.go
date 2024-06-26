/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rbacv1alpha1 "k8s.io/client-go/applyconfigurations/rbac/v1alpha1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ClusterRolesGetter has a method to return a ClusterRoleInterface.
// A group's client should implement this interface.
type ClusterRolesGetter interface {
	ClusterRoles() ClusterRoleInterface
}

// ClusterRoleInterface has methods to work with ClusterRole resources.
type ClusterRoleInterface interface {
	Create(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.CreateOptions) (*v1alpha1.ClusterRole, error)
	Update(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.UpdateOptions) (*v1alpha1.ClusterRole, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterRole, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRole, err error)
	Apply(ctx context.Context, clusterRole *rbacv1alpha1.ClusterRoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterRole, err error)
	ClusterRoleExpansion
}

// clusterRoles implements ClusterRoleInterface
type clusterRoles struct {
	*gentype.ClientWithListAndApply[*v1alpha1.ClusterRole, *v1alpha1.ClusterRoleList, *rbacv1alpha1.ClusterRoleApplyConfiguration]
}

// newClusterRoles returns a ClusterRoles
func newClusterRoles(c *RbacV1alpha1Client) *clusterRoles {
	return &clusterRoles{
		gentype.NewClientWithListAndApply[*v1alpha1.ClusterRole, *v1alpha1.ClusterRoleList, *rbacv1alpha1.ClusterRoleApplyConfiguration](
			"clusterroles",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.ClusterRole { return &v1alpha1.ClusterRole{} },
			func() *v1alpha1.ClusterRoleList { return &v1alpha1.ClusterRoleList{} }),
	}
}
