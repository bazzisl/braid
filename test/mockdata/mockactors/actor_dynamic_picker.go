package mockactors

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/router"
	"golang.org/x/time/rate"
)

type dynamicPickerActor struct {
	*actor.Runtime
	limiter *rate.Limiter
}

func NewDynamicPickerActor(p core.IActorBuilder) core.IActor {
	return &dynamicPickerActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: "MockDynamicPicker", Sys: p.GetSystem()},
		limiter: rate.NewLimiter(rate.Every(time.Second/200), 1), // 允许每秒10次调用
	}
}

func (a *dynamicPickerActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent("MockDynamicPick", func(ctx core.ActorContext) core.IChain {
		return &actor.DefaultChain{

			Handler: func(mw *router.MsgWrapper) error {

				// 使用限流器
				if err := a.limiter.Wait(mw.Ctx); err != nil {
					return err
				}

				actor_ty := mw.Req.Header.Custom["actor_ty"]

				// Select a node with low weight and relatively fewer registered actors of this type
				nodeaddr, err := ctx.AddressBook().GetLowWeightNodeForActor(mw.Ctx, actor_ty)
				if err != nil {
					return err
				}

				// rename
				mw.Req.Header.Custom["actor_id"] = nodeaddr.Node + "_" + actor_ty + "_" + uuid.NewString()

				// dispatcher to picker node
				return ctx.Call(router.Target{ID: nodeaddr.Node + "_" + "MockDynamicRegister", Ty: "MockDynamicRegister", Ev: "MockDynamicRegister"}, mw)
			},
		}
	})
}
