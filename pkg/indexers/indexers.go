/*
Copyright 2022 The Kube Bind Authors.

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

package indexers

import (
	kubebindv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

const (
	ByKubeconfigSecret = "byKubeconfigSecret"
)

func IndexByKubeconfigSecret(obj interface{}) ([]string, error) {
	binding, ok := obj.(*kubebindv1alpha1.ServiceBinding)
	if !ok {
		return nil, nil
	}
	return []string{ByKubeconfigSecretKey(binding)}, nil
}

func ByKubeconfigSecretKey(binding *kubebindv1alpha1.ServiceBinding) string {
	ref := &binding.Spec.KubeconfigSecretRef
	return ref.Namespace + "/" + ref.Name
}