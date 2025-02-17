package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	cdPipeApi "github.com/epam/edp-cd-pipeline-operator/v2/pkg/apis/edp/v1"
	edpCompApi "github.com/epam/edp-component-operator/pkg/apis/v1/v1"
	"github.com/epam/edp-gerrit-operator/v2/pkg/service/helpers"
	keycloakV1Api "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1"
	"github.com/pkg/errors"
	coreV1Api "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	appsV1Client "k8s.io/client-go/kubernetes/typed/apps/v1"
	networkingV1Client "k8s.io/client-go/kubernetes/typed/networking/v1"
	authV1Client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	jenkinsApi "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1"
	"github.com/epam/edp-jenkins-operator/v2/pkg/model"
	jenkinsDefaultSpec "github.com/epam/edp-jenkins-operator/v2/pkg/service/jenkins/spec"
	platformHelper "github.com/epam/edp-jenkins-operator/v2/pkg/service/platform/helper"
)

var log = ctrl.Log.WithName("platform")

// K8SService struct for K8S platform service
type K8SService struct {
	Scheme           *runtime.Scheme
	client           client.Client
	authClient       authV1Client.RbacV1Client
	networkingClient networkingV1Client.NetworkingV1Interface
	appsV1Client     appsV1Client.AppsV1Interface
}

// Init initializes K8SService
func (s *K8SService) Init(config *rest.Config, Scheme *runtime.Scheme, k8sClient client.Client) error {
	authClient, err := authV1Client.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "Failed to init auth V1 client for K8S")
	}

	s.Scheme = Scheme
	s.authClient = *authClient
	s.client = k8sClient

	ncl, err := networkingV1Client.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "Failed to init networking V1 client for K8S")
	}
	s.networkingClient = ncl

	appsClient, err := appsV1Client.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "Failed to init apps V1 client for K8S")
	}
	s.appsV1Client = appsClient

	return nil
}

// GetExternalEndpoint returns Ingress object and connection protocol from Kubernetes
func (s K8SService) GetExternalEndpoint(namespace string, name string) (string, string, string, error) {
	ingress, err := s.networkingClient.Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil && k8sErrors.IsNotFound(err) {
		return "", "", "", errors.New(fmt.Sprintf("Ingress %v in namespace %v not found", name, namespace))
	} else if err != nil {
		return "", "", "", err
	}

	return ingress.Spec.Rules[0].Host, jenkinsDefaultSpec.RouteHTTPSScheme,
		strings.TrimRight(ingress.Spec.Rules[0].HTTP.Paths[0].Path, platformHelper.UrlCutset), nil
}

// AddVolumeToInitContainer adds volume to Jenkins init container
func (s K8SService) AddVolumeToInitContainer(instance *jenkinsApi.Jenkins, containerName string,
	vol []coreV1Api.Volume, volMount []coreV1Api.VolumeMount) error {

	if len(vol) == 0 || len(volMount) == 0 {
		return nil
	}

	deployment, err := s.appsV1Client.Deployments(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil
	}

	initContainer, err := selectContainer(deployment.Spec.Template.Spec.InitContainers, containerName)
	if err != nil {
		return err
	}

	initContainer.VolumeMounts = updateVolumeMounts(initContainer.VolumeMounts, volMount)
	deployment.Spec.Template.Spec.InitContainers = append(deployment.Spec.Template.Spec.InitContainers, initContainer)
	volumes := deployment.Spec.Template.Spec.Volumes
	volumes = updateVolumes(volumes, vol)
	deployment.Spec.Template.Spec.Volumes = volumes

	jsonDc, err := json.Marshal(deployment)
	if err != nil {
		return err
	}

	_, err = s.appsV1Client.Deployments(deployment.Namespace).Patch(context.TODO(), deployment.Name, types.StrategicMergePatchType, jsonDc, metav1.PatchOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s K8SService) IsDeploymentReady(instance jenkinsApi.Jenkins) (bool, error) {
	deployment, err := s.appsV1Client.Deployments(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	if deployment.Status.UpdatedReplicas == 1 && deployment.Status.AvailableReplicas == 1 {
		return true, nil
	}

	return false, nil
}

//CreateSecret creates secret object in K8s cluster
func (s K8SService) CreateSecret(instance *jenkinsApi.Jenkins, name string, data map[string][]byte) error {
	_, err := s.getSecret(name, instance.Namespace)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			secret := &coreV1Api.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: instance.Namespace,
					Labels:    platformHelper.GenerateLabels(instance.Name),
				},
				Data: data,
				Type: "Opaque",
			}
			return s.createSecret(instance, secret)
		}
		return errors.Wrapf(err, "Couldn't get Secret %v object", name)
	}
	return nil
}

