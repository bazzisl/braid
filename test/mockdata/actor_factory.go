package mockdata

import (
	"github.com/pojol/braid/actors"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/def"
)

// MockActorFactory is a factory for creating actors
type MockActorFactory struct {
	constructors map[string]*core.ActorConstructor
}

// NewActorFactory create new actor factory
func BuildActorFactory() *MockActorFactory {
	factory := &MockActorFactory{
		constructors: make(map[string]*core.ActorConstructor),
	}

	factory.bind("MockUserActor", false, 80, 10000, NewUserActor)
	factory.bind("MockClacActor", false, 20, 5, NewClacActor)

	factory.bind(def.ActorDynamicPicker, false, 160, 3, actors.NewDynamicPickerActor)
	factory.bind(def.ActorDynamicRegister, true, 80, 0, actors.NewDynamicRegisterActor)

	return factory
}

// Bind associates an actor type with its constructor function
func (factory *MockActorFactory) bind(actorType string, unique bool, weight, limit int, f core.CreateFunc) {
	factory.constructors[actorType] = &core.ActorConstructor{
		NodeUnique:          unique,
		Weight:              weight,
		GlobalQuantityLimit: limit,
		Constructor:         f,
	}
}

func (factory *MockActorFactory) Get(actorType string) *core.ActorConstructor {
	if _, ok := factory.constructors[actorType]; ok {
		return factory.constructors[actorType]
	}

	return nil
}
