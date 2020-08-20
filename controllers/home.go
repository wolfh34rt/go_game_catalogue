package controllers

import (
    "bitbucket.org/pchandev/gamecatalogue-go/common"

    "html/template"
    "net/http"
    "fmt"
)

type HomeController struct {
    common.Controller
    controller_handlers     []common.ControllerHandler
}

func ( this HomeController ) StartController() {
    this.controller_handlers = []common.ControllerHandler{ HomeRootHandler{} }
    common.InitializeHandler( this.controller_handlers )
}

type HomeRootHandler struct {
    common.ControllerHandler
}

type TestPage struct {
    Title           string
    Body            []byte
}

func ( this HomeRootHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc( path_to_handle, this.HandlerFunction )
    return nil
}

func ( this HomeRootHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    view := &TestPage {
        Title: "Test Page",
        Body: []byte( "<h1>TEST</h1>" ),
    }

    parser := template.Must( template.ParseFiles( "../view/shared/index.html", "../view/home/index.html", "../view/shared/login.html" ) );

    if err := parser.Execute( writer, view ); err != nil {
        fmt.Println( err )
    }
}

func ( this HomeRootHandler ) PathToHandle() string {
    return "/"
}
