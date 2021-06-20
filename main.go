package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ScoreManagement/bootstrap"
)

var (
	listen   = flag.String("listen", ":80", "listen")
	userName = flag.String("userName", "huangshijie", "db userName")
	password = flag.String("password", "123456", "db password")
	ip       = flag.String("ip", "hsj.flyaha.top", "db host")
	port     = flag.String("port", "3306", "db port")
	dbName   = flag.String("dbName", "yyh", "db Name")
)

func init() {
	flag.Parse()
	path := strings.Join([]string{*userName, ":", *password, "@tcp(", *ip, ":", *port, ")/", *dbName, "?charset=utf8mb4"}, "")

	log.Println(path)
	if err := bootstrap.Bootstrap(path, *listen); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	select {
	case <-ctx.Done():
	}

}
