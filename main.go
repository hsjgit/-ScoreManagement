package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ScoreManagement/bootstrap"
)

var (
	listen   = flag.String("listen", ":8090", "listen")
	userName = flag.String("userName", "huangshijie", "db userName")
	password = flag.String("password", "123456", "db password")
	ip       = flag.String("ip", "hsj.flyaha.top", "db host")
	port     = flag.String("port", "3306", "db port")
	dbName   = flag.String("dbName", "yyh", "db Name")
)

func init() {
	flag.Parse()
	path := strings.Join([]string{*userName, ":", *password, "@tcp(", *ip, ":", *port, ")/", *dbName, "?charset=utf8mb4"}, "")
	if bootstrap.Bootstrap(path, *listen) != nil {
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
