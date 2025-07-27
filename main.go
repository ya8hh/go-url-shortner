package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//struct to store data on run time using it as mongo db
type URL struct {
	ID string `json:"id"`
	OrginalURL string `json:"original_url"`
	ShortURL string `json:"short_url"`
	CreationDate time.Time  `json:"creation_date"`

}

var urlDB =make(map[string]URL)
//function to create short url
func generateShortUrl( OrginalURL string)string {
	hasher:= md5.New()
	hasher.Write([]byte(OrginalURL)) //converts the orignal string in byte array or in golang we say bytes
	data:= hasher.Sum(nil)
	hash :=hex.EncodeToString(data)
	return hash[:8];
}
//function to store url in run time Db
func createUrl(OrginalURL string) string{
	shortURL :=generateShortUrl(OrginalURL)
	id:=shortURL
	urlDB[id] = URL{
		ID: id,
		OrginalURL: OrginalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}
//to retrive the orignal url 
func getUrl( id string) (URL,error) {
	url ,ok:= urlDB[id]
	if !ok {
		return URL{},errors.New("Url not found")
	}
	return url,nil
}
//handler
func handler( w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello yash")
	}

	//handing short url to return the actual url 
	func shortURLHandler(w http.ResponseWriter,r *http.Request){
		var data struct {
			URL string `json:"url"`
		}
		err:=json.NewDecoder(r.Body).Decode(&data)
		if err!=nil{
			http.Error(w,"Invalid request body",http.StatusBadRequest)
			return
		}
		shortUrl:= createUrl(data.URL)
		fmt.Fprintf(w,shortUrl)
	}
func main() {
	//fmt.Println("Straing Projects")
	//starting the http server
	fmt.Println("server started and running on the port number 8080")
	http.HandleFunc("/",handler)
	http.HandleFunc("/shorten",shortURLHandler)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		fmt.Println("error in running function",err)
	}
}          