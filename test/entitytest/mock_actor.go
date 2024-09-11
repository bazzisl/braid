package entitytest

import (
	"context"
	"fmt"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
)

type mockUserActor struct {
	*actor.Runtime
	entity *EntityWapper
}

func NewEntity(p *core.CreateActorParm) core.IActor {
	return &mockUserActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: "mockUserActor"},
		entity:  NewEntityWapper(p.ID),
	}
}

func (a *mockUserActor) Init() {
	a.Runtime.Init()
	err := a.entity.Load(context.TODO())
	if err != nil {
		panic(fmt.Errorf("load user actor err %v", err.Error()))
	}

	// Implement events
	a.RegisterEvent("entity_test", &actor.DefaultChain{
		Handler: func(ctx context.Context, m *router.MsgWrapper) error {

			if a.entity.Bag.EnoughItem(1001, 10) {
				a.entity.Bag.ConsumeItem(1001, 5, "test", "")

				// mark success
				fmt.Println("entity_test consume item success")
				m.Res.Header.Custom["code"] = "200"
			}

			return nil
		},
	})

	// one minute try sync to cache
	a.RegisterTimer(0, 1000*60, func() error {
		a.entity.Sync(context.TODO())

		return nil
	}, nil)
}