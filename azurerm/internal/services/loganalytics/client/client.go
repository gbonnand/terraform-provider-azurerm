package client

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2020-03-01-preview/operationalinsights"
	"github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	DataExportClient           *operationalinsights.DataExportsClient
	DataSourcesClient          *operationalinsights.DataSourcesClient
	LinkedServicesClient       *operationalinsights.LinkedServicesClient
	SavedSearchesClient        *operationalinsights.SavedSearchesClient
	SharedKeysClient           *operationalinsights.SharedKeysClient
	SolutionsClient            *operationsmanagement.SolutionsClient
	StorageInsightConfigClient *operationalinsights.StorageInsightConfigsClient
	WorkspacesClient           *operationalinsights.WorkspacesClient
}

func NewClient(o *common.ClientOptions) *Client {
	DataExportClient := operationalinsights.NewDataExportsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DataExportClient.Client, o.ResourceManagerAuthorizer)

	DataSourcesClient := operationalinsights.NewDataSourcesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DataSourcesClient.Client, o.ResourceManagerAuthorizer)

	WorkspacesClient := operationalinsights.NewWorkspacesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&WorkspacesClient.Client, o.ResourceManagerAuthorizer)

	SavedSearchesClient := operationalinsights.NewSavedSearchesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&SavedSearchesClient.Client, o.ResourceManagerAuthorizer)

	SharedKeysClient := operationalinsights.NewSharedKeysClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&SharedKeysClient.Client, o.ResourceManagerAuthorizer)

	SolutionsClient := operationsmanagement.NewSolutionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId, "Microsoft.OperationsManagement", "solutions", "testing")
	o.ConfigureClient(&SolutionsClient.Client, o.ResourceManagerAuthorizer)

	StorageInsightConfigClient := operationalinsights.NewStorageInsightConfigsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&StorageInsightConfigClient.Client, o.ResourceManagerAuthorizer)

	LinkedServicesClient := operationalinsights.NewLinkedServicesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&LinkedServicesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		DataExportClient:           &DataExportClient,
		DataSourcesClient:          &DataSourcesClient,
		LinkedServicesClient:       &LinkedServicesClient,
		SavedSearchesClient:        &SavedSearchesClient,
		SharedKeysClient:           &SharedKeysClient,
		SolutionsClient:            &SolutionsClient,
		StorageInsightConfigClient: &StorageInsightConfigClient,
		WorkspacesClient:           &WorkspacesClient,
	}
}
