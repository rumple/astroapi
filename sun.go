package main
import (
    "time"
    "net/http"
    "fmt"
    "strings"
    "strconv"
    "github.com/cpucycle/astrotime"
)

const ver = "v0"
const lenPath = 4

func SunHandlerV0(w http.ResponseWriter, r *http.Request) {
    var latitude, longitude *float64
    var lat, lon float64
    var err error
    if r.Method != "GET" {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }
    components := strings.Split(r.URL.Path[1:], "/")
    if len(components) != 4 {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }
    for _, comp := range components {
        switch {
            case len(comp) > 3 && comp[:3] == "lat":
                lat, err = strconv.ParseFloat(comp[3:], 64)
                if err == nil {
                  latitude = &lat
                } else {
                  http.Error(w, "Bad Request", http.StatusBadRequest)
                  return
                }
            case len(comp) > 3 && comp[:3] == "lon":
                lon, err = strconv.ParseFloat(comp[3:], 64)
                if err == nil {
                  longitude = &lon
                } else {
                  http.Error(w, "Bad Request", http.StatusBadRequest)
                  return
                }
        }
    }
    if latitude == nil || longitude == nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    if components[1] == "sunrise" {
        fmt.Fprint(w, astrotime.NextSunrise(time.Now(), *latitude, *longitude))
    }

    if components[1] == "sunset" {
        fmt.Fprint(w, astrotime.NextSunset(time.Now(), *latitude, *longitude))
    }
}

func init() {
    http.HandleFunc("/v0/", SunHandlerV0)
}
