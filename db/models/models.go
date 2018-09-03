package models

import (
    "time"
)

type User struct {
    ID int
	Name string
    Date_created time.Time
    Is_admin bool
}

type Post struct {
    ID int
    Body string
    Thumbnail_url string
    Creator_id int
    Time_created time.Time
}

type Comment struct {
    ID int
    Body string
    Creator_id int
    Post_id int
    Time_created time.Time
}

type Error struct {
    Message string
}

type ID struct {
    ID int
}
