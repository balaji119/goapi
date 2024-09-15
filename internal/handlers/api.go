package hndlers

import(
        "github.com/go-chi/chi"
        chimiddle "github.com/go-chi/chi/middleware"
        "github.com/balaji119/goapi/internal/middelware"
)


func Handler(r *chi.Mux){
        r.User(chimiddle.StripSlashes)
        r.Route("/account", func(router chi.Router){

                router.Use(middleware.Authorization)

                router.Get("/coins", GetCoinBalance)
        })
}
