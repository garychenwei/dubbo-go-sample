package app

import "github.com/iotcenter/golange/proto/erp"
import (
	context "context"
)
// import (
// 	perrors "github.com/pkg/errors"
// )

type MeterialProvider struct {
	erp.UnimplementedMaterialServiceServer
}

func (s *MeterialProvider) GetMeterial(context context.Context,req *erp.GetMeterialRequest) (*erp.GetMeterialResponse, error){
	responseData := &erp.GetMeterialResponse{
		Page: 1,
		Size: 10,
		Records: []*erp.Material{},
	}

	// err := perrors.New("")
	return responseData, nil
}