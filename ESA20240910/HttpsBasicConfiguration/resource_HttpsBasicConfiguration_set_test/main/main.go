// This file is auto-generated, don't edit it. Thanks.
package main

import (
  "os"
  esa20240910 "github.com/alibabacloud-go/esa-20240910/client"
  openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  credential "github.com/aliyun/credentials-go/credentials"
  util "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
  "time"
)


// Description:
// 
// Init Client
func CreateESA20240910Client () (_result *esa20240910.Client, _err error) {
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

func RatePlanInst (client *esa20240910.Client) (_result *esa20240910.PurchaseRatePlanResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call PurchaseRatePlan to create resource")
  purchaseRatePlanRequest := &esa20240910.PurchaseRatePlanRequest{
    Type: dara.String("NS"),
    ChargeType: dara.String("PREPAY"),
    AutoRenew: dara.Bool(false),
    Period: dara.Int32(1),
    Coverage: dara.String("overseas"),
    AutoPay: dara.Bool(true),
    PlanName: dara.String("high"),
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
        error := _t;
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
      return _result , _err
    }

    currentRetry++
  }
  _err = dara.NewSDKError(map[string]interface{}{
    "message": "Asynchronous check failed",
  })
  return _result, _err
}

func Site (ratePlanInstResponseBody *esa20240910.PurchaseRatePlanResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateSiteResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call CreateSite to create resource")
  createSiteRequest := &esa20240910.CreateSiteRequest{
    SiteName: dara.String("gositecdn.cn"),
    InstanceId: ratePlanInstResponseBody.InstanceId,
    Coverage: dara.String("overseas"),
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
        error := _t;
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
      return _result , _err
    }

    currentRetry++
  }
  _err = dara.NewSDKError(map[string]interface{}{
    "message": "Asynchronous check failed",
  })
  return _result, _err
}

func HttpsCfg (siteResponseBody *esa20240910.CreateSiteResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateHttpsBasicConfigurationResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call CreateHttpsBasicConfiguration to create resource")
  createHttpsBasicConfigurationRequest := &esa20240910.CreateHttpsBasicConfigurationRequest{
    SiteId: siteResponseBody.SiteId,
    RuleEnable: dara.String("on"),
    Https: dara.String("on"),
    Rule: dara.String("true"),
    RuleName: dara.String("test2"),
  }
  createHttpsBasicConfigurationResponse, _err := client.CreateHttpsBasicConfiguration(createHttpsBasicConfigurationRequest)
  if _err != nil {
    return _result, _err
  }

  fmt.Printf("[INFO] %s\n", "Call CreateHttpsBasicConfiguration success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(createHttpsBasicConfigurationResponse))
  _result = createHttpsBasicConfigurationResponse.Body
  return _result , _err
}

func UpdateHttpsCfg (siteResponseBody *esa20240910.CreateSiteResponseBody, createHttpsBasicConfigurationResponseBody *esa20240910.CreateHttpsBasicConfigurationResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call UpdateHttpsBasicConfiguration to update resource")
  updateHttpsBasicConfigurationRequest := &esa20240910.UpdateHttpsBasicConfigurationRequest{
    SiteId: siteResponseBody.SiteId,
    RuleEnable: dara.String("on"),
    Https: dara.String("off"),
    Rule: dara.String("true"),
    RuleName: dara.String("test2"),
    ConfigId: createHttpsBasicConfigurationResponseBody.ConfigId,
  }
  updateHttpsBasicConfigurationResponse, _err := client.UpdateHttpsBasicConfiguration(updateHttpsBasicConfigurationRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call UpdateHttpsBasicConfiguration success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(updateHttpsBasicConfigurationResponse))
  return _err
}

func DestroyHttpsCfg (siteResponseBody *esa20240910.CreateSiteResponseBody, createHttpsBasicConfigurationResponseBody *esa20240910.CreateHttpsBasicConfigurationResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call DeleteHttpsBasicConfiguration to destroy resource")
  deleteHttpsBasicConfigurationRequest := &esa20240910.DeleteHttpsBasicConfigurationRequest{
    SiteId: siteResponseBody.SiteId,
    ConfigId: createHttpsBasicConfigurationResponseBody.ConfigId,
  }
  deleteHttpsBasicConfigurationResponse, _err := client.DeleteHttpsBasicConfiguration(deleteHttpsBasicConfigurationRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call DeleteHttpsBasicConfiguration success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(deleteHttpsBasicConfigurationResponse))
  return _err
}

// Description:
// 
// Running code may affect the online resources of the current account, please proceed with caution!
func _main (args []*string) (_err error) {
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
  //resource_RatePlanInstance_set_test
  ratePlanInstRespBody, _err := RatePlanInst(esa20240910Client)
  if _err != nil {
    return _err
  }

  //resource_HttpBasicConfiguration_set_test
  siteRespBody, _err := Site(ratePlanInstRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }

  httpsCfgRespBody, _err := HttpsCfg(siteRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }

  // update resource
  _err = UpdateHttpsCfg(siteRespBody, httpsCfgRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }
  // destroy resource
  _err = DestroyHttpsCfg(siteRespBody, httpsCfgRespBody, esa20240910Client)
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
