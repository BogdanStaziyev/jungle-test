package service

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"mime/multipart"
	"net/url"
)

type imageService struct {
	ir  ImageRepo
	str FileStorage
}

func NewImageService(imageRepo ImageRepo, imageStorage FileStorage) *imageService {
	return &imageService{
		ir:  imageRepo,
		str: imageStorage,
	}
}

func (i *imageService) UploadImage(image *multipart.FileHeader, entityImage entity.Image) (string, error) {
	err := i.str.Save(image, &entityImage)
	if err != nil {
		return "", err
	}
	baseURL, err := url.Parse(entityImage.URL)
	if err != nil {
		return "", err
	}
	res := baseURL.ResolveReference(&url.URL{Path: entityImage.Path})
	entityImage.URL = res.String()
	err = i.ir.SaveImage(entityImage)
	if err != nil {
		return "", err
	}
	return entityImage.URL, nil
}

func (i *imageService) DownloadImages(id int64) ([]entity.Image, error) {
	images, err := i.ir.GetImages(id)
	if err != nil {
		return nil, err
	}
	return images, nil
}
