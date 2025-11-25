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

func RatePlanInstSmimea (client *esa20240910.Client) (_result *esa20240910.PurchaseRatePlanResponseBody, _err error) {
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

func SiteSmimea (ratePlanInstSmimeaResponseBody *esa20240910.PurchaseRatePlanResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateSiteResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call CreateSite to create resource")
  createSiteRequest := &esa20240910.CreateSiteRequest{
    SiteName: dara.String("gositecdn.cn"),
    InstanceId: ratePlanInstSmimeaResponseBody.InstanceId,
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

func RecordSmimea (siteSmimeaResponseBody *esa20240910.CreateSiteResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateRecordResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call CreateRecord to create resource")
  data := &esa20240910.CreateRecordRequestData{
    Usage: dara.Int32(1),
    MatchingType: dara.Int32(1),
    Certificate: dara.String("7777276264696475536f6d313237"),
    Selector: dara.Int32(1),
  }
  createRecordRequest := &esa20240910.CreateRecordRequest{
    RecordName: dara.String("www.gositecdn.cn"),
    Comment: dara.String("This is a remark"),
    SiteId: siteSmimeaResponseBody.SiteId,
    Type: dara.String("SMIMEA"),
    Data: data,
    Ttl: dara.Int32(100),
  }
  createRecordResponse, _err := CreateRecordWithRetry(client, createRecordRequest)
  if _err != nil {
    return _result, _err
  }

  fmt.Printf("[INFO] %s\n", "Call CreateRecord success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(createRecordResponse))
  _result = createRecordResponse.Body
  return _result , _err
}

func CreateRecordWithRetry (client *esa20240910.Client, createRecordRequest *esa20240910.CreateRecordRequest) (_result *esa20240910.CreateRecordResponse, _err error) {
  errorCode := ""
  retry1 := 0
  interval1 := 5000
  retry2 := 0
  interval2 := 5000
  for (retry1 < 10) || (retry2 < 20) {
    createRecordResponse, _err := client.CreateRecord(createRecordRequest)
    if _err != nil {
      if _t, ok := _err.(*dara.SDKError); ok {
        error := _t;
        errorCode = dara.StringValue(error.Code)
      }
    }
    fmt.Printf("[INFO] %s\n", "Call CreateRecord success, response: ")
    fmt.Printf("[INFO] %s\n", util.ToJSONString(createRecordResponse))
    _result = createRecordResponse
    return _result , _err
    if errorCode == "Site.ServiceBusy" {
      fmt.Printf("[INFO] %s\n", "Call CreateRecord failed, errorCode: Site.ServiceBusy, please retry")
      time.Sleep(time.Duration(interval1) * time.Millisecond)
      retry1++
    }

    if errorCode == "TooManyRequests" {
      fmt.Printf("[INFO] %s\n", "Call CreateRecord failed, errorCode: TooManyRequests, please retry")
      time.Sleep(time.Duration(interval2) * time.Millisecond)
      retry2++
    }

  }
  _err = dara.NewSDKError(map[string]interface{}{
    "message": "Call CreateRecord failed",
  })
  return _result, _err
}

func UpdateRecordSmimea (createRecordResponseBody *esa20240910.CreateRecordResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call UpdateRecord to update resource")
  data := &esa20240910.UpdateRecordRequestData{
    Usage: dara.Int32(3),
    MatchingType: dara.Int32(3),
    Certificate: dara.String("7737246264656475536f6d617256"),
    Selector: dara.Int32(3),
  }
  updateRecordRequest := &esa20240910.UpdateRecordRequest{
    Comment: dara.String("test_record_comment"),
    Data: data,
    Ttl: dara.Int32(86400),
    RecordId: createRecordResponseBody.RecordId,
  }
  updateRecordResponse, _err := UpdateRecordWithRetry(client, updateRecordRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call UpdateRecord success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(updateRecordResponse))
  return _err
}

func UpdateRecordWithRetry (client *esa20240910.Client, updateRecordRequest *esa20240910.UpdateRecordRequest) (_result *esa20240910.UpdateRecordResponse, _err error) {
  errorCode := ""
  retry1 := 0
  interval1 := 5000
  retry2 := 0
  interval2 := 3000
  for (retry1 < 20) || (retry2 < 10) {
    updateRecordResponse, _err := client.UpdateRecord(updateRecordRequest)
    if _err != nil {
      if _t, ok := _err.(*dara.SDKError); ok {
        error := _t;
        errorCode = dara.StringValue(error.Code)
      }
    }
    fmt.Printf("[INFO] %s\n", "Call UpdateRecord success, response: ")
    fmt.Printf("[INFO] %s\n", util.ToJSONString(updateRecordResponse))
    _result = updateRecordResponse
    return _result , _err
    if errorCode == "TooManyRequests" {
      fmt.Printf("[INFO] %s\n", "Call UpdateRecord failed, errorCode: TooManyRequests, please retry")
      time.Sleep(time.Duration(interval1) * time.Millisecond)
      retry1++
    }

    if errorCode == "Record.ServiceBusy" {
      fmt.Printf("[INFO] %s\n", "Call UpdateRecord failed, errorCode: Record.ServiceBusy, please retry")
      time.Sleep(time.Duration(interval2) * time.Millisecond)
      retry2++
    }

  }
  _err = dara.NewSDKError(map[string]interface{}{
    "message": "Call UpdateRecord failed",
  })
  return _result, _err
}

func DestroyRecordSmimea (createRecordResponseBody *esa20240910.CreateRecordResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call DeleteRecord to destroy resource")
  deleteRecordRequest := &esa20240910.DeleteRecordRequest{
    RecordId: createRecordResponseBody.RecordId,
  }
  deleteRecordResponse, _err := DeleteRecordWithRetry(client, deleteRecordRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call DeleteRecord success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(deleteRecordResponse))
  return _err
}

func DeleteRecordWithRetry (client *esa20240910.Client, deleteRecordRequest *esa20240910.DeleteRecordRequest) (_result *esa20240910.DeleteRecordResponse, _err error) {
  errorCode := ""
  retry1 := 0
  interval1 := 5000
  retry2 := 0
  interval2 := 1000
  for (retry1 < 20) || (retry2 < 10) {
    deleteRecordResponse, _err := client.DeleteRecord(deleteRecordRequest)
    if _err != nil {
      if _t, ok := _err.(*dara.SDKError); ok {
        error := _t;
        errorCode = dara.StringValue(error.Code)
      }
    }
    fmt.Printf("[INFO] %s\n", "Call DeleteRecord success, response: ")
    fmt.Printf("[INFO] %s\n", util.ToJSONString(deleteRecordResponse))
    _result = deleteRecordResponse
    return _result , _err
    if errorCode == "TooManyRequests" {
      fmt.Printf("[INFO] %s\n", "Call DeleteRecord failed, errorCode: TooManyRequests, please retry")
      time.Sleep(time.Duration(interval1) * time.Millisecond)
      retry1++
    }

    if errorCode == "Record.ServiceBusy" {
      fmt.Printf("[INFO] %s\n", "Call DeleteRecord failed, errorCode: Record.ServiceBusy, please retry")
      time.Sleep(time.Duration(interval2) * time.Millisecond)
      retry2++
    }

  }
  _err = dara.NewSDKError(map[string]interface{}{
    "message": "Call DeleteRecord failed",
  })
  return _result, _err
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
  ratePlanInstSmimeaRespBody, _err := RatePlanInstSmimea(esa20240910Client)
  if _err != nil {
    return _err
  }

  siteSmimeaRespBody, _err := SiteSmimea(ratePlanInstSmimeaRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }

  recordSmimeaRespBody, _err := RecordSmimea(siteSmimeaRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }

  // update resource
  _err = UpdateRecordSmimea(recordSmimeaRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }
  // destroy resource
  _err = DestroyRecordSmimea(recordSmimeaRespBody, esa20240910Client)
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
