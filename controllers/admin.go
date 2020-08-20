package controllers

import (
    "bitbucket.org/pchandev/gamecatalogue-go/common"

    "bitbucket.org/pchandev/gamecatalogue-go/models"
    "bitbucket.org/pchandev/gamecatalogue-go/utilities/db_helper"
    "bitbucket.org/pchandev/gamecatalogue-go/utilities/random_generators"
    "fmt"
    "html/template"
    "net/http"
    "encoding/json"
    _ "io/ioutil"
    _ "reflect"
)

type AdminController struct {
    common.Controller
    controller_handlers []common.ControllerHandler
}

func ( this AdminController ) StartController() {
    this.controller_handlers = []common.ControllerHandler{
        AdminRootHandler{},
        AdminGetItemsHandler{},
        AdminGetPlatformsHandler{},
        AdminGetMediaTypesHandler{},
        AdminPostOperationsHandler{},
    }

    common.InitializeHandler( this.controller_handlers )
}

type AdminRootHandler struct {
    common.ControllerHandler
}

type Page struct {
    Title string
    Body  []byte
}

// root admin handler setup

func ( this AdminRootHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc( path_to_handle, this.HandlerFunction )

    return nil
}

func ( this AdminRootHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    view := &Page{
        Title: "Admin Page",
        Body:  []byte("<h1>Admin Page</h1>"),
    }
    
    parser := template.Must( template.ParseFiles( "../view/shared/index.html", "../view/shared/login.html", "../view/admin/admin.html" ) )

    if err := parser.Execute( writer, view ); err != nil {
        fmt.Println( err )
    }
}

func ( this AdminRootHandler ) PathToHandle() string {
    return "/admin/"
}

// end root setup

// get items handler

type AdminGetItemsHandler struct {
    common.ControllerHandler
}

func ( this AdminGetItemsHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc( path_to_handle, this.HandlerFunction )

    return nil
}

func ( this AdminGetItemsHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    items := []models_db.Items{}
    db_map := utilities_db.InitializeDbMap()
    db_map.Select( &items, "SELECT * FROM items" )

    data_set := map[string]interface{} { 
        "rows" : items,
    }
    
    if json_object, err := json.Marshal( data_set ); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError )
        return
    } else {
        writer.Header().Set( "Content-Type", "application/json" )
        writer.Write( json_object )
    }
}

func ( this AdminGetItemsHandler ) PathToHandle() string {
    return "/admin/get_items"
}

// end handler get setup

// get platforms handler

type AdminGetPlatformsHandler struct {
    common.ControllerHandler
}

func ( this AdminGetPlatformsHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc( path_to_handle, this.HandlerFunction )

    return nil
}

func ( this AdminGetPlatformsHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    items := []models_db.Platforms{}
    db_map := utilities_db.InitializeDbMap()
    db_map.Select( &items, "SELECT * FROM platforms" )
    data_set := make( map[string]string )

    for _, value := range items {
        data_set[ value.Platformid ] =  value.Platformname
    }

   
    if json_object, err := json.Marshal( data_set ); err != nil {
        http.Error( writer, err.Error(), http.StatusInternalServerError )
        return
    } else {
        writer.Header().Set( "Content-Type", "application/json" )
        writer.Write( json_object )
    }
}

func ( this AdminGetPlatformsHandler ) PathToHandle() string {
    return "/admin/get_platforms"
}

// end handler get setup

// get media types handler

type AdminGetMediaTypesHandler struct {
    common.ControllerHandler
}

func ( this AdminGetMediaTypesHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc( path_to_handle, this.HandlerFunction )

    return nil
}

func ( this AdminGetMediaTypesHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    items := []models_db.Mediatypes{}
    db_map := utilities_db.InitializeDbMap()
    db_map.Select( &items, "SELECT * FROM mediatypes" )
    data_set := make( map[string]string )

    for _, value := range items {
        data_set[ value.Mediatypeid ] = value.Mediatypename
    }
    
    if json_object, err := json.Marshal( data_set ); err != nil {
        http.Error( writer, err.Error(), http.StatusInternalServerError )
        return
    } else {
        writer.Header().Set( "Content-Type", "application/json" )
        writer.Write( json_object )
    }
}

func ( this AdminGetMediaTypesHandler ) PathToHandle() string {
    return "/admin/get_mediatypes"
}

// end handler get setup

// admin edit handler setup

type AdminPostOperationsHandler struct {
    common.ControllerHandler
}

func ( this AdminPostOperationsHandler ) HandleHttp( path_to_handle string ) error {
    http.HandleFunc(path_to_handle, this.HandlerFunction)

    return nil
}

func ( this AdminPostOperationsHandler ) HandlerFunction( writer http.ResponseWriter, request *http.Request ) {
    result_object := make( map[string]string )

    if request.Method == "POST" {
        request.ParseForm()
        db_map := utilities_db.InitializeDbMap()
        operation := request.Form["oper"][0]

        if operation == "edit"  {
            item_id := request.Form["Itemid"][0]
            item_name := request.Form["Itemname"][0]
            media_type := request.Form["Mediatype"][0]
            platform := request.Form["Platform"][0]
            
            items := []models_db.Items{}
            
            db_map.Select( &items, "SELECT * FROM items WHERE itemid = '" + item_id + "'" )

            if( len( items ) == 1 ) {
                items[0].Itemname = item_name
                items[0].Mediatype = media_type
                items[0].Platform = platform

                db_map.Update( &items[0] )

                result_object["result"] = "true"
                result_object["message"] = "Item Successfully Edited!"
            }
        } else if operation == "add" {
            var item models_db.Items
            item_name := request.Form["Itemname"][0]
            media_type := request.Form["Mediatype"][0]
            platform := request.Form["Platform"][0]

            item.Itemid = string_generator.RandomId()
            item.Itemname = item_name
            item.Mediatype = media_type
            item.Platform = platform

            db_map.Insert( &item )

            result_object["result"] = "true"
            result_object["message"] = "Item Successfully Added!"
        } else if operation == "del" {
            items := []models_db.Items{}
            item_id := request.Form["id"][0]
            db_map.Select( &items, "SELECT * FROM items WHERE itemid = '" + item_id + "'" )

            if( len( items ) == 1 ) {
                db_map.Delete( &items[0] )

                result_object["result"] = "true"
                result_object["message"] = "Item Successfully Deleted!"
            }
        }
    }

    if json_object, err := json.Marshal( result_object ); err != nil {
        http.Error( writer, err.Error(), http.StatusInternalServerError )
        return
    } else {
        writer.Header().Set( "Content-Type", "application/json" )
        writer.Write( json_object )
    }
}

func ( this AdminPostOperationsHandler ) PathToHandle() string {
    return "/admin/post_operations"
}

// end admin edit handler setup