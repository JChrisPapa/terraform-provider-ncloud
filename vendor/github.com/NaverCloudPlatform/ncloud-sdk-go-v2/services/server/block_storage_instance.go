/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-08-20T06:35:09Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type BlockStorageInstance struct {

	// 블록스토리지인스턴스번호
BlockStorageInstanceNo *string `json:"blockStorageInstanceNo,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 서버명
ServerName *string `json:"serverName,omitempty"`

	// 블록스토리지구분
BlockStorageType *CommonCode `json:"blockStorageType,omitempty"`

	// 블록스토리지명
BlockStorageName *string `json:"blockStorageName,omitempty"`

	// 블록스토리지사이즈
BlockStorageSize *int64 `json:"blockStorageSize,omitempty"`

	// 디바이스명
DeviceName *string `json:"deviceName,omitempty"`

	// 회원서버이미지번호
MemberServerImageNo *string `json:"memberServerImageNo,omitempty"`

	// 블록스토리지상품코드
BlockStorageProductCode *string `json:"blockStorageProductCode,omitempty"`

	// 블록스토리지인스턴스상태
BlockStorageInstanceStatus *CommonCode `json:"blockStorageInstanceStatus,omitempty"`

	// 블록스토리지인스턴스OP
BlockStorageInstanceOperation *CommonCode `json:"blockStorageInstanceOperation,omitempty"`

	// 블록스토리지인스턴스상태명
BlockStorageInstanceStatusName *string `json:"blockStorageInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 블록스토리지인스턴스설명
BlockStorageInstanceDescription *string `json:"blockStorageInstanceDescription,omitempty"`

	// 디스크유형
DiskType *CommonCode `json:"diskType,omitempty"`

	// 디스크상세유형
DiskDetailType *CommonCode `json:"diskDetailType,omitempty"`

	// 최대 IOPS
MaxIopsThroughput *int32 `json:"maxIopsThroughput,omitempty"`

Region *Region `json:"region,omitempty"`

Zone *Zone `json:"zone,omitempty"`
}