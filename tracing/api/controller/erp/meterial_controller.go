package erp

import (
	"context"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gofiber/fiber/v2"
	provider "github.com/iotcenter/golange/api/gobal"
	"github.com/iotcenter/golange/proto/erp"
)

// MeterialController api
type MeterialController struct{}
// GetMeterial by dubbo
func(s *MeterialController) GetMeterial(c *fiber.Ctx) error {
	req := &erp.GetMeterialRequest{
		Name: "",
	}
	result, err := provider.MaterialServiceClientImpl.GetMeterial(context.Background(), req)
	if err != nil {
		logger.Error("inner error ", err)
		return err
	}
	
	return c.JSON(&result)
}
// MeterialControllerImpl impl
var MeterialControllerImpl = MeterialController{}