package controllers

import (
	"github.com/8treenet/freedom"
	"github.com/8treenet/freedom/example/http2/business/services"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		initiator.BindController("/shop", &ShopController{})
	})
}

type ShopController struct {
	Runtime  freedom.Runtime
	Shopping services.ShoppingInterface
}

// Get handles the GET: /shop/:id route.
func (s *ShopController) GetBy(id int) string {
	s.Runtime.Logger().Info("我是控制器", "ShopController.GetByID")
	return s.Shopping.Shopping(id)
}