func (s K8SService) getSecret(name, namespace string) (*coreV1Api.Secret, error) {
	secret := &coreV1Api.Secret{}
	err := s.client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, secret)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (s K8SService) createSecret(jenkins *jenkinsApi.Jenkins, secret *coreV1Api.Secret) error {
	if err := controllerutil.SetControllerReference(jenkins, secret, s.Scheme); err != nil {
		return errors.Wrapf(err, "Couldn't set reference for Secret %v object", secret.Name)
	}

	if err := s.client.Create(context.TODO(), secret); err != nil {
		return errors.Wrapf(err, "Couldn't create Secret %v object", secret.Name)
	}
	log.Info(fmt.Sprintf("Secret %v has been created", secret.Name))
	return nil
}

// GetSecretData return data field of Secret
func (s K8SService) GetSecretData(namespace, name string) (map[string][]byte, error) {
	secret, err := s.getSecret(name, namespace)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			log.Info(fmt.Sprintf("Secret %v in namespace %v not found", name, namespace))
			return nil, nil
		}
		return nil, err
	}
	return secret.Data, nil
}

func (s K8SService) CreateConfigMapWithUpdate(instance *jenkinsApi.Jenkins, name string, data map[string]string,
	labels ...map[string]string) (isUpdated bool, err error) {
	currentConfigMap, err := s.CreateConfigMap(instance, name, data, labels...)
	if err != nil {
		return false, errors.Wrap(err, "unable to create configmap")
	}

	if reflect.DeepEqual(data, currentConfigMap.Data) {
		return false, nil
	}

	currentConfigMap.Data = data
	if err := s.client.Update(context.TODO(), currentConfigMap); err != nil {
		return false, errors.Wrap(err, "unable to update config map")
	}

	return true, nil
}

func (s K8SService) CreateConfigMap(instance *jenkinsApi.Jenkins, name string, data map[string]string,
	labels ...map[string]string) (*coreV1Api.ConfigMap, error) {

	currentConfigMap, err := s.getConfigMap(name, instance.Namespace)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			resultLabels := platformHelper.GenerateLabels(instance.Name)
			if len(labels) != 0 {
				resultLabels = labels[0]
			}
			currentConfigMap = &coreV1Api.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: instance.Namespace,
					Labels:    resultLabels,
				},
				Data: data,
			}

			if err := s.createConfigMap(instance, currentConfigMap); err != nil {
				return nil, errors.Wrap(err, "unable to create config map")
			}

			return currentConfigMap, nil
		}

		return nil, errors.Wrapf(err, "Couldn't get ConfigMap %v object", name)
	}

	return currentConfigMap, nil
}

func (s K8SService) getConfigMap(name, namespace string) (*coreV1Api.ConfigMap, error) {
	cm := &coreV1Api.ConfigMap{}
	err := s.client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, cm)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (s K8SService) createConfigMap(jenkins *jenkinsApi.Jenkins, cm *coreV1Api.ConfigMap) error {
	if err := controllerutil.SetControllerReference(jenkins, cm, s.Scheme); err != nil {
		return errors.Wrapf(err, "Couldn't set reference for Config Map %v object", cm.Name)
	}

	if err := s.client.Create(context.TODO(), cm); err != nil {
		return errors.Wrapf(err, "Couldn't create Config Map %v object", cm.Name)
	}
	log.Info(fmt.Sprintf("ConfigMap %s/%s has been created", cm.Namespace, cm.Name))
	return nil
}

