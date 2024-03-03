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

func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
    var newUser user

    // Call BindJSON to bind the received JSON to
    // newUser.
    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    // Add the new user to the slice.
    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)
    router.POST("/users", postUsers)

    router.Run("localhost:8080")
}

