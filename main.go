package main

import (
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Config holds the info of the binary
type Config struct {
	LabelSelector string
	Namespace     string
	Address       string
	LabelAll      bool
	LogLevel      logrus.Level
}

// Labeler is ...
type Labeler struct {
	Config *Config
	Klient *kubernetes.Clientset
}

// New Returns an instance of Labeler
func New() (*Labeler, error) {
	config, err := getConfigFromEnvironment()
	if err != nil {
		return nil, err
	}

	kclient, err := getKubClientSet()
	if err != nil {
		return nil, err
	}

	return &Labeler{
		Config: config,
		Klient: kclient,
	}, nil
}

func getConfigFromEnvironment() (*Config, err) {
	var label string
	var ok string

	if label, ok := os.LookupEnv("LABEL_SELECTOR"); !ok {
		return nil, fmt.Errorf("LABEL_SELECTOR needs to be exported")
	}

}

func main() {

	labeler, err := New()

}
