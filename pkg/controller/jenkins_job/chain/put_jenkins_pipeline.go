package chain

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	cdPipeApi "github.com/epam/edp-cd-pipeline-operator/v2/pkg/apis/edp/v1"
	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	jenkinsApi "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1"
	jenkinsClient "github.com/epam/edp-jenkins-operator/v2/pkg/client/jenkins"
	jobhandler "github.com/epam/edp-jenkins-operator/v2/pkg/controller/jenkins_job/chain/handler"
	"github.com/epam/edp-jenkins-operator/v2/pkg/service/platform"
	"github.com/epam/edp-jenkins-operator/v2/pkg/util/consts"
	plutil "github.com/epam/edp-jenkins-operator/v2/pkg/util/platform"
)

type PutJenkinsPipeline struct {
	next   jobhandler.JenkinsJobHandler
	client client.Client
	ps     platform.PlatformService
	log    logr.Logger
}

func (h PutJenkinsPipeline) ServeRequest(jj *jenkinsApi.JenkinsJob) error {
	h.log.Info("start creating Jenkins CD Pipeline")
	if err := h.setStatus(jj, consts.StatusInProgress, jenkinsApi.CreateJenkinsPipeline, nil); err != nil {
		return errors.Wrap(err, "set status err")
	}
	if err := h.tryToCreateJob(jj); err != nil {
		if err := h.setStatus(jj, consts.StatusFailed, jenkinsApi.CreateJenkinsPipeline, &err); err != nil {
			return err
		}
		return err
	}
	if err := h.setStatus(jj, consts.StatusFinished, jenkinsApi.CreateJenkinsPipeline, nil); err != nil {
		return err
	}
	h.log.Info("end creating Jenkins CD Pipeline")
	return nextServeOrNil(h.next, jj)
}

func (h PutJenkinsPipeline) tryToCreateJob(jj *jenkinsApi.JenkinsJob) error {
	jc, err := h.initGoJenkinsClient(jj)
	if err != nil {
		return err
	}

	s, err := plutil.GetStageInstanceOwner(h.client, *jj)
	if err != nil {
		return err
	}

	json, err := h.ps.CreateStageJSON(*s)
	if err != nil {
		return err
	}

	conf, err := h.createStageConfig(s, json, jj.Spec.Job.Config)
	if err != nil {
		return err
	}

	if err := h.createJob(jc, conf, jj); err != nil {
		return errors.Wrap(err, "couldn't create jenkins job")
	}

	h.log.Info("job has been created", "name", jj.Spec.Job.Name)
	return nil
}

func (h PutJenkinsPipeline) createJob(jc *jenkinsClient.JenkinsClient, conf *string, jj *jenkinsApi.JenkinsJob) error {
	if jj.Spec.JenkinsFolder != nil && *jj.Spec.JenkinsFolder != "" {
		pfn := fmt.Sprintf("%v-%v", *jj.Spec.JenkinsFolder, "cd-pipeline")
		_, err := jc.GoJenkins.CreateJobInFolder(*conf, jj.Spec.Job.Name, pfn)
		if err != nil {
			return err
		}
		h.log.Info("job has been created",
			"name", fmt.Sprintf("%v/%v", pfn, jj.Spec.Job.Name))
		return nil
	}
	if _, err := jc.GoJenkins.CreateJob(*conf, jj.Spec.Job.Name); err != nil {
		return err
	}
	h.log.Info("job has been created", "name", jj.Spec.Job.Name)
	return nil
}

func (h PutJenkinsPipeline) initGoJenkinsClient(jj *jenkinsApi.JenkinsJob) (*jenkinsClient.JenkinsClient, error) {
	j, err := plutil.GetJenkinsInstanceOwner(h.client, jj.Name, jj.Namespace, jj.Spec.OwnerName, jj.GetOwnerReferences())
	if err != nil {
		return nil, errors.Wrapf(err, "an error has been occurred while getting owner jenkins for jenkins job %v", jj.Name)
	}
	h.log.Info("Jenkins instance has been created", "name", j.Name)
	return jenkinsClient.InitGoJenkinsClient(j, h.ps)
}

