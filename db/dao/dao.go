package dao

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
    "SampleGoServer/db/utils"
    "SampleGoServer/db/models"
)

//User db methods
func HandleCreateUserError(err error) (int, error) {
    return -1, err
}
func HandleReadUserError(err error) (models.User, error) {
    var u models.User
    return u, err
}
func HandleValidateUserError(err error) (bool, error) {
    return false, err
}
func CreateUser(name, pword string) (int, error) {
    var err error
    db := utils.OpenMySQL("root", "dbpassword")
    bytes_pword := []byte(pword)
    hashed_pword, err := bcrypt.GenerateFromPassword(bytes_pword, bcrypt.DefaultCost)
    if err != nil { return HandleCreateUserError(err) }
    str_hash := string(hashed_pword[:])
    cmd := "INSERT INTO users " +
          "(name, date_created, is_admin) " +
          fmt.Sprintf("VALUES ('%s', CURDATE(), FALSE)", name)
    res, err := db.Exec(cmd)
    if err != nil { return HandleCreateUserError(err) }
    id, err := res.LastInsertId()
    if err != nil { return HandleCreateUserError(err) }
    cmd = "INSERT INTO user_auth " +
           "(hash, user_id) " +
           fmt.Sprintf("VALUES ('%s', '%d')", str_hash, id)
    db.Exec(cmd)
    return int(id), nil
}
func ReadUser(id int) (models.User, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, name, date_created, is_admin FROM users " +
           fmt.Sprintf("WHERE id=%d", id)
    var usr models.User
    err := db.QueryRow(cmd).Scan(&usr.ID, &usr.Name, &usr.Date_created, &usr.Is_admin)
    if err != nil { return HandleReadUserError(err) }
    return usr, err
}
func DeleteUser(id int) error {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := fmt.Sprintf("DELETE FROM users where id=%d", id)
    res, err := db.Exec(cmd)
    if res == nil || err != nil {
        return err
    }
    return nil
}
func ValidateUser(id int, pword string) (bool, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := fmt.Sprintf("SELECT hash FROM user_auth WHERE id=%d", id)
    var hash string
    err := db.QueryRow(cmd).Scan(&hash)
    if err != nil { return HandleValidateUserError(err) }
    bytes_hash := []byte(hash)
    err = bcrypt.CompareHashAndPassword(bytes_hash, []byte(pword))
    if err != nil {
        return false, nil
    }
    return true, nil
}

func CreatePost(body, thumbnail_url string, creator_id int) (int, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "INSERT INTO posts (body, thumbnail_url, creator_id, time_created) " +
           fmt.Sprintf("VALUES ('%s', '%s', %d, CURRENT_TIMESTAMP)", body, thumbnail_url, creator_id)
    res, err := db.Exec(cmd)
    if res == nil || err != nil {
        return -1, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return -1, err
    }
    return int(id), nil
}
func ReadPost(id int) (models.Post, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, body, thumbnail_url, creator_id, time_created FROM posts " +
           fmt.Sprintf("WHERE id=%d", id)
    var p models.Post
    err := db.QueryRow(cmd).Scan(&p.ID, &p.Body, &p.Thumbnail_url, &p.Creator_id, &p.Time_created)
    if err != nil {
        return p, err
    }
    return p, nil
}
func PostsFromUser(user_id int) ([]models.Post, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    var posts []models.Post
    _, err := ReadUser(user_id)
    if err != nil {
        return posts, err
    }
    cmd := "SELECT id, body, thumbnail_url, creator_id, time_created FROM posts " +
           fmt.Sprintf("WHERE creator_id=%d", user_id)
    rows, err := db.Query(cmd)
    if err != nil {
        return posts, err
    }
    for rows.Next() {
        var p models.Post
        err := rows.Scan(&p.ID, &p.Body, &p.Thumbnail_url, &p.Creator_id, &p.Time_created)
        if err != nil {
            return posts, err
        }
        posts = append(posts, p)
    }
    rows.Close()
    return posts, nil
}

func CreateComment(cid, pid int, body string) (int, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    _, err := ReadUser(cid)
    if err != nil {
        return -1, err
    }
    _, err = ReadPost(pid)
    if err != nil {
        return -1, err
    }
    cmd := "INSERT INTO comments (body, creator_id, post_id, time_created) " +
           fmt.Sprintf("VALUES ('%s', %d, %d, CURRENT_TIMESTAMP)", body, cid, pid)
    res, err := db.Exec(cmd)
    if res == nil || err != nil {
        return -1, err
    }
    id, _ := res.LastInsertId()
    return int(id), nil
}
func ReadComment(id int) (models.Comment, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    cmd := "SELECT id, body, creator_id, post_id, time_created FROM comments " +
           fmt.Sprintf("WHERE id=%d", id)
    var c models.Comment
    err := db.QueryRow(cmd).Scan(&c.ID, &c.Body, &c.Creator_id, &c.Post_id, &c.Time_created)
    if err != nil {
        return c, err
    }
    return c, nil
}
func CommentsFromPost(pid int) ([]models.Comment, error) {
    db := utils.OpenMySQL("root", "dbpassword")
    var comments []models.Comment
    _, err := ReadPost(pid)
    if err != nil { return comments, err }
    cmd := "SELECT id, body, creator_id, post_id, time_created FROM comments " +
           fmt.Sprintf("WHERE post_id=%d", pid)
    rows, err := db.Query(cmd)
    for rows.Next() {
        var c models.Comment
        err := rows.Scan(&c.ID, &c.Body, &c.Creator_id, &c.Post_id, &c.Time_created)
        if err != nil {
            return comments, err
        }
        comments = append(comments, c)
    }
    rows.Close()
    return comments, nil
}
