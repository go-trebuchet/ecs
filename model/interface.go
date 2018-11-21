package model


type Entity interface{
	Id() uint64
	Component(string) Component
	Components() []Component
	AddComponent(Component)
}

type Component interface{
	Type() string
}


type Matcher interface{
	Match(Entity) bool
}

type System interface{
	Matcher() Matcher
	Init(Pool)
	Update(Pool)
	AddEntity(Entity)
	WillRemoveEntity(Entity)
	RemoveEntity(Entity)
}


type Pool interface {
	CreateEntity() Entity
	Entities() []Entity
	GetEntityByID(uint64) (Entity, error)
	Count() int
	HasEntity(uint64) bool
	DestroyEntity(uint64)
	DestroyAllEntities()
	AddSystem(System, priority int)
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
