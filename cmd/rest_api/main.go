package main

import (
	"flag"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/gikslab_assesment/internal/app/config"
	"github.com/kbiits/gikslab_assesment/internal/app/handler"
	"github.com/kbiits/gikslab_assesment/internal/app/http"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
	repoActivity "github.com/kbiits/gikslab_assesment/internal/app/repository/activity"
	repoRedis "github.com/kbiits/gikslab_assesment/internal/app/repository/redis"
	repoSkill "github.com/kbiits/gikslab_assesment/internal/app/repository/skill"
	repoUser "github.com/kbiits/gikslab_assesment/internal/app/repository/user"
	ucActivity "github.com/kbiits/gikslab_assesment/internal/app/usecase/activity"
	ucAuth "github.com/kbiits/gikslab_assesment/internal/app/usecase/auth"
	ucSkill "github.com/kbiits/gikslab_assesment/internal/app/usecase/skill"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {

	cfgPath := flag.String("config_file", "./config.json", "specify path to the config file")

	cfg, err := config.ParseConfig(*cfgPath)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sqlx.Connect("postgres", cfg.DbConfig.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	redisOpts, err := redis.ParseURL(cfg.Redis.Url)
	if err != nil {
		log.Fatalln(err)
	}

	rdb := redis.NewClient(redisOpts)
	redisRepo, err := repoRedis.NewRedisRepository(rdb)
	if err != nil {
		log.Fatalln(err)
	}

	jwt := jwt.New(cfg.Jwt, redisRepo)

	userRepo := repoUser.New(db)
	skillRepo := repoSkill.New(db)
	activityRepo := repoActivity.New(db)

	authUc := ucAuth.NewAuthUsecase(userRepo, skillRepo, jwt, redisRepo)
	skillUc := ucSkill.NewSkillUsecase(skillRepo)
	activityUc := ucActivity.NewActivityUsecase(skillRepo, userRepo, activityRepo)

	handlers := handler.NewHandlers(authUc, skillUc, activityUc)

	httpApp, err := http.NewHttpApp(cfg, handlers, jwt)

	if err != nil {
		log.Fatalln(err)
	}

	err = httpApp.Start()
	if err != nil {
		log.Fatalln(err)
	}

	// shutdown here
	db.Close()
}
