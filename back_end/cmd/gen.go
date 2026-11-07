package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "D:\\Projects\\test",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := os.Getenv("LOVEWALL_DSN")
	if dsn == "" {
		log.Fatal("LOVEWALL_DSN is required, e.g. postgres://user:password@127.0.0.1:5432/lovewall")
	}

	gormdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	g.UseDB(gormdb) // reuse your gorm db

	g.ApplyBasic(
		g.GenerateModel("goods"),
	)
	g.Execute()
}
