package workerthread

type SystemParm struct {
	ServiceName string
	NodeID      string
	Ip          string
	Port        int

	Constructors []ActorConstructor
}

type ActorConstructor struct {
	Type        string
	Constructor CreateFunc
}

type SystemOption func(*SystemParm)

func SystemService(serviceName, nodeID string) SystemOption {
	return func(sp *SystemParm) {
		sp.NodeID = nodeID
		sp.ServiceName = serviceName
	}
}

func SystemActorConstructor(lst []ActorConstructor) SystemOption {
	return func(sp *SystemParm) {
		sp.Constructors = append(sp.Constructors, lst...)
	}
}

func SystemWithAcceptor(port int) SystemOption {
	return func(sp *SystemParm) {
		sp.Port = port
	}
}

type CreateActorParm struct {
	ID     string
	Sys    ISystem
	InsPtr interface{}
}

type CreateActorOption func(*CreateActorParm)

func CreateActorWithID(id string) CreateActorOption {
	return func(p *CreateActorParm) {
		p.ID = id
	}
}

func CreateActorWithIns(ins interface{}) CreateActorOption {
	return func(p *CreateActorParm) {
		p.InsPtr = ins
	}
}
