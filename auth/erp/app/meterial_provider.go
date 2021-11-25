package app

import "github.com/iotcenter/golange/proto/erp"
import (
	context "context"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

// import (
// 	perrors "github.com/pkg/errors"
// )

type MeterialProvider struct {
	erp.UnimplementedMaterialServiceServer
}

func (s *MeterialProvider) GetMeterial(context context.Context,req *erp.GetMeterialRequest) (*erp.GetMeterialResponse, error){
	signKey := context.Value(constant.AttachmentKey)
  logger.Info("receive AttachmentKey from provider ", signKey)
	responseData := &erp.GetMeterialResponse{
		Page: 1,
		Size: 10,
		Records: []*erp.Material{},
	}

	// err := perrors.New("")
	return responseData, nil
}