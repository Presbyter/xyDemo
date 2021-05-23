package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"xyDemo/cmd/c_server/api"
)

func main() {
	ch := make(chan *http.Request, 1<<10)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 拿到用户传入的数据
		ch <- r
		w.WriteHeader(http.StatusOK)
	})
	go http.ListenAndServe(":3000", nil)

	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := api.NewCServiceClient(conn)
	//resp, err := client.Forward(context.Background(), &api.ForwardReq{UserName: "xingchao"})
	//if err != nil {
	//	panic(err)
	//}
	stream, err := client.Forward(context.Background())
	if err != nil {
		panic(err)
	}

	for v := range ch {
		t, err := json.Marshal(v.Header)
		if err != nil {
			panic(err)
		}
		if err := stream.Send(&api.ForwardReq{UserName: string(t)}); err != nil {
			panic(err)
		}

		r, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("stream EOF")
			return
		} else if err != nil {
			panic(err)
		}
		fmt.Println(r.GetResponse())
	}

}
