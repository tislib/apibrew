package test

import (
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/test/setup"
)

var authenticationClient stub.AuthenticationClient
var dataSourceClient stub.DataSourceClient
var resourceClient stub.ResourceClient
var recordClient stub.RecordClient
var userClient stub.UserClient
var namespaceClient stub.NamespaceClient

var container abs.Container

func init() {
	dhClient := setup.GetTestDhClient()
	recordClient = dhClient.GetRecordClient()
	authenticationClient = dhClient.GetAuthenticationClient()
	resourceClient = dhClient.GetResourceClient()
	dataSourceClient = dhClient.GetDataSourceClient()
	userClient = dhClient.GetUserClient()
	namespaceClient = dhClient.GetNamespaceClient()
	container = setup.GetContainer()
}
