package services 

import (
  "testing" 
  "log"
  "net/http"
)

func createHTTPServer() {
  s := &http.Server{
    Addr: "42069",
  }
}

func TestWeatherService(t *testing.T) {
  log.SetFlags(log.LstdFlags | log.Lshortfile)

    
}




