package main

import (
    "log"
    "net/http"
    "encoding/json"
)

type Track struct {
  Id int `db:"TrackId" json:"id"`
  Title string `db:"Name" json:"title"`
  AlbumId int `db:"AlbumId" json:"albumId"`
  MediaTypeId int `db:"MediaTypeId" json:"mediaTypeId"`
  GenreId int `db:"GenreId" json:"genreId"`
  Composer string `db:"Composer" json:"composer"`
  Duration int `db:"Milliseconds" json:"duration"`
  Size int `db:"Bytes" json:"size"`
  Price int `db:"UnitPrice" json:"price"`
}

func getTracks(w http.ResponseWriter, r *http.Request){
  s.Header().Set("Content-Type", "application/json")
  
  title := r.URL.Query().Get("title")

  trackRows, err := db.Queryx("SELECT * FROM tracks WHERE name LIKE ?","%" + title + "%")
  if len(trackRows) == 0 {
    // send 204
    log.Fatal("no data found")
  } else if err != nil {
    // send 404
    log.Fatal("unable to get data", err)
  }
  defer trackRows.Close()
  
  var tracks []Track

  for trackRows.Next() {
    var track Track
    err = trackRows.StructScan(&track)
    if err != nil {
      // send 404
      log.Fatal("unable to read data", err)
    }

    tracks = append(tracks, track)
  }

  json.NewEncoder(w).Encode(tracks)
}

func handleRequests() {
  http.HandleFunc("/tracks", getTracks)
  log.Fatal(http.ListenAndServe(":4041", nil))
}

func main() {
  handleRequests()
}

