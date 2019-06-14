package starter

import "k8s.io/client-go/kubernetes"

func ClientWithoutSecret(client *kubernetes.Clientset) *kubernetes.Clientset {
	return client
}
