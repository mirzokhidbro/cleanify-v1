package postgres

import (
	"bw-erp/storage/repo"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) repo.CommentStorageI {
	return &commentRepo{db: db}
}

func (repo commentRepo) Delete(id int) error {
	query := "DELETE FROM comments WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("comment with id %d not found", id)
	}

	return nil
}

func (repo commentRepo) UploadVoice(filePath string) (string, error) {
	// Voice fayllarni saqlash uchun papka
	uploadDir := "uploads/voices"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	// Yangi fayl nomi yaratish
	fileName := uuid.New().String() + filepath.Ext(filePath)
	newPath := filepath.Join(uploadDir, fileName)

	// Faylni ko'chirish
	if err := os.Rename(filePath, newPath); err != nil {
		return "", err
	}

	// URL qaytarish (sizning sistemangizga qarab o'zgartiring)
	return "/uploads/voices/" + fileName, nil
}
