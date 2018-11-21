package model

type Entity interface{
	Id() uint64

	AddComponent(Component) error
	//ReplaceComponent(Component)
	//WillRemoveComponent(Component)
	RemoveComponent(Component) error
	//RemoveAllComponents()

	Component(uuid string) Component
	Components() []Component
	//ComponentTypes() []string
	HasComponent(uuid string) bool

	AddedComponent() AddedComponentPublisher
	RemovedComponent() RemovedComponentPublisher
}

type Component interface{
	Type() string
}


type Matcher interface{
	Matches(Entity) bool
}

//type Group interface{
//	Entities() []Entity
//	HandleEntity(Entity)
//	UpdateEntity(Entity)
//	WillRemoveEntity(Entity)
//	Matches(Entity) bool
//	ContainsEntity(Entity) bool
//}

type System interface{
	Matcher() Matcher
	AddPool(Pool)
	AddEntity(Entity)
	//WillRemoveEntity(Entity)
	RemoveEntity(Entity)
}

type SetupSystem interface{
	Setup()
	System
}

type UpdateSystem interface {
	Update()
	System
}

type CleanUpSystem interface {
	CleanUp()
	System
}

type TearDownSystem interface {
	TearDown()
	System
}

type Pool interface {
	CreateEntity() Entity
	Entities() []Entity
	Entity(uint64) (Entity, error)
	Count() int
	HasEntity(uint64) bool
	DestroyEntity(uint64)
	DestroyAllEntities()
	//Group(Matcher) Group
	AddSystem(System, priority int)
	//System(string) System
	//Systems() []System
	//RemoveSystem(string)
	//RemoveSystem(System)
	//RemoveAllSystems()
	Update()


	RemovedComponentSubscriber
	AddedComponentSubscriber
}






type AddedComponentSubscriber interface {
	AddedComponent(Entity, string)
}

type AddedComponentPublisher interface{
	Subscribe(subscriber AddedComponentSubscriber)
	Unsubscribe(subscriber AddedComponentSubscriber)
	Notify(Entity, string)
}

type RemovedComponentSubscriber interface {
	RemovedComponent(Entity, string)
}

type RemovedComponentPublisher interface{
	Subscribe(subscriber RemovedComponentSubscriber)
	Unsubscribe(subscriber RemovedComponentSubscriber)
	Notify(Entity, string)
}
