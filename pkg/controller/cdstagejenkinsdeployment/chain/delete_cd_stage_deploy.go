package chain

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	jenkinsApi "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1"
	"github.com/epam/edp-jenkins-operator/v2/pkg/controller/cdstagejenkinsdeployment/helper"
	"github.com/epam/edp-jenkins-operator/v2/pkg/util/consts"
)

type DeleteCDStageDeploy struct {
	client client.Client
	log    logr.Logger
}

func (h DeleteCDStageDeploy) ServeRequest(jenkinsDeploy *jenkinsApi.CDStageJenkinsDeployment) error {
	log := h.log.WithValues("name", jenkinsDeploy.Spec.Job)
	log.Info("deleting CDStageDeploy")

	if err := h.deleteCDStageDeploy(jenkinsDeploy); err != nil {
		return err
	}

	log.Info("CDStageDeploy has been deleted")
	return nil
}

func (h DeleteCDStageDeploy) deleteCDStageDeploy(jenkinsDeploy *jenkinsApi.CDStageJenkinsDeployment) error {
	s, err := helper.GetCDStageDeploy(h.client, jenkinsDeploy.Labels[consts.CdStageDeployKey], jenkinsDeploy.Namespace)
	if err != nil {
		return err
	}
	return h.client.Delete(context.TODO(), s)
}
