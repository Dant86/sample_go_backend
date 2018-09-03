package main

import (
    "C"
    "SampleGoServer/db/dao"
    "SampleGoServer/db/models"
    "encoding/json"
)

func SerializeError(err error) *C.char {
    error_ := models.Error{ Message: err.Error() }
    json_, jsonError := json.Marshal(error_)
    if jsonError != nil {}
    return C.CString(string(json_))
}

//export CreateUser
func CreateUser(name, pword string) *C.char {
    id, err := dao.CreateUser(name, pword)
    if err != nil {
        return SerializeError(err)
    }
    encoded, err := json.Marshal(models.ID{ ID: id })
    if err != nil {}
    return C.CString(string(encoded))
}

//export GetUser
func GetUser(id int) *C.char {
    usr, err := dao.ReadUser(id)
    if err != nil {
        return SerializeError(err)
    }
    encoded, err := json.Marshal(usr)
    if err != nil {}
    return C.CString(string(encoded))
}

//export DeleteUser
func DeleteUser(id int) *C.char {
    err := dao.DeleteUser(id)
    if err != nil {
        return SerializeError(err)
    }
    return C.CString("{}")
}

//export CreatePost
func CreatePost(body, thumbnail_url string, creator_id int) *C.char {
    id, err := dao.CreatePost(body, thumbnail_url, creator_id)
    if err != nil {
        return SerializeError(err)
    }
    encoded, err := json.Marshal(models.ID{ ID: id })
    if err != nil {}
    return C.CString(string(encoded))
}

//export GetPost
func GetPost(id int) *C.char {
    post, err := dao.ReadPost(id)
    if err != nil {
        return SerializeError(err)
    }
    encoded, err := json.Marshal(post)
    return C.CString(string(encoded))
}

//export GetPostsFromUser
func GetPostsFromUser(id int) *C.char {
    posts, err := dao.PostsFromUser(id)
    if err != nil {
        return SerializeError(err)
    }
    encoded := "["
    for ind, p := range posts {
        individEncoded, err := json.Marshal(p)
        if err != nil {}
        encoded += string(individEncoded)
        if ind != len(posts) - 1 {
            encoded += ", "
        }
    }
    encoded += "]"
    return C.CString(encoded)
}

//export CreateComment
func CreateComment(cid, pid int, body string) *C.char {
    id, err := dao.CreateComment(cid, pid, body)
    if err != nil {
        return SerializeError(err)
    }
    encoded, err := json.Marshal(models.ID{ ID: id })
    if err != nil {}
    return C.CString(string(encoded))
}

//export GetComment
func GetComment(id int) *C.char {
    comment, err := dao.ReadComment(id)
    if err != nil {
        return SerializeError(err)
    }
    encoded, err := json.Marshal(comment)
    return C.CString(string(encoded))
}

//export GetCommentsFromPost
func GetCommentsFromPost(id int) *C.char {
    comments, err := dao.CommentsFromPost(id)
    if err != nil {
        return SerializeError(err)
    }
    encoded := "["
    for ind, c := range comments {
        individEncoded, err := json.Marshal(c)
        if err != nil {}
        encoded += string(individEncoded)
        if ind != len(comments) - 1 {
            encoded += ", "
        }
    }
    encoded += "]"
    return C.CString(encoded)
}

func main() {}
