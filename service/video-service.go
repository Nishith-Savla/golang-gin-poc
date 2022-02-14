package service

import "github.com/Nishith-Savla/golang-gin-poc/entity"

type VideoService interface {
	Save(entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (service *videoService) Save(video entity.Video) {
	service.videos = append(service.videos, video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
