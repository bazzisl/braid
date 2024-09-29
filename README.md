# braid
> 

[![image.png](https://i.postimg.cc/3xNDLTwR/image.png)](https://postimg.cc/ts0TT8WQ)

[![Go Report Card](https://goreportcard.com/badge/github.com/pojol/braid?style=flat-square)](https://goreportcard.com/report/github.com/pojol/braid)
[![Demo](https://img.shields.io/badge/demo-braid--demo-brightgreen?style=flat-square)](https://github.com/pojol/braid-demo)
[![Discord](https://img.shields.io/discord/1210543471593791488?color=7289da&label=Discord&logo=discord&style=flat-square)](https://discord.gg/yXJgTrkWxT)


### Sample

1. register actor
```go
// factory  e.g. test/mockdata/actor_factory
factory.bind("MockClacActor", 
    false,          // whether the node is unique
    20,             // weight of the actor
    50000,          // maximum number of actors to be built in the cluster
    NewClacActor,   // constructor function for the actor
)
```

2. builder actor
```go

// register ActorDynamicRegister type actor to the local node
sys.Loader().Builder(def.ActorDynamicRegister).WithID("nodeid-register").RegisterLocally()

// Or register a MockClacActor type actor to the cluster dynamically (via load balancing)
sys.Loader().Builder("MockClacActor").WithID("001").RegisterDynamically()
```

3. Implement logic for the actor
```go

// Bind message handler
clacActor.RegisterEvent("ev_clac", func(actorCtx context.Context) *actor.DefaultChain {
    
    // use middleware
    unpackcfg := &middleware.MsgUnpackCfg[proto.xxx]{}
    sys := core.GetSystem(actorCtx)

    return &actor.DefaultChain{
        Before: []Base.MiddlewareHandler{
            middleware.MsgUnpack(unpackcfg),
        },
        Handler: func(ctx context.Context, msg *router.MsgWrapper) error {

            realmsg, ok := unpackcfg.Msg.(*proto.xxx)
            // todo ...

            // Pass the message downstream
            sys.Call(...)

            return nil
        }
    }
})

// Register a periodic processing function
clacActor.RegisterTimer(0, 1000, func(actorCtx context.Context) error {

    state := core.GetState(actorCtx).(*xxxState)

    if state.State == Init {
        // todo & state transitions
        state.State = Running
    } else if state.State == Running {

    }

    return nil
})

// Subscribe to messages (chat messages sent by others when offline)
//  topic: Offline chat messages
//  channel: The actor itself
//  succ: Callback after successful subscription
//  ttl: Time-to-live for messages in the cache
clacActor.SubscriptionEvent(events.EvChatMessageStore, a.Id, func() {

    // After successful subscription, bind a handler function for the message
    a.RegisterEvent(events.EvChatMessageStore, events.MakeChatStoreMessage)
    
}, pubsub.WithTTL(time.Hour*24*30))
```


---

### benchmark