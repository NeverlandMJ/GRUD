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
	g := int(req.GetPage())
	s.GetPosts(g)
	return &collectpb.Empty{}, nil
}

func (s Server) GetPosts(g int) error {
	// var mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(g)

	for i := 1; i <= g; i++ {
		k := strconv.Itoa(i)
		go func() {
			defer wg.Done()
			url := "https://gorest.co.in/public/v1/posts?page=" + url.QueryEscape(k)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			data := new(entity.Post)
			err = json.Unmarshal(body, data)
			if err != nil {
				panic(err)
			}

			err = s.svc.SavePost(context.Background(), data.Data)
			if err != nil {
				panic(err)
			}

		}()
	}

	wg.Wait()

	return nil
}
