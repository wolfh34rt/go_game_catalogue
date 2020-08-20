package common

import (
    "net/http"
)

type Page struct {
    Title                           string
    Body                            []byte
    page_functions                  PageFunctions
}

type PageFunctions interface {
    Title()                         string
    Body()                          []byte
    LoadTemplates( string, []byte ) ( error, Page )
}

func (this *Page) InitPage( files_to_load []string ) ( error, Page )  {
    return this.page_functions.LoadTemplates( this.page_functions.Title(), this.page_functions.Body() )
}

type Controller interface {
    StartController()
    InitializeHandler()             error
}

type ControllerHandler interface {
    PathToHandle()                  string
    HandleHttp( string )            error
    HandlerFunction( http.ResponseWriter, *http.Request )
}

func InitializeHandler( controller_handlers []ControllerHandler ) {
    for _, handler := range controller_handlers {
        handler.HandleHttp( handler.PathToHandle() )
    }
}

func NewPage( files_to_load []string, new_page_functions PageFunctions ) Page {
    return Page{ Title: new_page_functions.Title(), Body: new_page_functions.Body(), page_functions: new_page_functions }
}
