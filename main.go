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

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of users, looking for
    // an user whose ID value matches the parameter.
    for _, a := range users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
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
    router.GET("/users/:id", getUserByID)
    router.POST("/users", postUsers)

    router.Run(":8080")
}

