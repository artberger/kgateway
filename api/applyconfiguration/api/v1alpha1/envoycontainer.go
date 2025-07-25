// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// EnvoyContainerApplyConfiguration represents a declarative configuration of the EnvoyContainer type for use
// with apply.
type EnvoyContainerApplyConfiguration struct {
	Bootstrap       *EnvoyBootstrapApplyConfiguration `json:"bootstrap,omitempty"`
	Image           *ImageApplyConfiguration          `json:"image,omitempty"`
	SecurityContext *v1.SecurityContext               `json:"securityContext,omitempty"`
	Resources       *v1.ResourceRequirements          `json:"resources,omitempty"`
	Env             []v1.EnvVar                       `json:"env,omitempty"`
}

// EnvoyContainerApplyConfiguration constructs a declarative configuration of the EnvoyContainer type for use with
// apply.
func EnvoyContainer() *EnvoyContainerApplyConfiguration {
	return &EnvoyContainerApplyConfiguration{}
}

// WithBootstrap sets the Bootstrap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bootstrap field is set to the value of the last call.
func (b *EnvoyContainerApplyConfiguration) WithBootstrap(value *EnvoyBootstrapApplyConfiguration) *EnvoyContainerApplyConfiguration {
	b.Bootstrap = value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *EnvoyContainerApplyConfiguration) WithImage(value *ImageApplyConfiguration) *EnvoyContainerApplyConfiguration {
	b.Image = value
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *EnvoyContainerApplyConfiguration) WithSecurityContext(value v1.SecurityContext) *EnvoyContainerApplyConfiguration {
	b.SecurityContext = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *EnvoyContainerApplyConfiguration) WithResources(value v1.ResourceRequirements) *EnvoyContainerApplyConfiguration {
	b.Resources = &value
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *EnvoyContainerApplyConfiguration) WithEnv(values ...v1.EnvVar) *EnvoyContainerApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}
