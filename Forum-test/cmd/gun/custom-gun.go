// create a package
package main

import (
	"context"
	"log"
	"time"

	"github.com/spf13/afero"
	"github.com/storm5758/Forum-test/pkg/api"
	"github.com/yandex/pandora/cli"
	"github.com/yandex/pandora/core"
	"github.com/yandex/pandora/core/aggregator/netsample"
	coreimport "github.com/yandex/pandora/core/import"
	"github.com/yandex/pandora/core/register"
	"google.golang.org/grpc"
)

type Ammo struct {
	Tag  string
	Name string
}

type Sample struct {
	URL              string
	ShootTimeSeconds float64
}

type GunConfig struct {
	Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
	// Configured on construction.
	client grpc.ClientConn
	conf   GunConfig
	// Configured on Bind, before shooting
	aggr core.Aggregator // May be your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	// create gRPC stub at gun initialization
	conn, err := grpc.Dial(
		g.conf.Target,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
		grpc.WithUserAgent("load test, pandora custom shooter"))
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	g.client = *conn
	g.aggr = aggr
	g.GunDeps = deps
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo)
	g.shoot(customAmmo)
}

func (g *Gun) case1_method(client api.UserClient, ammo *Ammo) int {
	code := 0
	// prepare list of ids from ammo
	var name string
	name = ammo.Name
	out, err := client.UserGetOne(
		context.TODO(), &api.UserGetOneRequest{
			Nickname: name})

	if err != nil {
		log.Printf("FATAL: %s", err)
		code = 500
	}

	if out != nil {
		code = 200
	}
	return code
}

func (g *Gun) shoot(ammo *Ammo) {
	code := 0
	sample := netsample.Acquire(ammo.Tag)

	conn := g.client
	client := api.NewUserClient(&conn)

	switch ammo.Tag {
	case "/MyCase1":
		code = g.case1_method(client, ammo)
	default:
		code = 404
	}

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(sample)
	}()
}

func main() {
	//debug.SetGCPercent(-1)
	// Standard imports.
	fs := afero.NewOsFs()
	coreimport.Import(fs)

	// Custom imports. Integrate your custom types into configuration system.
	coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })

	register.Gun("My_custom_gun_name", NewGun, func() GunConfig {
		return GunConfig{
			Target: "localhost:6000",
		}
	})

	cli.Run()
}
