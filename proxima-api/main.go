package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)


type user struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    Email     string  `json:"email"`
    Username  string  `json:"username"`
    Pronouns  string  `json:"pronouns"`
    Secret    string  `json:"secret"`

    Visibility string `json:"visibility"`
}

var users = []user{
    {ID: "1", Name: "Klaire", Email: "no@csh.rit.edu", Username: "osmiu", Pronouns: "she/they", Secret: "no", Visibility: "USER_INVISIBLE"},
    {ID: "2", Name: "laire", Email: "no@csh.rit.edu", Username: "osmi", Pronouns: "she/they", Secret: "no", Visibility: "USER_INVISIBLE"},
    {ID: "3", Name: "aire", Email: "no@csh.rit.edu", Username: "osm", Pronouns: "she/they", Secret: "no", Visibility: "USER_INVISIBLE"},
}

// getAlbums responds with the list of all albums as JSON.
func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)

    router.Run("localhost:8080")
}

