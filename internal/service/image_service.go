package service

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"mime/multipart"
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

func (i *imageService) UploadImage(image *multipart.FileHeader, domainImage *entity.Image) (id int64, err error) {
	err = i.str.Save(image, domainImage)
	if err != nil {
		return id, err
	}
	id, err = i.ir.SaveImage(*domainImage)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (i *imageService) DownloadImages(id int64) ([]entity.Image, error) {
	images, err := i.ir.GetImages(id)
	if err != nil {
		return nil, err
	}
	return images, nil
}
