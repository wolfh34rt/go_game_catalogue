package main

import (
    _ "fmt"
    "net/http"
    _ "strconv"

    _ "bitbucket.org/pchandev/gamecatalogue-go/models"
    "bitbucket.org/pchandev/gamecatalogue-go/controllers"
    "bitbucket.org/pchandev/gamecatalogue-go/common"
    _ "bitbucket.org/pchandev/gamecatalogue-go/utilities/db_helper"
    _ "bitbucket.org/pchandev/gamecatalogue-go/utilities/random_generators"
    _ "bitbucket.org/pchandev/gamecatalogue-go/_vendor/github.com/jmoiron/modl"
)

func main() {
    controllers := []common.Controller{
        controllers.HomeController{},
        controllers.AdminController{},
    }

    for _, controller := range controllers {
        controller.StartController()
    }

    // Add static global handlers
    // TODO: Need to move into static handlers class.
    http.Handle( "/content/", http.StripPrefix( "/content/", http.FileServer( http.Dir( "../content" ) ) ) )
    http.Handle( "/scripts/", http.StripPrefix( "/scripts/", http.FileServer( http.Dir( "../scripts" ) ) ) )
    http.Handle( "/images/", http.StripPrefix( "/images/", http.FileServer( http.Dir( "../images" ) ) ) )

    // listen on specific port
    http.ListenAndServe( ":8080", nil )
}
