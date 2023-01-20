package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/muxCrud/utils"
)

type blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
type bloglist struct {
	blogStore []blog
}

func Service() *bloglist {
	return &bloglist{}
}
func (b *bloglist) Createblog(w http.ResponseWriter, r *http.Request) {

	if !utils.Checkmethod(r.Method, utils.POST) {
		utils.ResponseWriter(w, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	b.LoadFromJson()
	var newblog blog
	err := json.NewDecoder(r.Body).Decode(&newblog)
	if err != nil {
		log.Fatal("error found in decoding data")
		panic(err)
	}

	newblog.Id = b.AddnewId()
	b.blogStore = append(b.blogStore, newblog)
	b.SavetoJson()
	// marshal blog to write in response
	data, err := json.Marshal(&newblog)
	utils.Errorhandlefordataconversion(err)

	utils.ResponseWriter(w, http.StatusOK, data, "New Blog created Successfully:")
}
func (b *bloglist) Getblog(w http.ResponseWriter, r *http.Request) {
	if !utils.Checkmethod(r.Method, utils.GET) {
		utils.ResponseWriter(w, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	b.LoadFromJson()
	data, err := json.Marshal(&b.blogStore)
	utils.Errorhandlefordataconversion(err)
	utils.ResponseWriter(w, http.StatusOK, data, "Data Present is:")
}
func (b *bloglist) Deleteblog(w http.ResponseWriter, r *http.Request) {
	if !utils.Checkmethod(r.Method, utils.DELETE) {
		utils.ResponseWriter(w, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	b.LoadFromJson()
	id := utils.Getidfromurl(r.URL.Path)
	for key, v := range b.blogStore {

		if v.Id == id {
			b.blogStore = append(b.blogStore[:key], b.blogStore[key+1:]...)
			b.SavetoJson()
			utils.ResponseWriter(w, http.StatusOK, []byte(strconv.Itoa(id)), "Data Deleted with id is:")
			return
		}

	}
	utils.ResponseWriter(w, http.StatusOK, []byte(strconv.Itoa(id)), "No data deleted for the given id:")
}

func (b *bloglist) Updateblog(w http.ResponseWriter, r *http.Request) {
	if !utils.Checkmethod(r.Method, utils.PATCH) {
		utils.ResponseWriter(w, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	b.LoadFromJson()
	id := utils.Getidfromurl(r.URL.Path)

	var newblog blog
	err := json.NewDecoder(r.Body).Decode(&newblog)
	utils.Errorhandlefordataconversion(err)

	for key, v := range b.blogStore {

		if v.Id == id {
			v.Body = newblog.Body
			v.Title = newblog.Title
			b.blogStore[key] = v
			b.SavetoJson()
			utils.ResponseWriter(w, http.StatusOK, []byte(strconv.Itoa(id)), "Data Updates with id is:")
			return
		}

	}
	utils.ResponseWriter(w, http.StatusOK, []byte(strconv.Itoa(id)), "No updates for the given id:")

}

func (b *bloglist) Getblogbyid(w http.ResponseWriter, r *http.Request) {

	if !utils.Checkmethod(r.Method, utils.GET) {
		utils.ResponseWriter(w, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	b.LoadFromJson()
	id := utils.Getidfromurl(r.URL.Path)
	var newblog blog
	for _, v := range b.blogStore {

		if v.Id == id {
			newblog = v
			data, err := json.Marshal(&newblog)
			utils.Errorhandlefordataconversion(err)
			utils.ResponseWriter(w, http.StatusOK, data, "Data with the given id")
			return
		}

	}

	utils.ResponseWriter(w, http.StatusOK, nil, "No blog with the given id")

}
