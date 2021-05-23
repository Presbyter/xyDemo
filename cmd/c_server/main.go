package main

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"xyDemo/cmd/c_server/api"
)

//go:generate protoc --go_out=api --go-grpc_out=api api/c.proto
func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterCServiceServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}

type Server struct {
	api.UnimplementedCServiceServer
}

func (s Server) Forward(stream api.CService_ForwardServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		fmt.Println("Req Message: ", req.GetUserName())

		resp, err := http.Get("http://localhost:3002/")
		if err != nil {
			fmt.Println("request c server fail. error: %s", err)
			return errors.New("request c server fail")
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		err = stream.Send(&api.ForwardResp{
			Response: fmt.Sprintf("Hello %s. C return %s ~", req.GetUserName(), string(body)),
		})
		if err != nil {
			return err
		}
	}
}

func (s Server) mustEmbedUnimplementedCServiceServer() {
	panic("implement me")
}

//func (s Server) Forward(ctx context.Context, req *api.ForwardReq) (*api.ForwardResp, error) {
//	resp, err := http.Get("http://localhost:3002/")
//	if err != nil {
//		fmt.Println("request c server fail. error: %s", err)
//		return nil, errors.New("request c server fail")
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(req.GetUserName() + string(body))
//	return &api.ForwardResp{
//		Response: fmt.Sprintf("Hello %s. C return %s ~", req.GetUserName(), string(body)),
//	}, nil
//}
//
//func (s Server) mustEmbedUnimplementedCServiceServer() {
//	return
//}
