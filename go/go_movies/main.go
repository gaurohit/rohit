package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}
type Director struct{
 Firstname string `json:"firstname"`
 Lastname string `json:"lastname"`
}
var movies []Movie

func main()  {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID:"1",Isbn: "3562",Title: "ro",Director: &Director{Firstname:"did",Lastname:"ro"}})
	movies = append(movies, Movie{ID:"2",Isbn: "7433",Title: "to",Director: &Director{Firstname:"ste",Lastname:"sm"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("staring server")

	log.Fatal(http.ListenAndServe(":8000",r))
}

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type","application/json") //set json type

	json.NewEncoder(w).Encode(movies) //sending json respone
}

func deleteMovie(w http.ResponseWriter , r *http.Request)  {
	w.Header().Set("Content-type","application/json") 
	params := mux.Vars(r) // looking for id 
    for index, movie := range movies{
		if movie.ID == params["id"]{
           movies = append(movies[:index],movies[index+1:]... )
		   break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type","application/json") 
    params := mux.Vars(r)

	for _,movie := range movies{
		if movie.ID == params["id"]{
          json.NewEncoder(w).Encode(movie)
		  return
		}
	}
	
}

func createMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-type","application/json") 
	var movie Movie
	 _ = json.NewDecoder(r.Body).Decode(&movie) // new movie come is inside &movie
    movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies,movie)
    json.NewEncoder(w).Encode(movie)

}

func updateMovies(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-type","application/json") 
	params := mux.Vars(r)

	for index,movie := range movies{
		if movie.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			var item Movie
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.ID = params["id"]
			movies = append(movies, item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}


