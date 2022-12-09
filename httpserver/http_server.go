package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

func RunServer(lifecycle fx.Lifecycle, r *gin.Engine) {
	mux := NewMuxServer(r)

	port := viper.GetInt("server.port")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("HTTP server starting at port:", port)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal("HTTP server start error: ", err)
				}
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			log.Println("HTTP server Shutting down...")
			timeout := viper.GetInt("server.shutdown-timeout-sec")
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
			defer cancel()
			return srv.Shutdown(ctx)
		},
	})
}

func NewMuxServer(r *gin.Engine) *http.ServeMux {
	mux := http.NewServeMux()
	if r != nil {
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			r.ServeHTTP(w, req)
		})
	}
	return mux
}
