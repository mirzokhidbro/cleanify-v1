package repo

type CommentStorageI interface {
	Delete(id int) error
	UploadVoice(filePath string) (string, error) // voice faylni yuklab, URL qaytaradi
}
