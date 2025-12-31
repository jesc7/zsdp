package store

type IStore interface {
	SendOffer(string, ...any) (int, string, error)
	SendAnswer(int, string, string, ...any) (any, error)
}

var Store IStore
