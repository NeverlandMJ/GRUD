package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/NeverlandMJ/GRUD/collecter-service/collectpb"
	"github.com/NeverlandMJ/GRUD/collecter-service/entity"
	"github.com/NeverlandMJ/GRUD/collecter-service/service"
)

type Server struct {
	collectpb.UnimplementedCollecterServiceServer
	svc service.Service
}

func New(svc service.Service) Server {
	return Server{
		svc: svc,
	}
}

func (s Server) SavePosts(ctx context.Context, req *collectpb.DownloadReq) (*collectpb.Empty, error) {
	pages := int(req.GetPage())

	var wg sync.WaitGroup
	ch := make(chan error)

	for i := 1; i <= pages; i++ {
		wg.Add(1)
		go func(page int) {
			err := s.GetAndSavePosts(page)
			if err != nil {
				ch <- err
				return
			}
			ch <- nil
			defer wg.Done()
		}(i)
	}

	err := <-ch
	if err != nil {
		return nil, err
	}
	return &collectpb.Empty{}, nil
}

func (s Server) GetAndSavePosts(page int) error {
	p := strconv.Itoa(page)
	url := "https://gorest.co.in/public/v1/posts?page=" + url.QueryEscape(p)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	data := new(entity.Post)
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	err = s.svc.SavePost(context.Background(), data.Data)
	if err != nil {
		return err
	}
	return nil
}
