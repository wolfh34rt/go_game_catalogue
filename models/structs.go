package models_db

import (
    "time"
)

func GetAllDbStructs() map[string]DatabaseInfo {
    return map[string]DatabaseInfo {
        "borrowed_items" : DatabaseInfo{
            DatabaseObject : BorrowedItems{},
            DatabaseKey : "Borrowid",
            IsAutoIncrementKey : false,
        },
        "items" : DatabaseInfo{
            DatabaseObject : Items{},
            DatabaseKey : "Itemid",
            IsAutoIncrementKey : false,
        },
        "mediatypes" : DatabaseInfo{
            DatabaseObject : Mediatypes{},
            DatabaseKey : "Mediatypeid",
            IsAutoIncrementKey : false,
        },
        "oauthmembership" : DatabaseInfo{
            DatabaseObject : Oauthmembership{},
            DatabaseKey : "Oathmembershipid",
            IsAutoIncrementKey : false,
        },
        "platforms" : DatabaseInfo{
            DatabaseObject : Platforms{},
            DatabaseKey : "Platformid",
            IsAutoIncrementKey : false,
        },
        "roles" : DatabaseInfo{
            DatabaseObject : Roles{},
            DatabaseKey : "Roleid",
            IsAutoIncrementKey : false,
        },
        "userprofile" : DatabaseInfo {
            DatabaseObject : Userprofile{},
            DatabaseKey : "Profileid",
            IsAutoIncrementKey : false,
        },
        "usersinroles" : DatabaseInfo{
            DatabaseObject : Usersinroles{},
            DatabaseKey : "Usersinroleid",
            IsAutoIncrementKey : false,
        },
    }
}

type DatabaseInfo struct {
    DatabaseObject          interface{}
    DatabaseKey             string
    IsAutoIncrementKey      bool
}

type BorrowedItems struct {
    Borrowid            string      `db:"borrowid"`
    Userid              string      `db:"userid"`
    Itemid              string      `db:"itemid"`
    Borrowdate          time.Time   `db:"borrowdate"`
}

type Items struct {
    Itemid              string      `db:"itemid"`
    Itemname            string      `db:"itemname"`
    Mediatype           string      `db:"mediatype"`
    Platform            string      `db:"platform"`
}

type Mediatypes struct{
    Mediatypeid         string      `db:"mediatypeid"`
    Mediatypename       string      `db:"mediatypename"`
}

type Oauthmembership struct{
    Oathmembershipid    string      `db:"oathmembershipid"`
    Provider            string      `db:"provider"`
    Provideruserid      string      `db:"provideruserid"`
    Userid              string      `db:"userid"`
}

type Platforms struct{
    Platformid          string      `db:"platformid"`
    Platformname        string      `db:"platformname"`
}

type Roles struct{
    Roleid              string      `db:"roleid"`
    Rolename            string      `db:"rolename"`
}

type Userprofile struct{
    Profileid           string      `db:"profileid"`
    Userid              string      `db:"userid"`
    Email               string      `db:"email"`
}

type Usersinroles struct{
    Usersinroleid       string      `db:"usersinroleid"`
    Userid              string      `db:"userid"`
    Roleid              string      `db:"roleid"`
}