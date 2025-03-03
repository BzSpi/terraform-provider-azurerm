package securityinsight

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ThreatIntelligenceIndicatorClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource
// provider
type ThreatIntelligenceIndicatorClient struct {
	BaseClient
}

// NewThreatIntelligenceIndicatorClient creates an instance of the ThreatIntelligenceIndicatorClient client.
func NewThreatIntelligenceIndicatorClient(subscriptionID string) ThreatIntelligenceIndicatorClient {
	return NewThreatIntelligenceIndicatorClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewThreatIntelligenceIndicatorClientWithBaseURI creates an instance of the ThreatIntelligenceIndicatorClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewThreatIntelligenceIndicatorClientWithBaseURI(baseURI string, subscriptionID string) ThreatIntelligenceIndicatorClient {
	return ThreatIntelligenceIndicatorClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AppendTags append tags to a threat intelligence indicator.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// name - threat intelligence indicator name field.
// threatIntelligenceAppendTags - the threat intelligence append tags request body
func (client ThreatIntelligenceIndicatorClient) AppendTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags ThreatIntelligenceAppendTags) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.AppendTags")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "AppendTags", err.Error())
	}

	req, err := client.AppendTagsPreparer(ctx, resourceGroupName, workspaceName, name, threatIntelligenceAppendTags)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "AppendTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.AppendTagsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "AppendTags", resp, "Failure sending request")
		return
	}

	result, err = client.AppendTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "AppendTags", resp, "Failure responding to request")
		return
	}

	return
}

// AppendTagsPreparer prepares the AppendTags request.
func (client ThreatIntelligenceIndicatorClient) AppendTagsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags ThreatIntelligenceAppendTags) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}/appendTags", pathParameters),
		autorest.WithJSON(threatIntelligenceAppendTags),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AppendTagsSender sends the AppendTags request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) AppendTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// AppendTagsResponder handles the response to the AppendTags request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) AppendTagsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create update a threat Intelligence indicator.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// name - threat intelligence indicator name field.
// threatIntelligenceProperties - properties of threat intelligence indicators to create and update.
func (client ThreatIntelligenceIndicatorClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel) (result ThreatIntelligenceInformationModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceName, name, threatIntelligenceProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ThreatIntelligenceIndicatorClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}", pathParameters),
		autorest.WithJSON(threatIntelligenceProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) CreateResponder(resp *http.Response) (result ThreatIntelligenceInformationModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateIndicator create a new threat intelligence indicator.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// threatIntelligenceProperties - properties of threat intelligence indicators to create and update.
func (client ThreatIntelligenceIndicatorClient) CreateIndicator(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel) (result ThreatIntelligenceInformationModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.CreateIndicator")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "CreateIndicator", err.Error())
	}

	req, err := client.CreateIndicatorPreparer(ctx, resourceGroupName, workspaceName, threatIntelligenceProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "CreateIndicator", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateIndicatorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "CreateIndicator", resp, "Failure sending request")
		return
	}

	result, err = client.CreateIndicatorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "CreateIndicator", resp, "Failure responding to request")
		return
	}

	return
}

