package storage

type StorageI interface {
	User() UserRepoI
}

type UserRepoI interface {
	// Create(ctx context.Context, entity *pb.CreateUserRequest) (pKey *pb.UserPrimaryKey, err error)
}
