// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"fmt"
	"os"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	esa20240910 "github.com/alibabacloud-go/esa-20240910/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/dara"
	credential "github.com/aliyun/credentials-go/credentials"
)

// Description:
//
// Init Client
func CreateESA20240910Client() (_result *esa20240910.Client, _err error) {
	config := &openapi.Config{}
	config.Credential, _err = credential.NewCredential(nil)
	if _err != nil {
		return _result, _err
	}

	// Endpoint please refer to https://api.aliyun.com/product/ESA
	config.Endpoint = dara.String("esa.cn-hangzhou.aliyuncs.com")
	_result, _err = esa20240910.NewClient(config)
	return _result, _err
}

func RatePlanInst(client *esa20240910.Client) (_result *esa20240910.PurchaseRatePlanResponseBody, _err error) {
	fmt.Printf("[INFO] %s\n", "Begin Call PurchaseRatePlan to create resource")
	purchaseRatePlanRequest := &esa20240910.PurchaseRatePlanRequest{
		Type:       dara.String("NS"),
		ChargeType: dara.String("PREPAY"),
		AutoRenew:  dara.Bool(false),
		Period:     dara.Int32(1),
		Coverage:   dara.String("overseas"),
		AutoPay:    dara.Bool(true),
		PlanName:   dara.String("high"),
	}
	purchaseRatePlanResponse, _err := client.PurchaseRatePlan(purchaseRatePlanRequest)
	if _err != nil {
		return _result, _err
	}

	describeRatePlanInstanceStatusRequest := &esa20240910.DescribeRatePlanInstanceStatusRequest{
		InstanceId: purchaseRatePlanResponse.Body.InstanceId,
	}
	currentRetry := 0
	delayedTime := 10000
	interval := 10000
	for currentRetry < 10 {
		sleepTime := 0
		if currentRetry == 0 {
			sleepTime = delayedTime
		} else {
			sleepTime = interval
		}

		fmt.Printf("[INFO] %s\n", "Polling for asynchronous results...")
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		describeRatePlanInstanceStatusResponse, _err := client.DescribeRatePlanInstanceStatus(describeRatePlanInstanceStatusRequest)
		if _err != nil {
			if _t, ok := _err.(*dara.SDKError); ok {
				error := _t
				_err = dara.NewSDKError(map[string]interface{}{
					"message": dara.StringValue(error.Message),
				})
				return _result, _err
			}
		}
		instanceStatus := dara.StringValue(describeRatePlanInstanceStatusResponse.Body.InstanceStatus)
		if instanceStatus == "running" {
			fmt.Printf("[INFO] %s\n", "Call PurchaseRatePlan success, response: ")
			fmt.Printf("[INFO] %s\n", util.ToJSONString(purchaseRatePlanResponse))
			_result = purchaseRatePlanResponse.Body
			return _result, _err
		}

		currentRetry++
	}
	_err = dara.NewSDKError(map[string]interface{}{
		"message": "Asynchronous check failed",
	})
	return _result, _err
}

func Site(ratePlanInstResponseBody *esa20240910.PurchaseRatePlanResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateSiteResponseBody, _err error) {
	fmt.Printf("[INFO] %s\n", "Begin Call CreateSite to create resource")
	createSiteRequest := &esa20240910.CreateSiteRequest{
		SiteName:   dara.String("gositecdn.cn"),
		InstanceId: ratePlanInstResponseBody.InstanceId,
		Coverage:   dara.String("overseas"),
		AccessType: dara.String("NS"),
	}
	createSiteResponse, _err := client.CreateSite(createSiteRequest)
	if _err != nil {
		return _result, _err
	}

	getSiteRequest := &esa20240910.GetSiteRequest{
		SiteId: createSiteResponse.Body.SiteId,
	}
	currentRetry := 0
	delayedTime := 60000
	interval := 10000
	for currentRetry < 5 {
		sleepTime := 0
		if currentRetry == 0 {
			sleepTime = delayedTime
		} else {
			sleepTime = interval
		}

		fmt.Printf("[INFO] %s\n", "Polling for asynchronous results...")
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		getSiteResponse, _err := client.GetSite(getSiteRequest)
		if _err != nil {
			if _t, ok := _err.(*dara.SDKError); ok {
				error := _t
				_err = dara.NewSDKError(map[string]interface{}{
					"message": dara.StringValue(error.Message),
				})
				return _result, _err
			}
		}
		status := dara.StringValue(getSiteResponse.Body.SiteModel.Status)
		if status == "pending" {
			fmt.Printf("[INFO] %s\n", "Call CreateSite success, response: ")
			fmt.Printf("[INFO] %s\n", util.ToJSONString(createSiteResponse))
			_result = createSiteResponse.Body
			return _result, _err
		}

		currentRetry++
	}
	_err = dara.NewSDKError(map[string]interface{}{
		"message": "Asynchronous check failed",
	})
	return _result, _err
}

