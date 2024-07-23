package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/spells", app.spells)
	router.HandlerFunc(http.MethodPost, "/spells", app.updateSpellSlot)
	router.HandlerFunc(http.MethodGet, "/edit", app.edit)
	router.HandlerFunc(http.MethodGet, "/class/traits", app.classTraits)
	router.HandlerFunc(http.MethodGet, "/features", app.features)

	// I don't think these should be get requests since they change server state.
	// we can fix this with htmx
	router.HandlerFunc(http.MethodGet, "/save", app.save)
	router.HandlerFunc(http.MethodGet, "/load", app.load)

	router.HandlerFunc(http.MethodPost, "/hitdice", app.useHitDice)
	router.HandlerFunc(http.MethodPost, "/moxie", app.useMoxie)
	router.HandlerFunc(http.MethodPost, "/resetmoxie", app.resetMoxie)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)

}
