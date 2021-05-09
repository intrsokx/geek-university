package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//基于 errgroup 实现一个 http server 的启动和关闭，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	//这里返回的是一个errGroup和一个cancelCtx对象，但是我们没有发现手动调用cancel的方法
	//这是因为在errGroup在wait结束时已经帮我们调用了cancel方法
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.Server{
		Addr:           ":8080",
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}

	//svr start
	g.Go(func() error {
		fmt.Println("http server start...")
		return svr.ListenAndServe()
	})

	//svr shutdown
	g.Go(func() error {
		<-ctx.Done()
		fmt.Println("http server shutdown.")
		c, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()
		return svr.Shutdown(c)
	})

	// signal
	g.Go(func() error {
		sigs := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGUSR2}
		ch := make(chan os.Signal, len(sigs))
		signal.Notify(ch, sigs...)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case s := <-ch:
				fmt.Println("receive sig:", s.String())
				return errors.New("receive sig " + s.String())
			}
		}
	})

	//mock error
	g.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("mock error")
	})

	if err := g.Wait(); err != nil {
		fmt.Println("occurred err:", err)
	}
}
