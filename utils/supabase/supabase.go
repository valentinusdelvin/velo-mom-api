package supabase

import (
	"mime/multipart"
	"os"

	supabase_storage_uploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type InterSupabase interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(link string) error
}

type storageSupabase struct {
	client *supabase_storage_uploader.Client
}

func Init() InterSupabase {
	spbClient := supabase_storage_uploader.New(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_TOKEN"),
		os.Getenv("SUPABASE_BUCKET"),
	)

	return &storageSupabase{
		client: spbClient,
	}
}

func (s *storageSupabase) Upload(file *multipart.FileHeader) (string, error) {
	url, err := s.client.Upload(file)
	if err != nil {
		return url, err
	}

	return url, nil
}

func (s *storageSupabase) Delete(link string) error {
	err := s.client.Delete(link)
	if err != nil {
		return err
	}

	return nil
}