func ReqHdrModRule(siteResponseBody *esa20240910.CreateSiteResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateHttpRequestHeaderModificationRuleResponseBody, _err error) {
	fmt.Printf("[INFO] %s\n", "Begin Call CreateHttpRequestHeaderModificationRule to create resource")
	requestHeaderModification := &esa20240910.CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification{
		Type:      dara.String("static"),
		Value:     dara.String("add"),
		Operation: dara.String("add"),
		Name:      dara.String("testadd"),
	}
	requestHeaderModification1 := &esa20240910.CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification{
		Operation: dara.String("del"),
		Name:      dara.String("testdel"),
	}
	requestHeaderModification2 := &esa20240910.CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification{
		Type:      dara.String("dynamic"),
		Value:     dara.String("ip.geoip.country"),
		Operation: dara.String("modify"),
		Name:      dara.String("testmodify"),
	}
	createHttpRequestHeaderModificationRuleRequest := &esa20240910.CreateHttpRequestHeaderModificationRuleRequest{
		SiteId:                    siteResponseBody.SiteId,
		RuleEnable:                dara.String("on"),
		Rule:                      dara.String("(http.host eq \"video.example.com\")"),
		Sequence:                  dara.Int32(1),
		SiteVersion:               dara.Int32(0),
		RuleName:                  dara.String("test"),
		RequestHeaderModification: []*esa20240910.CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification{requestHeaderModification, requestHeaderModification1, requestHeaderModification2},
	}
	createHttpRequestHeaderModificationRuleResponse, _err := client.CreateHttpRequestHeaderModificationRule(createHttpRequestHeaderModificationRuleRequest)
	if _err != nil {
		return _result, _err
	}

	fmt.Printf("[INFO] %s\n", "Call CreateHttpRequestHeaderModificationRule success, response: ")
	fmt.Printf("[INFO] %s\n", util.ToJSONString(createHttpRequestHeaderModificationRuleResponse))
	_result = createHttpRequestHeaderModificationRuleResponse.Body
	return _result, _err
}

func UpdateReqHdrModRule(siteResponseBody *esa20240910.CreateSiteResponseBody, createHttpRequestHeaderModificationRuleResponseBody *esa20240910.CreateHttpRequestHeaderModificationRuleResponseBody, client *esa20240910.Client) (_err error) {
	fmt.Printf("[INFO] %s\n", "Begin Call UpdateHttpRequestHeaderModificationRule to update resource")
	requestHeaderModification := &esa20240910.UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification{
		Type:      dara.String("static"),
		Value:     dara.String("modify1"),
		Operation: dara.String("modify"),
		Name:      dara.String("testmodify1"),
	}
	updateHttpRequestHeaderModificationRuleRequest := &esa20240910.UpdateHttpRequestHeaderModificationRuleRequest{
		SiteId:                    siteResponseBody.SiteId,
		RuleEnable:                dara.String("off"),
		Rule:                      dara.String("(http.request.uri eq \"/content?page=1234\")"),
		Sequence:                  dara.Int32(1),
		RuleName:                  dara.String("test_modify"),
		RequestHeaderModification: []*esa20240910.UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification{requestHeaderModification},
		ConfigId:                  createHttpRequestHeaderModificationRuleResponseBody.ConfigId,
	}
	updateHttpRequestHeaderModificationRuleResponse, _err := client.UpdateHttpRequestHeaderModificationRule(updateHttpRequestHeaderModificationRuleRequest)
	if _err != nil {
		return _err
	}

	fmt.Printf("[INFO] %s\n", "Call UpdateHttpRequestHeaderModificationRule success, response: ")
	fmt.Printf("[INFO] %s\n", util.ToJSONString(updateHttpRequestHeaderModificationRuleResponse))
	return _err
}

func DestroyReqHdrModRule(siteResponseBody *esa20240910.CreateSiteResponseBody, createHttpRequestHeaderModificationRuleResponseBody *esa20240910.CreateHttpRequestHeaderModificationRuleResponseBody, client *esa20240910.Client) (_err error) {
	fmt.Printf("[INFO] %s\n", "Begin Call DeleteHttpRequestHeaderModificationRule to destroy resource")
	deleteHttpRequestHeaderModificationRuleRequest := &esa20240910.DeleteHttpRequestHeaderModificationRuleRequest{
		SiteId:   siteResponseBody.SiteId,
		ConfigId: createHttpRequestHeaderModificationRuleResponseBody.ConfigId,
	}
	deleteHttpRequestHeaderModificationRuleResponse, _err := client.DeleteHttpRequestHeaderModificationRule(deleteHttpRequestHeaderModificationRuleRequest)
	if _err != nil {
		return _err
	}

	fmt.Printf("[INFO] %s\n", "Call DeleteHttpRequestHeaderModificationRule success, response: ")
	fmt.Printf("[INFO] %s\n", util.ToJSONString(deleteHttpRequestHeaderModificationRuleResponse))
	return _err
}

// Description:
//
// Running code may affect the online resources of the current account, please proceed with caution!
func _main(args []*string) (_err error) {
	// The code may contain api calls involving fees. Please ensure that you fully understand the charging methods and prices before running.
	// Set the environment variable COST_ACK to true or delete the following judgment to run the sample code.
	costAcknowledged := os.Getenv("COST_ACK")
	if dara.IsNil(dara.String(costAcknowledged)) || !(costAcknowledged == "true") {
		fmt.Printf("[WARNING] %s\n", "Running code may affect the online resources of the current account, please proceed with caution!")
		return
	}

	// Init client
	esa20240910Client, _err := CreateESA20240910Client()
	if _err != nil {
		return _err
	}

	// Init resource
	ratePlanInstRespBody, _err := RatePlanInst(esa20240910Client)
	if _err != nil {
		return _err
	}

	siteRespBody, _err := Site(ratePlanInstRespBody, esa20240910Client)
	if _err != nil {
		return _err
	}

	reqHdrModRuleRespBody, _err := ReqHdrModRule(siteRespBody, esa20240910Client)
	if _err != nil {
		return _err
	}

	// update resource
	_err = UpdateReqHdrModRule(siteRespBody, reqHdrModRuleRespBody, esa20240910Client)
	if _err != nil {
		return _err
	}
	// destroy resource
	_err = DestroyReqHdrModRule(siteRespBody, reqHdrModRuleRespBody, esa20240910Client)
	if _err != nil {
		return _err
	}
	return _err
}

func main() {
	err := _main(dara.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
