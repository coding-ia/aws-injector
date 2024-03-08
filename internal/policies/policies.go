package policies

type Policies struct {
	DaemonSets  []*Policy `json:"DaemonSets"`
	Deployments []*Policy `json:"Deployments"`
	ConfigMaps  []*Policy `json:"ConfigMaps"`
}

type Policy struct {
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Type        string `json:"keyType"`
	Value       string `json:"value,omitempty"`
	SkipAdd     bool   `json:"skipAdd"`
	SkipReplace bool   `json:"skipReplace"`
}

func FindDeploymentPolicy(namespace string, name string, keyType string) (*Policy, error) {
	var policy *Policy

	if namespace == "sentinel" && name == "" && keyType == "env" {
		policy = &Policy{
			Namespace: namespace,
			Name:      name,
			Key:       "CLUSTER_NAME",
			Type:      keyType,
		}
	}

	return policy, nil
}

func FindDaemonSetPolicy(namespace string, name string, keyType string) (*Policy, error) {
	var policy *Policy

	if namespace == "aqua" && name == "aqua-enforcer-ds" && keyType == "env" {
		policy = &Policy{
			Namespace: namespace,
			Name:      name,
			Key:       "AQUA_LOGICAL_NAME",
			Type:      keyType,
		}
	}

	return policy, nil
}

func FindConfigMapPolicy(namespace string, name string, keyType string) (*Policy, error) {
	var policy *Policy

	if namespace == "aqua" && name == "aqua-csp-kube-enforcer" && keyType == "" {
		policy = &Policy{
			Namespace: namespace,
			Name:      name,
			Key:       "AQUA_LOGICAL_NAME",
			Type:      "",
		}
	}

	return policy, nil
}