// CreateConfigMapFromFileOrDir performs creating ConfigMap in K8S
func (s K8SService) CreateConfigMapFromFileOrDir(instance *jenkinsApi.Jenkins, configMapName string,
	configMapKey *string, path string, ownerReference metav1.Object, customLabels ...map[string]string) error {
	configMapData, err := s.fillConfigMapData(path, configMapKey)
	if err != nil {
		return errors.Wrapf(err, "Couldn't generate Config Map data for %v", configMapName)
	}

	labels := platformHelper.GenerateLabels(instance.Name)
	if len(customLabels) == 1 {
		for key, value := range customLabels[0] {
			labels[key] = value
		}
	}

	_, err = s.CreateConfigMap(instance, configMapName, configMapData, labels)
	if err != nil {
		return errors.Wrapf(err, "Failed to create Config Map %v", configMapName)
	}

	return nil
}

func (s K8SService) CreateEDPComponentIfNotExist(jen jenkinsApi.Jenkins, url string, icon string) error {
	c, err := s.getEDPComponent(jen.Name, jen.Namespace)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			return s.createEDPComponent(jen, url, icon)
		}
		return errors.Wrapf(err, "failed to get edp component: %v", jen.Name)
	}

	log.V(1).Info("edp component already exists", "name", c.Name)
	return nil
}

func (s K8SService) getEDPComponent(name, namespace string) (*edpCompApi.EDPComponent, error) {
	c := &edpCompApi.EDPComponent{}
	err := s.client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func (s K8SService) createEDPComponent(jen jenkinsApi.Jenkins, url string, icon string) error {
	c := &edpCompApi.EDPComponent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jen.Name,
			Namespace: jen.Namespace,
		},
		Spec: edpCompApi.EDPComponentSpec{
			Type:    "jenkins",
			Url:     url,
			Icon:    icon,
			Visible: true,
		},
	}
	if err := controllerutil.SetControllerReference(&jen, c, s.Scheme); err != nil {
		return err
	}

	return s.client.Create(context.TODO(), c)
}

func (s K8SService) fillConfigMapData(path string, configMapKey *string) (map[string]string, error) {
	var configMapData map[string]string
	pathInfo, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Couldn't open path %v.", path))
	}
	if pathInfo.Mode().IsDir() {
		configMapData, err = s.fillConfigMapFromDir(path)
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("Couldn't generate config map data from directory %v", path))
		}
	} else {
		configMapData, err = s.fillConfigMapFromFile(path, configMapKey)
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("Couldn't generate config map data from file %v", path))
		}
	}
	return configMapData, nil
}

func (s K8SService) fillConfigMapFromFile(path string, configMapKey *string) (map[string]string, error) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Couldn't read file %v.", path))
	}
	key := filepath.Base(path)
	if configMapKey != nil {
		key = *configMapKey
	}
	configMapData := map[string]string{
		key: string(content),
	}
	return configMapData, nil
}

func (s K8SService) fillConfigMapFromDir(path string) (map[string]string, error) {
	configMapData := make(map[string]string)
	directory, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Couldn't open path %v.", path))
	}
	for _, file := range directory {
		content, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", path, file.Name()))
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("Couldn't open path %v.", path))
		}
		configMapData[file.Name()] = string(content)
	}
	return configMapData, nil
}

// GetConfigMapData return data field of ConfigMap
func (s K8SService) GetConfigMapData(namespace, name string) (map[string]string, error) {
	cm, err := s.getConfigMap(name, namespace)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			return nil, errors.Wrapf(err, "Config map %v in namespace %v not found", name, namespace)
		}
		return nil, errors.Wrapf(err, "Couldn't get ConfigMap %v object", name)
	}
	return cm.Data, nil
}

func (s K8SService) CreateKeycloakClient(kc *keycloakV1Api.KeycloakClient) error {
	nsn := types.NamespacedName{
		Namespace: kc.Namespace,
		Name:      kc.Name,
	}

	err := s.client.Get(context.TODO(), nsn, kc)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			err := s.client.Create(context.TODO(), kc)
			if err != nil {
				return errors.Wrapf(err, "Failed to create Keycloak client %s/%s", kc.Namespace, kc.Name)
			}
			log.Info(fmt.Sprintf("Keycloak client %s/%s created", kc.Namespace, kc.Name))
			return nil
		}
		return errors.Wrapf(err, "Failed to create Keycloak client %s/%s", kc.Namespace, kc.Name)
	}

	return nil
}

