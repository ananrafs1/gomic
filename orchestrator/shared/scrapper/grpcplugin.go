package plugin
import (
	"github.com/ananrafs1/gomic/orchestrator/proto"
	"github.com/ananrafs1/gomic/model"
	"log"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"context"
)

type GRPCPlugin struct {
	plugin.Plugin
	Impl Scrapper
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterScrapperServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewScrapperClient(c)}, nil
}

type GRPCClient struct{ client proto.ScrapperClient }

func (m *GRPCClient) Scrap(Title string, Page, Quantity int) model.Comic {
	resp, err := m.client.Scrap(context.Background(), &proto.ScrapRequest{
		Title:   Title,
		Page: int64(Page),
		Quantity: int64(Quantity),
	})
	if err != nil {
		log.Fatal(err)
		return model.Comic{}
	}
	ret := model.Comic{
		ComicFlat : model.ComicFlat{
			Id : int(resp.Comic.ComicFlat.Id),
			Name : resp.Comic.ComicFlat.Name,
			Host : resp.Comic.ComicFlat.Host,
		},
		Chapters : make([]model.Chapter,0),
	}
	for _, v := range resp.Comic.Chapters {
		chpter := model.Chapter{
			ChapterFlat : model.ChapterFlat{
				Id : v.ChapterFlat.Id,
			},
			Images : make([]model.ImageProvider, 0),
		}
		for _, img := range v.Images {
			image := model.ImageProvider{
				Episode :int(img.Episode),
				Provider : img.Provider,
				Link : img.Link,
			}
			chpter.Images = append(chpter.Images, image)
		}
		ret.Chapters = append(ret.Chapters, chpter)
	}
	return ret
}

func (m *GRPCClient) ScrapPerChapter(Title, Id string) model.Chapter {
	resp, err := m.client.ScrapPerChapter(context.Background(), &proto.ScrapPerChapterRequest{
		Title: Title,
		Id : Id,
	})
	if err != nil {
		log.Fatal(err)
		return model.Chapter{}
	}
	ret := model.Chapter{
		ChapterFlat: model.ChapterFlat{
			Id : resp.Chapter.ChapterFlat.Id,
		},
		Images : make([]model.ImageProvider, 0),
	}
	for _, img := range resp.Chapter.Images {
		image := model.ImageProvider{
			Episode : int(img.Episode),
			Provider : img.Provider,
			Link : img.Link,
		}
		ret.Images = append(ret.Images, image)
	}
	return ret
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Scrapper
}

func (m *GRPCServer) Scrap(
	ctx context.Context,
	req *proto.ScrapRequest) (*proto.ScrapResponse, error) {
		val := m.Impl.Scrap(req.Title, int(req.Page), int(req.Quantity))

		ret := proto.ScrapResponse{
			Comic : &proto.Comic{
				ComicFlat : &proto.ComicFlat{
					Id : int64(val.ComicFlat.Id),
					Name: val.ComicFlat.Name,
					Host: val.ComicFlat.Host,
				},
			},
		}
		chapterHolder := make([](*proto.Chapter), 0)
		for _, v := range val.Chapters {
			chpter := proto.Chapter{
				ChapterFlat: &proto.ChapterFlat{
					Id : v.ChapterFlat.Id,
				},
			}
			imageHolder := make([](*proto.ImageProvider), 0)
			for _, img := range v.Images {
				img := proto.ImageProvider{
					Episode: int64(img.Episode),
					Provider: img.Provider,
					Link : img.Link,
				}
				imageHolder = append(imageHolder, &img)
			}
			chpter.Images = imageHolder
			chapterHolder = append(chapterHolder, &chpter)
		}
		ret.Comic.Chapters = chapterHolder

	return &ret, nil
}

func (m *GRPCServer) ScrapPerChapter(
	ctx context.Context,
	req *proto.ScrapPerChapterRequest) (*proto.ScrapPerChapterResponse, error) {
	v := m.Impl.ScrapPerChapter(req.Title, req.Id)
	ret := proto.ScrapPerChapterResponse{
		Chapter : &proto.Chapter{
			ChapterFlat: &proto.ChapterFlat{
				Id : v.ChapterFlat.Id,
			},
		},
	}
	imageHolder := make([](*proto.ImageProvider), 0)
	for _, img := range v.Images {
		img := proto.ImageProvider{
				Episode : int64(img.Episode),
				Provider: img.Provider,
				Link : img.Link,
				}
			imageHolder = append(imageHolder, &img)
	}
	ret.Chapter.Images = imageHolder

	return &ret, nil

}