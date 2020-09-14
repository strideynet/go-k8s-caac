package main

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
)

// This file shows a very simple example of a very simple deployment
// This gives us a feel for writing definitions in Go.

var podLabels = map[string]string{
	"app.kubernetes.io/name":      "sheep",
	"app.kubernetes.io/component": "api",
	"app.kubernetes.io/part-of":   "farm",
}

var replicaCount int32 = 3

var containers = []corev1.Container{
	{
		Name:  "api",
		Image: "gcr.io/farm-manager/sheep-api:latest",

		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: 80,
			},
		},
	},
}

var deployment = appsv1.Deployment{
	TypeMeta: v1.TypeMeta{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "sheep-api",
		Namespace: "default",
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: &replicaCount,
		Selector: &v1.LabelSelector{
			MatchLabels: podLabels,
		},

		Template: corev1.PodTemplateSpec{
			ObjectMeta: v1.ObjectMeta{
				Labels: podLabels,
			},

			Spec: corev1.PodSpec{
				Containers: containers,
			},
		},
	},
}

func main() {
	// marshal basic deployment spec
	jsonBytes, err := json.Marshal(deployment)
	if err != nil {
		log.Fatalf("err occured marshalling json: %s", err)
	}

	// write pod spec to file so we can inspect
	err = ioutil.WriteFile("output.json", jsonBytes, os.ModePerm)
	if err != nil {
		log.Fatalf("err occured writing json: %s", err)
	}

	// some random, unvetted, third party package to see how YAML would work
	// using the default yaml package doesn't work as the structs arent tagged correctly
	yamlBytes, err := yaml.Marshal(deployment)
	if err != nil {
		log.Fatalf("err occurred marshalling yaml: %s", err)
	}

	err = ioutil.WriteFile("output.yaml", yamlBytes, os.ModePerm)
	if err != nil {
		log.Fatalf("err occured writing json: %s", err)
	}
}
