// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationExecutionInfo) DeepCopyInto(out *AuthenticationExecutionInfo) {
	*out = *in
	if in.RequirementChoices != nil {
		in, out := &in.RequirementChoices, &out.RequirementChoices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationExecutionInfo.
func (in *AuthenticationExecutionInfo) DeepCopy() *AuthenticationExecutionInfo {
	if in == nil {
		return nil
	}
	out := new(AuthenticationExecutionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticatorConfig) DeepCopyInto(out *AuthenticatorConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticatorConfig.
func (in *AuthenticatorConfig) DeepCopy() *AuthenticatorConfig {
	if in == nil {
		return nil
	}
	out := new(AuthenticatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfig) DeepCopyInto(out *BackupConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfig.
func (in *BackupConfig) DeepCopy() *BackupConfig {
	if in == nil {
		return nil
	}
	out := new(BackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentity) DeepCopyInto(out *FederatedIdentity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentity.
func (in *FederatedIdentity) DeepCopy() *FederatedIdentity {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keycloak) DeepCopyInto(out *Keycloak) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keycloak.
func (in *Keycloak) DeepCopy() *Keycloak {
	if in == nil {
		return nil
	}
	out := new(Keycloak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Keycloak) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIAuthenticationExecution) DeepCopyInto(out *KeycloakAPIAuthenticationExecution) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIAuthenticationExecution.
func (in *KeycloakAPIAuthenticationExecution) DeepCopy() *KeycloakAPIAuthenticationExecution {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIAuthenticationExecution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIAuthenticationFlow) DeepCopyInto(out *KeycloakAPIAuthenticationFlow) {
	*out = *in
	if in.AuthenticationExecutions != nil {
		in, out := &in.AuthenticationExecutions, &out.AuthenticationExecutions
		*out = make([]KeycloakAPIAuthenticationExecution, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIAuthenticationFlow.
func (in *KeycloakAPIAuthenticationFlow) DeepCopy() *KeycloakAPIAuthenticationFlow {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIAuthenticationFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIAuthenticatorConfig) DeepCopyInto(out *KeycloakAPIAuthenticatorConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIAuthenticatorConfig.
func (in *KeycloakAPIAuthenticatorConfig) DeepCopy() *KeycloakAPIAuthenticatorConfig {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIAuthenticatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIClient) DeepCopyInto(out *KeycloakAPIClient) {
	*out = *in
	if in.DefaultRoles != nil {
		in, out := &in.DefaultRoles, &out.DefaultRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RedirectUris != nil {
		in, out := &in.RedirectUris, &out.RedirectUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WebOrigins != nil {
		in, out := &in.WebOrigins, &out.WebOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProtocolMappers != nil {
		in, out := &in.ProtocolMappers, &out.ProtocolMappers
		*out = make([]KeycloakProtocolMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OptionalClientScopes != nil {
		in, out := &in.OptionalClientScopes, &out.OptionalClientScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultClientScopes != nil {
		in, out := &in.DefaultClientScopes, &out.DefaultClientScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIClient.
func (in *KeycloakAPIClient) DeepCopy() *KeycloakAPIClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIPasswordReset) DeepCopyInto(out *KeycloakAPIPasswordReset) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIPasswordReset.
func (in *KeycloakAPIPasswordReset) DeepCopy() *KeycloakAPIPasswordReset {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIPasswordReset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIRealm) DeepCopyInto(out *KeycloakAPIRealm) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]*KeycloakAPIUser, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakAPIUser)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Clients != nil {
		in, out := &in.Clients, &out.Clients
		*out = make([]*KeycloakAPIClient, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakAPIClient)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]*KeycloakIdentityProvider, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakIdentityProvider)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.EventsListeners != nil {
		in, out := &in.EventsListeners, &out.EventsListeners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EventsEnabled != nil {
		in, out := &in.EventsEnabled, &out.EventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AdminEventsEnabled != nil {
		in, out := &in.AdminEventsEnabled, &out.AdminEventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AdminEventsDetailsEnabled != nil {
		in, out := &in.AdminEventsDetailsEnabled, &out.AdminEventsDetailsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ClientScopes != nil {
		in, out := &in.ClientScopes, &out.ClientScopes
		*out = make([]KeycloakClientScope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthenticationFlows != nil {
		in, out := &in.AuthenticationFlows, &out.AuthenticationFlows
		*out = make([]KeycloakAPIAuthenticationFlow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthenticatorConfig != nil {
		in, out := &in.AuthenticatorConfig, &out.AuthenticatorConfig
		*out = make([]KeycloakAPIAuthenticatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIRealm.
func (in *KeycloakAPIRealm) DeepCopy() *KeycloakAPIRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIUser) DeepCopyInto(out *KeycloakAPIUser) {
	*out = *in
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClientRoles != nil {
		in, out := &in.ClientRoles, &out.ClientRoles
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.RequiredActions != nil {
		in, out := &in.RequiredActions, &out.RequiredActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FederatedIdentities != nil {
		in, out := &in.FederatedIdentities, &out.FederatedIdentities
		*out = make([]FederatedIdentity, len(*in))
		copy(*out, *in)
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]KeycloakCredential, len(*in))
		copy(*out, *in)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIUser.
func (in *KeycloakAPIUser) DeepCopy() *KeycloakAPIUser {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAWSSpec) DeepCopyInto(out *KeycloakAWSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAWSSpec.
func (in *KeycloakAWSSpec) DeepCopy() *KeycloakAWSSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakAWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakBackup) DeepCopyInto(out *KeycloakBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakBackup.
func (in *KeycloakBackup) DeepCopy() *KeycloakBackup {
	if in == nil {
		return nil
	}
	out := new(KeycloakBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakBackupList) DeepCopyInto(out *KeycloakBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakBackupList.
func (in *KeycloakBackupList) DeepCopy() *KeycloakBackupList {
	if in == nil {
		return nil
	}
	out := new(KeycloakBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakBackupSpec) DeepCopyInto(out *KeycloakBackupSpec) {
	*out = *in
	out.AWS = in.AWS
	if in.InstanceSelector != nil {
		in, out := &in.InstanceSelector, &out.InstanceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakBackupSpec.
func (in *KeycloakBackupSpec) DeepCopy() *KeycloakBackupSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakBackupStatus) DeepCopyInto(out *KeycloakBackupStatus) {
	*out = *in
	if in.SecondaryResources != nil {
		in, out := &in.SecondaryResources, &out.SecondaryResources
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakBackupStatus.
func (in *KeycloakBackupStatus) DeepCopy() *KeycloakBackupStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClient) DeepCopyInto(out *KeycloakClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClient.
func (in *KeycloakClient) DeepCopy() *KeycloakClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientList) DeepCopyInto(out *KeycloakClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientList.
func (in *KeycloakClientList) DeepCopy() *KeycloakClientList {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientScope) DeepCopyInto(out *KeycloakClientScope) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProtocolMappers != nil {
		in, out := &in.ProtocolMappers, &out.ProtocolMappers
		*out = make([]KeycloakProtocolMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientScope.
func (in *KeycloakClientScope) DeepCopy() *KeycloakClientScope {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientSpec) DeepCopyInto(out *KeycloakClientSpec) {
	*out = *in
	if in.RealmSelector != nil {
		in, out := &in.RealmSelector, &out.RealmSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(KeycloakAPIClient)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientSpec.
func (in *KeycloakClientSpec) DeepCopy() *KeycloakClientSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientStatus) DeepCopyInto(out *KeycloakClientStatus) {
	*out = *in
	if in.SecondaryResources != nil {
		in, out := &in.SecondaryResources, &out.SecondaryResources
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientStatus.
func (in *KeycloakClientStatus) DeepCopy() *KeycloakClientStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakCredential) DeepCopyInto(out *KeycloakCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakCredential.
func (in *KeycloakCredential) DeepCopy() *KeycloakCredential {
	if in == nil {
		return nil
	}
	out := new(KeycloakCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakExternalAccess) DeepCopyInto(out *KeycloakExternalAccess) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakExternalAccess.
func (in *KeycloakExternalAccess) DeepCopy() *KeycloakExternalAccess {
	if in == nil {
		return nil
	}
	out := new(KeycloakExternalAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakExternalDatabase) DeepCopyInto(out *KeycloakExternalDatabase) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakExternalDatabase.
func (in *KeycloakExternalDatabase) DeepCopy() *KeycloakExternalDatabase {
	if in == nil {
		return nil
	}
	out := new(KeycloakExternalDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakIdentityProvider) DeepCopyInto(out *KeycloakIdentityProvider) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakIdentityProvider.
func (in *KeycloakIdentityProvider) DeepCopy() *KeycloakIdentityProvider {
	if in == nil {
		return nil
	}
	out := new(KeycloakIdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakList) DeepCopyInto(out *KeycloakList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Keycloak, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakList.
func (in *KeycloakList) DeepCopy() *KeycloakList {
	if in == nil {
		return nil
	}
	out := new(KeycloakList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakProtocolMapper) DeepCopyInto(out *KeycloakProtocolMapper) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakProtocolMapper.
func (in *KeycloakProtocolMapper) DeepCopy() *KeycloakProtocolMapper {
	if in == nil {
		return nil
	}
	out := new(KeycloakProtocolMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealm) DeepCopyInto(out *KeycloakRealm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealm.
func (in *KeycloakRealm) DeepCopy() *KeycloakRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmList) DeepCopyInto(out *KeycloakRealmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmList.
func (in *KeycloakRealmList) DeepCopy() *KeycloakRealmList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmSpec) DeepCopyInto(out *KeycloakRealmSpec) {
	*out = *in
	if in.InstanceSelector != nil {
		in, out := &in.InstanceSelector, &out.InstanceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(KeycloakAPIRealm)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmOverrides != nil {
		in, out := &in.RealmOverrides, &out.RealmOverrides
		*out = make([]*RedirectorIdentityProviderOverride, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RedirectorIdentityProviderOverride)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmSpec.
func (in *KeycloakRealmSpec) DeepCopy() *KeycloakRealmSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmStatus) DeepCopyInto(out *KeycloakRealmStatus) {
	*out = *in
	if in.SecondaryResources != nil {
		in, out := &in.SecondaryResources, &out.SecondaryResources
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmStatus.
func (in *KeycloakRealmStatus) DeepCopy() *KeycloakRealmStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ExternalAccess = in.ExternalAccess
	out.ExternalDatabase = in.ExternalDatabase
	out.PodDisruptionBudget = in.PodDisruptionBudget
	in.KeycloakDeploymentSpec.DeepCopyInto(&out.KeycloakDeploymentSpec)
	in.PostgresDeploymentSpec.DeepCopyInto(&out.PostgresDeploymentSpec)
	out.Migration = in.Migration
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakStatus) DeepCopyInto(out *KeycloakStatus) {
	*out = *in
	if in.SecondaryResources != nil {
		in, out := &in.SecondaryResources, &out.SecondaryResources
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakStatus.
func (in *KeycloakStatus) DeepCopy() *KeycloakStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUser) DeepCopyInto(out *KeycloakUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUser.
func (in *KeycloakUser) DeepCopy() *KeycloakUser {
	if in == nil {
		return nil
	}
	out := new(KeycloakUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserList) DeepCopyInto(out *KeycloakUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserList.
func (in *KeycloakUserList) DeepCopy() *KeycloakUserList {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserRole) DeepCopyInto(out *KeycloakUserRole) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserRole.
func (in *KeycloakUserRole) DeepCopy() *KeycloakUserRole {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserSpec) DeepCopyInto(out *KeycloakUserSpec) {
	*out = *in
	if in.RealmSelector != nil {
		in, out := &in.RealmSelector, &out.RealmSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.User.DeepCopyInto(&out.User)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserSpec.
func (in *KeycloakUserSpec) DeepCopy() *KeycloakUserSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserStatus) DeepCopyInto(out *KeycloakUserStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserStatus.
func (in *KeycloakUserStatus) DeepCopy() *KeycloakUserStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrateConfig) DeepCopyInto(out *MigrateConfig) {
	*out = *in
	out.Backups = in.Backups
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrateConfig.
func (in *MigrateConfig) DeepCopy() *MigrateConfig {
	if in == nil {
		return nil
	}
	out := new(MigrateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetConfig) DeepCopyInto(out *PodDisruptionBudgetConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetConfig.
func (in *PodDisruptionBudgetConfig) DeepCopy() *PodDisruptionBudgetConfig {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedirectorIdentityProviderOverride) DeepCopyInto(out *RedirectorIdentityProviderOverride) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedirectorIdentityProviderOverride.
func (in *RedirectorIdentityProviderOverride) DeepCopy() *RedirectorIdentityProviderOverride {
	if in == nil {
		return nil
	}
	out := new(RedirectorIdentityProviderOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenResponse) DeepCopyInto(out *TokenResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenResponse.
func (in *TokenResponse) DeepCopy() *TokenResponse {
	if in == nil {
		return nil
	}
	out := new(TokenResponse)
	in.DeepCopyInto(out)
	return out
}
