package usecase

import (
	"github.com/77InnovationLabs/nebula-back/pessoa/internal/domain/entity"
	"github.com/77InnovationLabs/nebula-back/pessoa/internal/domain/repository"
	"github.com/77InnovationLabs/nebula-back/pessoa/pkg/event_dispatcher"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
	OrderCreated    event_dispatcher.EventInterface
	EventDispatcher event_dispatcher.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository repository.OrderRepositoryInterface,
	OrderCreated event_dispatcher.EventInterface,
	EventDispatcher event_dispatcher.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
