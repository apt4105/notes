module Main exposing (..)

import Time
import Html exposing (text)
import Browser
import Html as H
import Html.Attributes as A
import Html.Events as E



main =
    Browser.sandbox
        { init = init
        , update = update
        , view = view
        }

type alias Model =
    { user: User -- the current user
    , notes: List Note -- the user's notes
    }

init : Model
init =
    { user =
        { id = 0
        , name = "me"
        , email = "me@example.com"
        }
    , notes = []
    }

type Msg = Nil

update : Msg -> Model -> Model
update msg model = model

view : Model -> H.Html Msg
view model =
  viewUser model.user

type alias User =
    { id: Int
    , name: String
    , email: String
    }

viewUser : User -> H.Html Msg
viewUser u =
    H.text ("user " ++ u.name ++ " here")

type alias Collaboration =
    { noteId: Int
    , userId: Int
    , read: Bool
    , write: Bool
    , delete: Bool
    }

type alias Note =
    { id: Int
    , name: String
    , creator: User
    , created: Time.Posix
    , update: Time.Posix
    , collaborations: List Collaboration
    }

