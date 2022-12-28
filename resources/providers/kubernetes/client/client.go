package client

// // New returns a kubernetes client, defaults to in cluster
// func New() (*kubernetes.Clientset, error) {
// 	return NewInClusterClient()
// }

// // NewInClusterClient returns a client that should run inside the cluster
// func NewInClusterClient() (*kubernetes.Clientset, error) {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return kubernetes.NewForConfig(config)
// }

// // NewExternalClient returns an out of cluster (ie: kubectl style) client
// func NewExternalClient() (*kubernetes.Clientset, error) {
// 	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return kubernetes.NewForConfig(config)
// }
