package order_service

import (
	"7kzu-order-service/internal/data/model"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

func (s *orderService) UpdateOrder(ctx context.Context, order *model.Order) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderService.UpdateOrder")
	defer span.Finish()

	if order.IsFoundCourier() {
		fmt.Println("found courier status")
		if err := s.repo.UpdateOrder(ctx, order.ID, func(old *model.Order) *model.Order {
			return &model.Order{
				ID:          old.ID,
				CourierID:   order.CourierID,
				MapUrl:      old.MapUrl,
				PutUrl:      old.PutUrl,
				Status:      order.Status,
				CompanyId:   old.CompanyId,
				Price:       old.Price,
				TransportId: old.TransportId,
				IsDraft:     old.IsDraft,
				Code:        old.Code,
			}
		}); err != nil {
			return err
		}
	}

	if order.IsInShop() {
		fmt.Println("in shop status")
		mapUrl := `https://7kzu.kz/map/` + order.ID.String()
		if err := s.repo.UpdateOrder(ctx, order.ID, func(old *model.Order) *model.Order {
			return &model.Order{
				ID:          old.ID,
				CourierID:   old.CourierID,
				MapUrl:      mapUrl,
				PutUrl:      old.PutUrl,
				Status:      order.Status,
				CompanyId:   old.CompanyId,
				Price:       old.Price,
				TransportId: old.TransportId,
				IsDraft:     old.IsDraft,
				Code:        old.Code,
			}
		}); err != nil {
			return err
		}
	}

	if order.IsInHome() {
		code := `1233`
		if err := s.repo.UpdateOrder(ctx, order.ID, func(old *model.Order) *model.Order {
			return &model.Order{
				ID:          old.ID,
				CourierID:   old.CourierID,
				MapUrl:      old.MapUrl,
				PutUrl:      old.PutUrl,
				Status:      order.Status,
				CompanyId:   old.CompanyId,
				Price:       old.Price,
				TransportId: old.TransportId,
				IsDraft:     old.IsDraft,
				Code:        code,
			}
		}); err != nil {
			return err
		}
	}

	if order.IsSuccess() {
		if err := s.repo.UpdateOrder(ctx, order.ID, func(old *model.Order) *model.Order {
			return &model.Order{
				ID:          old.ID,
				CourierID:   old.CourierID,
				MapUrl:      old.MapUrl,
				PutUrl:      old.PutUrl,
				Status:      order.Status,
				CompanyId:   old.CompanyId,
				Price:       old.Price,
				TransportId: old.TransportId,
				IsDraft:     old.IsDraft,
				Code:        " ",
			}
		}); err != nil {
			return err
		}
	}

	return nil
}
