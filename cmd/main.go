package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sadensmol/test_gymshark/internal/service"
	"github.com/sadensmol/test_gymshark/internal/store"
)

var router *chi.Mux
var packService *service.PackService

func main() {
	var err error

	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	packStore, err := store.NewPackStore()
	catch(err)
	packService, err = service.NewPackService(packStore)
	catch(err)

	router.Get("/", HandleMain)
	router.Delete("/pack", HandleDeletePack)
	router.Post("/pack", HandleAddPack)
	router.Post("/buy", HandleBuy)

	err = http.ListenAndServe(":8005", router)
	catch(err)
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, packService.AvailablePacks())
	catch(err)
}

func HandleDeletePack(w http.ResponseWriter, r *http.Request) {
	packSize := r.URL.Query().Get("s")
	ps, err := strconv.Atoi(packSize)
	catch(err)
	err = packService.Remove(ps)
	catch(err)
}
func HandleAddPack(w http.ResponseWriter, r *http.Request) {
	packSize := r.FormValue("size")
	ps, err := strconv.Atoi(packSize)
	catch(err)
	err = packService.Add(ps)
	catch(err)

	htmlStr := fmt.Sprintf("<tr id='packs-list-%d'><td>%d</td><td><p><button style='width: 40px;' hx-delete='/pack?s=%d' hx-target='#packs-list-%d' hx-swap='delete'>del</button></p></td></tr>", ps, ps, ps, ps)
	t, _ := template.New("t").Parse(htmlStr)
	err = t.Execute(w, nil)
	catch(err)
}

func HandleBuy(w http.ResponseWriter, r *http.Request) {
	numStr := r.FormValue("num")
	num, err := strconv.Atoi(numStr)
	catch(err)
	packs, err := packService.PackItems(num)
	catch(err)

	htmlStr := fmt.Sprintf("The following items were used to ship your purchase: <br/> %v <br/> Total packs: %d <br/> Total size: %d", packs, len(packs), packs.TotalSize())

	t, _ := template.New("t").Parse(htmlStr)
	err = t.Execute(w, nil)
	catch(err)
}
