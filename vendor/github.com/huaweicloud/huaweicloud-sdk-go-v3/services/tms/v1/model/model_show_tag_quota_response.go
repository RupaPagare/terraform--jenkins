package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTagQuotaResponse struct {

	// ้้ขๅ่กจ
	Quotas         *[]TagQuota `json:"quotas,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTagQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowTagQuotaResponse", string(data)}, " ")
}
