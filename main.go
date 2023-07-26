// main.go
package main
import (
   "github.com/gin-gonic/gin"
   "gitlab.com/alistairr/issue-tracker/db"
   "log"
   "net/http"
)
var (
   ListenAddr = "localhost:8080"
   RedisAddr = "localhost:6379"
)

func main() {
    database, err := db.NewDatabase(RedisAddr)
    if err != nil {
        log.Fatalf("Failed to connect to redis: %s", err.Error())
    }

    router := gin.Default()

    router.POST("/issue", func(c *gin.Context) {
        var issue Issue
        if err := c.ShouldBindJSON(&issue); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        issueID := database.CreateIssue(&issue)

        c.JSON(http.StatusOK, gin.H{"status": "Created", "id": issueID})
    })

    router.GET("/issue/:id", func(c *gin.Context) {
        issueID := c.Param("id")
        issue := database.GetIssue(issueID)

        c.JSON(http.StatusOK, issue)
    })

    router.PATCH("/issue/:id", func(c *gin.Context) {
        issueID := c.Param("id")

        var updatedFields map[string]interface{}
        if err := c.ShouldBindJSON(&updatedFields); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        database.UpdateIssue(issueID, updatedFields)

        c.JSON(http.StatusOK, gin.H{"status": "Updated", "id": issueID})
    })

    router.DELETE("/issue/:id", func(c *gin.Context) {
        issueID := c.Param("id")
        database.DeleteIssue(issueID)

        c.JSON(http.StatusOK, gin.H{"status": "Deleted", "id": issueID})
    })

    router.Run(ListenAddr)
}