func (h PutJenkinsPipeline) createStageConfig(s *cdPipeApi.Stage, ps, conf string) (*string, error) {
	pipeSrc := map[string]interface{}{
		"type":    "default",
		"library": map[string]string{},
	}

	if s.Spec.Source.Type == consts.LibraryCodebase {
		h.setPipeSrcParams(s, pipeSrc)
	}

	tmpl, err := template.New("cd-pipeline.tmpl").Parse(conf)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"name":               s.Spec.Name,
		"gitServerCrVersion": "v2",
		"pipelineStages":     ps,
		"source":             pipeSrc,
	}
	var cdPipelineBuffer bytes.Buffer
	if err := tmpl.Execute(&cdPipelineBuffer, params); err != nil {
		return nil, err
	}
	pipeConf := cdPipelineBuffer.String()
	return &pipeConf, nil
}

func (h PutJenkinsPipeline) setPipeSrcParams(stage *cdPipeApi.Stage, pipeSrc map[string]interface{}) {
	cb, err := h.getLibraryParams(stage.Spec.Source.Library.Name, stage.Namespace)
	if err != nil {
		h.log.Error(err, "couldn't retrieve parameters for pipeline's library, default source type will be used",
			"Library name", stage.Spec.Source.Library.Name)
		return
	}
	gs, err := h.getGitServerParams(cb.Spec.GitServer, stage.Namespace)
	if err != nil {
		h.log.Error(err, "couldn't retrieve parameters for git server, default source type will be used",
			"Git server", cb.Spec.GitServer)
		return
	} else {
		pipeSrc["type"] = "library"
		pipeSrc["library"] = map[string]string{
			"url": fmt.Sprintf("ssh://%v@%v:%v%v", gs.Spec.GitUser, gs.Spec.GitHost, gs.Spec.SshPort,
				getPathToRepository(string(cb.Spec.Strategy), stage.Spec.Source.Library.Name, cb.Spec.GitUrlPath)),
			"credentials": gs.Spec.NameSshKeySecret,
			"branch":      stage.Spec.Source.Library.Branch,
		}
	}
}

func (h PutJenkinsPipeline) getLibraryParams(name, ns string) (*codebaseApi.Codebase, error) {
	nsn := types.NamespacedName{
		Namespace: ns,
		Name:      name,
	}
	i := &codebaseApi.Codebase{}
	if err := h.client.Get(context.TODO(), nsn, i); err != nil {
		return nil, err
	}
	return i, nil
}

func (h PutJenkinsPipeline) getGitServerParams(name string, ns string) (*codebaseApi.GitServer, error) {
	nsn := types.NamespacedName{
		Namespace: ns,
		Name:      name,
	}
	i := &codebaseApi.GitServer{}
	if err := h.client.Get(context.TODO(), nsn, i); err != nil {
		return nil, err
	}
	return i, nil
}

func getPathToRepository(strategy, name string, url *string) string {
	if strategy == consts.ImportStrategy {
		return *url
	}
	return "/" + name
}

func (h PutJenkinsPipeline) setStatus(jj *jenkinsApi.JenkinsJob, status string, action jenkinsApi.ActionType, err *error) error {
	jj.Status = jenkinsApi.JenkinsJobStatus{
		Status:          status,
		Available:       status == consts.StatusFinished,
		LastTimeUpdated: metav1.NewTime(time.Now()),
		Action:          action,
		Result:          getResult(status),
		Username:        "system",
		Value:           getValue(status),
	}

	if err != nil {
		errV := *err
		jj.Status.DetailedMessage = errV.Error()
	}

	return updateStatus(h.client, jj)
}

func getResult(status string) jenkinsApi.Result {
	if status == consts.StatusFinished {
		return jenkinsApi.Success
	} else if status == consts.StatusFailed {
		return jenkinsApi.Error
	}
	return "success"
}

func getValue(status string) string {
	if status == consts.StatusFinished {
		return "active"
	} else if status == consts.StatusFailed {
		return "failed"
	}
	return "inactive"
}

func updateStatus(c client.Client, jj *jenkinsApi.JenkinsJob) error {
	if err := c.Status().Update(context.TODO(), jj); err != nil {
		if err := c.Update(context.TODO(), jj); err != nil {
			return errors.Wrap(err, "couldn't update jenkins job status")
		}
	}
	log.Info("JenkinsJob status has been updated", "name", jj.Name)
	return nil
}
