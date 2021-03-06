package workflows

import (
	"fmt"
	"github.com/stelligent/mu/common"
)

// DatabaseSetPassword sets a database password for an environment
func DatabaseSetPassword(ctx *common.Context, environmentName string, serviceName string, newPassword string) Executor {
	workflow := new(databaseWorkflow)

	return newPipelineExecutor(
		workflow.databaseInput(ctx, serviceName),
		workflow.databaseSetPassword(ctx, environmentName, newPassword),
	)
}

// DatabaseSetPassword sets a database password for an environment
func (workflow *databaseWorkflow) databaseSetPassword(ctx *common.Context, environmentName string, newPassword string) Executor {
	return func() error {
		dbStackName := common.CreateStackName(ctx.Config.Namespace, common.StackTypeDatabase, workflow.serviceName, environmentName)
		return ctx.ParamManager.SetParam(fmt.Sprintf("%s-%s", dbStackName, "DatabaseMasterPassword"), newPassword)
	}
}

// DatabaseGetPassword gets a database password for an environment
func DatabaseGetPassword(ctx *common.Context, environmentName string, serviceName string) Executor {
	workflow := new(databaseWorkflow)

	return newPipelineExecutor(
		workflow.databaseInput(ctx, serviceName),
		workflow.databaseGetPassword(ctx, environmentName),
	)
}

func (workflow *databaseWorkflow) databaseGetPassword(ctx *common.Context, environmentName string) Executor {
	return func() error {
		dbStackName := common.CreateStackName(ctx.Config.Namespace, common.StackTypeDatabase, workflow.serviceName, environmentName)
		log.Debugf("Getting password for dbStackName:%s", dbStackName)
		dbPass, _ := ctx.ParamManager.GetParam(fmt.Sprintf("%s-%s", dbStackName, "DatabaseMasterPassword"))
		log.Noticef("%s", dbPass)
		return nil
	}
}
