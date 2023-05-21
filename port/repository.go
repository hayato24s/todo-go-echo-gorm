package port

type IRepository interface {
	// Begin begins a transaction.
	Begin() (IRepository, error)

	// Rollback rollbacks the changes in a transaction.
	Rollback() error

	// Commit commits the changes in a transaction.
	Commit() error

	IUserRepository
	ITaskRepository
}