// CreateIndicatorPreparer prepares the CreateIndicator request.
func (client ThreatIntelligenceIndicatorClient) CreateIndicatorPreparer(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/createIndicator", pathParameters),
		autorest.WithJSON(threatIntelligenceProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateIndicatorSender sends the CreateIndicator request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) CreateIndicatorSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateIndicatorResponder handles the response to the CreateIndicator request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) CreateIndicatorResponder(resp *http.Response) (result ThreatIntelligenceInformationModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a threat intelligence indicator.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// name - threat intelligence indicator name field.
func (client ThreatIntelligenceIndicatorClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ThreatIntelligenceIndicatorClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get view a threat intelligence indicator by name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// name - threat intelligence indicator name field.
func (client ThreatIntelligenceIndicatorClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string) (result ThreatIntelligenceInformationModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ThreatIntelligenceIndicatorClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) GetResponder(resp *http.Response) (result ThreatIntelligenceInformationModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// QueryIndicators query threat intelligence indicators as per filtering criteria.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// threatIntelligenceFilteringCriteria - filtering criteria for querying threat intelligence indicators.
func (client ThreatIntelligenceIndicatorClient) QueryIndicators(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria) (result ThreatIntelligenceInformationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.QueryIndicators")
		defer func() {
			sc := -1
			if result.tiil.Response.Response != nil {
				sc = result.tiil.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "QueryIndicators", err.Error())
	}

	result.fn = client.queryIndicatorsNextResults
	req, err := client.QueryIndicatorsPreparer(ctx, resourceGroupName, workspaceName, threatIntelligenceFilteringCriteria)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "QueryIndicators", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryIndicatorsSender(req)
	if err != nil {
		result.tiil.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "QueryIndicators", resp, "Failure sending request")
		return
	}

	result.tiil, err = client.QueryIndicatorsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "QueryIndicators", resp, "Failure responding to request")
		return
	}
	if result.tiil.hasNextLink() && result.tiil.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// QueryIndicatorsPreparer prepares the QueryIndicators request.
func (client ThreatIntelligenceIndicatorClient) QueryIndicatorsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/queryIndicators", pathParameters),
		autorest.WithJSON(threatIntelligenceFilteringCriteria),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryIndicatorsSender sends the QueryIndicators request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) QueryIndicatorsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// QueryIndicatorsResponder handles the response to the QueryIndicators request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) QueryIndicatorsResponder(resp *http.Response) (result ThreatIntelligenceInformationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// queryIndicatorsNextResults retrieves the next set of results, if any.
func (client ThreatIntelligenceIndicatorClient) queryIndicatorsNextResults(ctx context.Context, lastResults ThreatIntelligenceInformationList) (result ThreatIntelligenceInformationList, err error) {
	req, err := lastResults.threatIntelligenceInformationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "queryIndicatorsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.QueryIndicatorsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "queryIndicatorsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.QueryIndicatorsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "queryIndicatorsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// QueryIndicatorsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ThreatIntelligenceIndicatorClient) QueryIndicatorsComplete(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria) (result ThreatIntelligenceInformationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.QueryIndicators")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.QueryIndicators(ctx, resourceGroupName, workspaceName, threatIntelligenceFilteringCriteria)
	return
}

// ReplaceTags replace tags added to a threat intelligence indicator.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// name - threat intelligence indicator name field.
// threatIntelligenceReplaceTags - tags in the threat intelligence indicator to be replaced.
func (client ThreatIntelligenceIndicatorClient) ReplaceTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags ThreatIntelligenceIndicatorModel) (result ThreatIntelligenceInformationModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ThreatIntelligenceIndicatorClient.ReplaceTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorClient", "ReplaceTags", err.Error())
	}

	req, err := client.ReplaceTagsPreparer(ctx, resourceGroupName, workspaceName, name, threatIntelligenceReplaceTags)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "ReplaceTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReplaceTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "ReplaceTags", resp, "Failure sending request")
		return
	}

	result, err = client.ReplaceTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorClient", "ReplaceTags", resp, "Failure responding to request")
		return
	}

	return
}

// ReplaceTagsPreparer prepares the ReplaceTags request.
func (client ThreatIntelligenceIndicatorClient) ReplaceTagsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags ThreatIntelligenceIndicatorModel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}/replaceTags", pathParameters),
		autorest.WithJSON(threatIntelligenceReplaceTags),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReplaceTagsSender sends the ReplaceTags request. The method will close the
// http.Response Body if it receives an error.
func (client ThreatIntelligenceIndicatorClient) ReplaceTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ReplaceTagsResponder handles the response to the ReplaceTags request. The method always
// closes the http.Response Body.
func (client ThreatIntelligenceIndicatorClient) ReplaceTagsResponder(resp *http.Response) (result ThreatIntelligenceInformationModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
