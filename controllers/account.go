package controllers

import (
    "bitbucket.org/pchandev/gamecatalogue-go/common"

    _ "bitbucket.org/pchandev/gamecatalogue-go/models"
    _ "bitbucket.org/pchandev/gamecatalogue-go/utilities/db_helper"
    _ "bitbucket.org/pchandev/gamecatalogue-go/utilities/random_generators"
    "fmt"
    "html/template"
    "net/http"
    _ "encoding/json"
    _ "io/ioutil"
    _ "reflect"
)

type AccountController struct {
    common.Controller
    controller_handlers []common.ControllerHandler
}

func ( this AccountController ) StartController() {
    this.controller_handlers = []common.ControllerHandler{
        AccountRootHandler{},
    }

    common.InitializeHandler( this.controller_handlers )
}

type AccountRootHandler struct {
    common.ControllerHandler
}

// root admin handler setup

func ( this AccountRootHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc( path_to_handle, this.HandlerFunction )

    return nil
}

func ( this AccountRootHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    view := &Page{
        Title: "Login Page",
        Body:  []byte("<h1>Login Page</h1>"),
    }
    
    parser := template.Must( template.ParseFiles( "../view/shared/index.html" ) )

    if err := parser.Execute( writer, view ); err != nil {
        fmt.Println( err )
    }
}

func ( this AccountRootHandler ) PathToHandle() string {
    return "/login/"
}

// end root setup
