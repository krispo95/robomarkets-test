package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"geobase/config"
	"geobase/internal/repository"
	"geobase/internal/usecase"
	"github.com/valyala/fasthttp"
	"log"
)

func StartServer(cfg *config.Config, uc usecase.Usecase) {
	// Parse command-line flags.
	flag.Parse()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ip/location":
			fmt.Println("1")
			IpLocationHandler(ctx, uc)
		case "/city/locations":
			fmt.Println("2")
			CityLocationHandler(ctx, uc)
		default:

		}
	}

	// Start HTTP server.
	if len(cfg.Port) > 0 {
		log.Printf("Starting HTTP server on %q", cfg.Port)
		go func() {
			if err := fasthttp.ListenAndServe(cfg.Port, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServe: %s", err)
			}
		}()
	}

	select {}
}

func CityLocationHandler(ctx *fasthttp.RequestCtx, uc usecase.Usecase) {
	cityName := ctx.QueryArgs().Peek("city")
	location := uc.FindLocationByName(string(cityName))
	bts, err := json.Marshal(location)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(bts)
}

func IpLocationHandler(ctx *fasthttp.RequestCtx, uc usecase.Usecase) {
	ipBts := ctx.QueryArgs().Peek("ip")
	location := uc.FindLocationByIP(repository.ParseUint32(ipBts))
	bts, err := json.Marshal(location)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(bts)
}