func (s K8SService) GetKeycloakClient(name, namespace string) (keycloakV1Api.KeycloakClient, error) {
	out := keycloakV1Api.KeycloakClient{}
	nsn := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}

	err := s.client.Get(context.TODO(), nsn, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (s K8SService) CreateJenkinsScript(namespace, name string, forceRecreate bool) (*jenkinsApi.JenkinsScript, error) {
	js, err := s.getJenkinsScript(name, namespace)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			js := &jenkinsApi.JenkinsScript{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: jenkinsApi.JenkinsScriptSpec{
					SourceCmName: name,
				},
			}
			if err := s.client.Create(context.TODO(), js); err != nil {
				return nil, errors.Wrap(err, "unable to create jenkins script")
			}
			return js, nil
		}
		return nil, err
	}

	if forceRecreate {
		if err := s.client.Delete(context.TODO(), js); err != nil {
			return nil, errors.Wrap(err, "unable to delete jenkins script")
		}

		js = &jenkinsApi.JenkinsScript{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: jenkinsApi.JenkinsScriptSpec{
				SourceCmName: name,
			},
		}

		if err := s.client.Create(context.TODO(), js); err != nil {
			return nil, errors.Wrap(err, "unable to create jenkins script")
		}
	}

	return js, nil
}

func (s K8SService) getJenkinsScript(name, namespace string) (*jenkinsApi.JenkinsScript, error) {
	js := &jenkinsApi.JenkinsScript{}
	err := s.client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, js)
	if err != nil {
		return nil, err
	}
	return js, nil
}

func selectContainer(containers []coreV1Api.Container, name string) (coreV1Api.Container, error) {
	for _, c := range containers {
		if c.Name == name {
			return c, nil
		}
	}

	return coreV1Api.Container{}, errors.New("No matching container in spec found!")
}

func updateVolumes(existing []coreV1Api.Volume, vol []coreV1Api.Volume) []coreV1Api.Volume {
	var out []coreV1Api.Volume
	var covered []string

	for _, v := range existing {
		newer, ok := findVolume(vol, v.Name)
		if ok {
			covered = append(covered, v.Name)
			out = append(out, newer)
			continue
		}
		out = append(out, v)
	}
	for _, v := range vol {
		if helpers.IsStringInSlice(v.Name, covered) {
			continue
		}
		covered = append(covered, v.Name)
		out = append(out, v)
	}
	return out
}

func updateVolumeMounts(existing []coreV1Api.VolumeMount, volMount []coreV1Api.VolumeMount) []coreV1Api.VolumeMount {
	var out []coreV1Api.VolumeMount
	var covered []string

	for _, v := range existing {
		newer, ok := findVolumeMount(volMount, v.Name)
		if ok {
			covered = append(covered, v.Name)
			out = append(out, newer)
			continue
		}
		out = append(out, v)
	}
	for _, v := range volMount {
		if helpers.IsStringInSlice(v.Name, covered) {
			continue
		}
		covered = append(covered, v.Name)
		out = append(out, v)
	}
	return out
}

func findVolumeMount(volMount []coreV1Api.VolumeMount, name string) (coreV1Api.VolumeMount, bool) {
	for _, v := range volMount {
		if v.Name == name {
			return v, true
		}
	}
	return coreV1Api.VolumeMount{}, false
}

func findVolume(vol []coreV1Api.Volume, name string) (coreV1Api.Volume, bool) {
	for _, v := range vol {
		if v.Name == name {
			return v, true
		}
	}
	return coreV1Api.Volume{}, false
}

func (s K8SService) CreateStageJSON(stage cdPipeApi.Stage) (string, error) {
	j := []model.PipelineStage{
		{
			Name:     "deploy-helm",
			StepName: "deploy-helm",
		},
	}

	for _, ps := range stage.Spec.QualityGates {
		i := model.PipelineStage{
			Name:     ps.QualityGateType,
			StepName: ps.StepName,
		}

		j = append(j, i)
	}
	j = append(j, model.PipelineStage{Name: "promote-images", StepName: "promote-images"})

	o, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(o), err
}
