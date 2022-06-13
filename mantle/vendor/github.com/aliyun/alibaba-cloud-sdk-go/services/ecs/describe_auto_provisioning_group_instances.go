package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeAutoProvisioningGroupInstances invokes the ecs.DescribeAutoProvisioningGroupInstances API synchronously
func (client *Client) DescribeAutoProvisioningGroupInstances(request *DescribeAutoProvisioningGroupInstancesRequest) (response *DescribeAutoProvisioningGroupInstancesResponse, err error) {
	response = CreateDescribeAutoProvisioningGroupInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoProvisioningGroupInstancesWithChan invokes the ecs.DescribeAutoProvisioningGroupInstances API asynchronously
func (client *Client) DescribeAutoProvisioningGroupInstancesWithChan(request *DescribeAutoProvisioningGroupInstancesRequest) (<-chan *DescribeAutoProvisioningGroupInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoProvisioningGroupInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoProvisioningGroupInstances(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeAutoProvisioningGroupInstancesWithCallback invokes the ecs.DescribeAutoProvisioningGroupInstances API asynchronously
func (client *Client) DescribeAutoProvisioningGroupInstancesWithCallback(request *DescribeAutoProvisioningGroupInstancesRequest, callback func(response *DescribeAutoProvisioningGroupInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoProvisioningGroupInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoProvisioningGroupInstances(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeAutoProvisioningGroupInstancesRequest is the request struct for api DescribeAutoProvisioningGroupInstances
type DescribeAutoProvisioningGroupInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber              requests.Integer `position:"Query" name:"PageNumber"`
	PageSize                requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	AutoProvisioningGroupId string           `position:"Query" name:"AutoProvisioningGroupId"`
}

// DescribeAutoProvisioningGroupInstancesResponse is the response struct for api DescribeAutoProvisioningGroupInstances
type DescribeAutoProvisioningGroupInstancesResponse struct {
	*responses.BaseResponse
	PageSize   int                                               `json:"PageSize" xml:"PageSize"`
	RequestId  string                                            `json:"RequestId" xml:"RequestId"`
	PageNumber int                                               `json:"PageNumber" xml:"PageNumber"`
	TotalCount int                                               `json:"TotalCount" xml:"TotalCount"`
	Instances  InstancesInDescribeAutoProvisioningGroupInstances `json:"Instances" xml:"Instances"`
}

// CreateDescribeAutoProvisioningGroupInstancesRequest creates a request to invoke DescribeAutoProvisioningGroupInstances API
func CreateDescribeAutoProvisioningGroupInstancesRequest() (request *DescribeAutoProvisioningGroupInstancesRequest) {
	request = &DescribeAutoProvisioningGroupInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoProvisioningGroupInstances", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutoProvisioningGroupInstancesResponse creates a response to parse from DescribeAutoProvisioningGroupInstances response
func CreateDescribeAutoProvisioningGroupInstancesResponse() (response *DescribeAutoProvisioningGroupInstancesResponse) {
	response = &DescribeAutoProvisioningGroupInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
