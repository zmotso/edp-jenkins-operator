package client

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	jenkinsApi "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1"
)

//var k8sConfig clientcmd.ClientConfig
var SchemeGroupVersion = schema.GroupVersion{Group: "v2.edp.epam.com", Version: "v1"}

type EdpV1Client struct {
	crClient *rest.RESTClient
}

type JenkinsInterface interface {
	Get(name string, namespace string, options metav1.GetOptions) (result *jenkinsApi.Jenkins, err error)
	Create(jsa *jenkinsApi.Jenkins, namespace string) (result *jenkinsApi.Jenkins, err error)
	Update(jsa *jenkinsApi.Jenkins) (result *jenkinsApi.Jenkins, err error)
}

func NewForConfig(config *rest.Config) (*EdpV1Client, error) {
	if err := createCrdClient(config); err != nil {
		return nil, err
	}
	crClient, err := rest.RESTClientFor(config)
	if err != nil {
		return nil, err
	}
	return &EdpV1Client{crClient: crClient}, nil
}

func (c *EdpV1Client) Get(ctx context.Context, name string, namespace string, options metav1.GetOptions) (result *jenkinsApi.Jenkins, err error) {
	result = &jenkinsApi.Jenkins{}
	err = c.crClient.Get().
		Namespace(namespace).
		Resource("jenkins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// Create takes the representation of a Jenkins and creates it.  Returns the server's representation of the Jenkins, and an error, if there is any.
func (c *EdpV1Client) Create(ctx context.Context, jsa *jenkinsApi.Jenkins, namespace string) (result *jenkinsApi.Jenkins, err error) {
	result = &jenkinsApi.Jenkins{}
	err = c.crClient.Post().
		Namespace(namespace).
		Resource("jenkins").
		Body(jsa).
		Do(ctx).
		Into(result)
	return
}

func (c *EdpV1Client) Update(ctx context.Context, jsa *jenkinsApi.Jenkins) (result *jenkinsApi.Jenkins, err error) {
	result = &jenkinsApi.Jenkins{}
	err = c.crClient.Put().
		Namespace(jsa.Namespace).
		Resource("jenkins").
		Name(jsa.Name).
		Body(jsa).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Jenkins that match those selectors.
func (c *EdpV1Client) List(ctx context.Context, opts metav1.ListOptions, namespace string) (result *jenkinsApi.JenkinsList, err error) {
	result = &jenkinsApi.JenkinsList{}
	err = c.crClient.Get().
		Namespace(namespace).
		Resource("jenkins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func createCrdClient(cfg *rest.Config) error {
	scheme := runtime.NewScheme()
	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return err
	}
	config := cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme)

	return nil
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&jenkinsApi.Jenkins{},
		&jenkinsApi.JenkinsList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
