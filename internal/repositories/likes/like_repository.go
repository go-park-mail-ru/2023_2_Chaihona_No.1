package likes

type LikeRepository interface {
	CreateNewLike(userId int, postId int) error
	DeleteLike(userId int, postId int) error
}
