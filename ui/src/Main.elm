module Main exposing (..)

import Browser
import Html as H
import Html.Attributes as A
import Html.Events as E

import Task
import Time
import DateFormat
import String

-- utils


main =
    Browser.element
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        }

-- model

type alias Model =
    { user: User -- the current user
    , notes: List Note -- the user's notes
    , timeZone: Time.Zone -- the user's time zone
    }

type alias User =
    { name: String
    , email: String
    }


type alias Note =
    { id: Int
    , name: String
    , creator: String -- username of creator
    , created: Time.Posix
    , updated: Time.Posix
    , collaborations: List Collaboration
    }

type alias Collaboration =
    { noteId: Int
    , userId: Int
    , read: Bool
    , write: Bool
    , delete: Bool
    }

mockModel : Model
mockModel =
    { user = mockUser
    , notes = [ mockNote, mockNote ]
    , timeZone = Time.utc
    }

mockUser : User
mockUser =
    { name = "me"
    , email = "me@example.com"
    }

mockNote : Note
mockNote  =
    { id = 0
    , name = "my note"
    , creator = "me"
    , created = Time.millisToPosix 0
    , updated = Time.millisToPosix 1
    , collaborations = []
    }


init : () -> (Model, Cmd Msg)
init _ =
    ( mockModel
    , Task.perform AdjustTimeZone Time.here
    )

-- update

type Msg
    = AdjustTimeZone Time.Zone

update : Msg -> Model -> (Model, Cmd Msg)
update msg model = (model, Cmd.none)

-- subscriptions

subscriptions : Model -> Sub Msg
subscriptions model = Sub.none

-- view

view : Model -> H.Html Msg
view model =
    H.div []
        [ viewUser model.user
        , H.h1 [] [ H.text "my notes" ]
        , H.div [] (List.map (\x -> viewNote model x) model.notes)
        ]


viewUser : User -> H.Html Msg
viewUser u =
    H.div []
        [ viewUserName u
        , H.br [] []
        , H.text (u.email)
        ]


viewUserName : User -> H.Html Msg
viewUserName u =
    H.a
        [ A.href ("/users/" ++ u.name) ]
        [ H.text (u.name) ]

viewNote : Model -> Note -> H.Html Msg
viewNote m n =
    H.div
        [ A.style "border" "solid black 1px"
        , A.style "display" "inline-block"
        , A.style "margin" "3px"
        , A.style "padding" "3px"
        ]
        [ H.h3
            []
            [ H.a
                [ A.href ("/notes/" ++ String.fromInt n.id) ]
                [ H.text n.name ]
            ]
        , H.span [] [ H.text "updated: " ]
        , H.span [] [ H.text (timeFormat m.timeZone n.updated) ]
        ]

timeFormat : Time.Zone -> Time.Posix -> String
timeFormat =
    DateFormat.format
        [ DateFormat.dayOfMonthNumber
        , DateFormat.text " "
        , DateFormat.monthNameFull
        , DateFormat.text " "
        , DateFormat.yearNumber
        ]

