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

func RatePlanInstCltCert (client *esa20240910.Client) (_result *esa20240910.PurchaseRatePlanResponseBody, _err error) {
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

func SiteCltCert (ratePlanInstCltCertResponseBody *esa20240910.PurchaseRatePlanResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateSiteResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call CreateSite to create resource")
  createSiteRequest := &esa20240910.CreateSiteRequest{
    SiteName: dara.String("gositecdn.cn"),
    InstanceId: ratePlanInstCltCertResponseBody.InstanceId,
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

func CltCert (siteCltCertResponseBody *esa20240910.CreateSiteResponseBody, client *esa20240910.Client) (_result *esa20240910.CreateClientCertificateResponseBody, _err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call CreateClientCertificate to create resource")
  createClientCertificateRequest := &esa20240910.CreateClientCertificateRequest{
    SiteId: siteCltCertResponseBody.SiteId,
    PkeyType: dara.String("RSA"),
    ValidityDays: dara.Int64(365),
  }
  createClientCertificateResponse, _err := client.CreateClientCertificate(createClientCertificateRequest)
  if _err != nil {
    return _result, _err
  }

  fmt.Printf("[INFO] %s\n", "Call CreateClientCertificate success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(createClientCertificateResponse))
  _result = createClientCertificateResponse.Body
  return _result , _err
}

func UpdateCltCert (siteCltCertResponseBody *esa20240910.CreateSiteResponseBody, createClientCertificateResponseBody *esa20240910.CreateClientCertificateResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call RevokeClientCertificate to update resource")
  revokeClientCertificateRequest := &esa20240910.RevokeClientCertificateRequest{
    SiteId: siteCltCertResponseBody.SiteId,
    Id: createClientCertificateResponseBody.Id,
  }
  revokeClientCertificateResponse, _err := client.RevokeClientCertificate(revokeClientCertificateRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call RevokeClientCertificate success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(revokeClientCertificateResponse))
  return _err
}

func UpdateCltCert1 (siteCltCertResponseBody *esa20240910.CreateSiteResponseBody, createClientCertificateResponseBody *esa20240910.CreateClientCertificateResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call ActivateClientCertificate to update resource")
  activateClientCertificateRequest := &esa20240910.ActivateClientCertificateRequest{
    SiteId: siteCltCertResponseBody.SiteId,
    Id: createClientCertificateResponseBody.Id,
  }
  activateClientCertificateResponse, _err := client.ActivateClientCertificate(activateClientCertificateRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call ActivateClientCertificate success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(activateClientCertificateResponse))
  return _err
}

func UpdateCltCert2 (siteCltCertResponseBody *esa20240910.CreateSiteResponseBody, createClientCertificateResponseBody *esa20240910.CreateClientCertificateResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call RevokeClientCertificate to update resource")
  revokeClientCertificateRequest := &esa20240910.RevokeClientCertificateRequest{
    SiteId: siteCltCertResponseBody.SiteId,
    Id: createClientCertificateResponseBody.Id,
  }
  revokeClientCertificateResponse, _err := client.RevokeClientCertificate(revokeClientCertificateRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call RevokeClientCertificate success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(revokeClientCertificateResponse))
  return _err
}

func DestroyCltCert (siteCltCertResponseBody *esa20240910.CreateSiteResponseBody, createClientCertificateResponseBody *esa20240910.CreateClientCertificateResponseBody, client *esa20240910.Client) (_err error) {
  fmt.Printf("[INFO] %s\n", "Begin Call DeleteClientCertificate to destroy resource")
  deleteClientCertificateRequest := &esa20240910.DeleteClientCertificateRequest{
    SiteId: siteCltCertResponseBody.SiteId,
    Id: createClientCertificateResponseBody.Id,
  }
  deleteClientCertificateResponse, _err := client.DeleteClientCertificate(deleteClientCertificateRequest)
  if _err != nil {
    return _err
  }

  fmt.Printf("[INFO] %s\n", "Call DeleteClientCertificate success, response: ")
  fmt.Printf("[INFO] %s\n", util.ToJSONString(deleteClientCertificateResponse))
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
  //resource_RatePlanInstance_ClientCertificate_set_test
  ratePlanInstCltCertRespBody, _err := RatePlanInstCltCert(esa20240910Client)
  if _err != nil {
    return _err
  }

  //resource_Site_ClientCerticificate_set_test
  siteCltCertRespBody, _err := SiteCltCert(ratePlanInstCltCertRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }

  cltCertRespBody, _err := CltCert(siteCltCertRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }

  // update resource
  _err = UpdateCltCert(siteCltCertRespBody, cltCertRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }
  _err = UpdateCltCert1(siteCltCertRespBody, cltCertRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }
  _err = UpdateCltCert2(siteCltCertRespBody, cltCertRespBody, esa20240910Client)
  if _err != nil {
    return _err
  }
  // destroy resource
  _err = DestroyCltCert(siteCltCertRespBody, cltCertRespBody, esa20240910Client)
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
