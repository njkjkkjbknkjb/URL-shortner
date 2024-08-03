package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
  
)
type URL struct{
  ID  string `json:"id"`
  OrginalURL string `json:"Orginal_URL"`
  ShortURL string  `json:"short_url"`
  CreateionDate time.Time `json:"creation_Date"`
  
}
// 
var urlDB = make(map[string]URL)
func genShortURL(OrginalURL string) string {
  hasher := md5.New()
  hasher.Write([]byte(OrginalURL)) // It convehrts rthe orginal string to a byte slice
  fmt.Println("hasher: ", hasher)
  data := hasher .Sum(nil)
   fmt.Println("hasher data: ", hasher)
  hash := hex.EncodeToString(data)
  fmt.Println("EncodeString: ", hash )
  fmt.Println("final string:", hash[:8] )
  return hash[:8];
}

func createURL(orginalURL string) {
  shortURL := genShortURL(orginalURL)
  id := shortURL
  urlDB[id] =URL{
    ID : id,
    OrginalURL: orginalURL,
    ShortURL: shortURL,
    CreateionDate:time.Now(),
  }
  return shortURL
}
func getURL(id string) (URL, error) {
  url, ok := urlDB[id]
  if !ok {
    return URL{}, errors.New("URL not found")
  }
  return url,nil
}
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "helloworld")
}
func main() {
	fmt.Println("Starting url shortner")
  OriginalURL :="https://github.com/njkjkkjbknkjb"
  genShortURL(OriginalURL)
  
  http.HandleFunc("/", handler)

  // Start the http server on port 8080
  err := http.ListenAndServe(":3000", nil)
  if err != nil {    fmt.Println("Error on starting server:", err)
  }
}
