package abstractions

type ICommandHanler[T any] interface {
	Handle() (T, error)
}
