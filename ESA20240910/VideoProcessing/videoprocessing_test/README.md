#### Project Overview
This project is a complete engineering example for "creating, updating, and deleting video processing configurations" under the VideoProcessing resource of the Alibaba Cloud product ESA (version: 2024-09-10). This example demonstrates how to:
- Create a new video processing configuration.
- Update an existing video processing configuration.
- Delete a previously created video processing configuration.

#### Notes
- **Operation Costs**:
  - Running the sample code may incur costs for online resource operations on the current account. Please proceed with caution!

- **Dependencies**:
  - The `VideoProcessing` resource depends on the `Site` resource, so you must successfully create a site before creating a video processing configuration.
  - The `Site` resource depends on the `RatePlanInstance` resource, so you must successfully purchase a plan before creating a site.

- **Asynchronous Operations**:
  - Purchasing a new plan and creating a site are both asynchronous operations. You need to wait for their status to update to the specified state before proceeding to the next step.

#### Workflow
1. **Initialize Client**
   - To run this example, you must first configure your credentials as described in [Credential Configuration](https://help.aliyun.com/zh/sdk/developer-reference/v2-manage-net-access-credentials). Create a client instance.
   
2. **Purchase Plan**
   - Call the [PurchaseRatePlan](https://api.aliyun.com/api/ESA/2024-09-10/PurchaseRatePlan) API to purchase a new resource plan and wait for its status to change to "running".

3. **Create Site**
   - After successfully purchasing the plan, call the [CreateSite](https://api.aliyun.com/api/ESA/2024-09-10/CreateSite) API to create a new site and wait for its status to change to "pending".

4. **Create Video Processing Configuration**
   - After the site is successfully created, call the [CreateVideoProcessing](https://api.aliyun.com/api/ESA/2024-09-10/CreateVideoProcessing) API to create a new video processing configuration.

5. **Update Video Processing Configuration**
   - Call the [UpdateVideoProcessing](https://api.aliyun.com/api/ESA/2024-09-10/UpdateVideoProcessing) API to update the existing video processing configuration.

6. **Delete Video Processing Configuration**
   - Finally, call the [DeleteVideoProcessing](https://api.aliyun.com/api/ESA/2024-09-10/DeleteVideoProcessing) API to delete the previously created video processing configuration, completing the entire lifecycle management.

#### How to Run
- *Requires Go 1.10.x or later.*
- *Install SDK core OpenAPI*
```sh
go get github.com/alibabacloud-go/darabonba-openapi/v2/client
go get github.com/alibabacloud-go/main
go get github.com/alibabacloud-go/tea/dara
```
- *Command*
```sh
GOPROXY=https://goproxy.cn,direct go run ./main
```

#### More Samples
For additional examples, visit: https://github.com/aliyun/alibabacloud-go-sdk-samples