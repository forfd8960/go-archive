package main

import (
	"flag"
	"net"
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/forfd8960/go-archive/pb"
	"github.com/forfd8960/go-archive/conf"
	"github.com/forfd8960/go-archive/db"
	"github.com/forfd8960/go-archive/archive"
	"github.com/forfd8960/go-archive/model"
)

func main() {
	configPath := flag.String("config", "conf/config.toml", "config file's path")
	flag.Parse()

	conf.LoadConfig(*configPath)
	db.InitMysql(conf.Conf.MysqlConf)

	listen, err := net.Listen("tcp", conf.Conf.ListenAddr)
	if err != nil {
		log.Fatalf("serve error: %v\n", err)
	}

	ctx := context.Background()
	server := archive.NewServer(model.NewArchiveItemer(ctx, db.MySql))

	s := grpc.NewServer()
	pb.RegisterGoArchiveServer(s, server)
	if err = s.Serve(listen); err != nil {
		log.Fatalf("serve error: %v\n", err)
	}
}
