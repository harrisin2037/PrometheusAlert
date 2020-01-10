package green

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

// DescribeAuditContentItem invokes the green.DescribeAuditContentItem API synchronously
// api document: https://help.aliyun.com/api/green/describeauditcontentitem.html
func (client *Client) DescribeAuditContentItem(request *DescribeAuditContentItemRequest) (response *DescribeAuditContentItemResponse, err error) {
	response = CreateDescribeAuditContentItemResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAuditContentItemWithChan invokes the green.DescribeAuditContentItem API asynchronously
// api document: https://help.aliyun.com/api/green/describeauditcontentitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuditContentItemWithChan(request *DescribeAuditContentItemRequest) (<-chan *DescribeAuditContentItemResponse, <-chan error) {
	responseChan := make(chan *DescribeAuditContentItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAuditContentItem(request)
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

// DescribeAuditContentItemWithCallback invokes the green.DescribeAuditContentItem API asynchronously
// api document: https://help.aliyun.com/api/green/describeauditcontentitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuditContentItemWithCallback(request *DescribeAuditContentItemRequest, callback func(response *DescribeAuditContentItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAuditContentItemResponse
		var err error
		defer close(result)
		response, err = client.DescribeAuditContentItem(request)
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

// DescribeAuditContentItemRequest is the request struct for api DescribeAuditContentItem
type DescribeAuditContentItemRequest struct {
	*requests.RpcRequest
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	TaskId       string           `position:"Query" name:"TaskId"`
	TotalCount   requests.Integer `position:"Query" name:"TotalCount"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	ResourceType string           `position:"Query" name:"ResourceType"`
}

// DescribeAuditContentItemResponse is the response struct for api DescribeAuditContentItem
type DescribeAuditContentItemResponse struct {
	*responses.BaseResponse
	RequestId            string             `json:"RequestId" xml:"RequestId"`
	PageSize             int                `json:"PageSize" xml:"PageSize"`
	CurrentPage          int                `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount           int                `json:"TotalCount" xml:"TotalCount"`
	AuditContentItemList []AuditContentItem `json:"AuditContentItemList" xml:"AuditContentItemList"`
}

// CreateDescribeAuditContentItemRequest creates a request to invoke DescribeAuditContentItem API
func CreateDescribeAuditContentItemRequest() (request *DescribeAuditContentItemRequest) {
	request = &DescribeAuditContentItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeAuditContentItem", "green", "openAPI")
	return
}

// CreateDescribeAuditContentItemResponse creates a response to parse from DescribeAuditContentItem response
func CreateDescribeAuditContentItemResponse() (response *DescribeAuditContentItemResponse) {
	response = &DescribeAuditContentItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}