package store

var client Factory

// Factory defines the TikSound-backend platform storage interface.
type Factory interface {
	User() UserStore
	Video() VideoStore
	Comment() CommentStore
	Close() error
}
