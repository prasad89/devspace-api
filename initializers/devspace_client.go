package initializers

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	customclientset "github.com/prasad89/devspace-operator/pkg/generated/clientset/versioned"
)

var DevspaceClient customclientset.Interface

// InitDevspaceClient initializes the global DevspaceClient
func InitDevspaceClient() {
	var kubeConfig string

	// Parse kubeconfig from flag
	flag.StringVar(&kubeConfig, "kubeConfig", "", "Path to the kubeconfig file")
	flag.Parse()

	// Check environment variable if flag is not specified
	if kubeConfig == "" {
		kubeConfig = os.Getenv("KUBECONFIG")
	}

	// Fallback to ~/.kube/config
	if kubeConfig == "" {
		home, err := os.UserHomeDir()
		if err == nil {
			kubeConfig = filepath.Join(home, ".kube", "config")
		}
	}

	var config *rest.Config
	var err error

	if _, err := os.Stat(kubeConfig); err == nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			log.Fatal("❌ Failed to build config from kubeconfig:", err)
			os.Exit(1)
		}
		log.Println("✅ Using kubeconfig from:", kubeConfig)
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatal("❌ Failed to load in-cluster config:", err)
			os.Exit(1)
		}
		log.Println("✅ Using in-cluster Kubernetes config")
	}

	DevspaceClient, err = customclientset.NewForConfig(config)
	if err != nil {
		log.Fatal("❌ Failed to create DevSpace clientset:", err)
		os.Exit(1)
	}

	log.Println("✅ DevSpace clientset initialized successfully!")
}
