package routes

import (
	"algo/trees"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SafeTree struct {
	avl   *trees.AVL
	mutex sync.Mutex
}

var avlTree = &SafeTree{
	avl: &trees.AVL{},
}

func addTreeRoutes(rg *gin.RouterGroup) {
	treesGroup := rg.Group("trees")

	treesGroup.GET("/", func(c *gin.Context) {
		avlTree.mutex.Lock()
		hierarchy := avlTree.avl.Print()
		avlTree.mutex.Unlock()
		c.JSON(http.StatusOK, hierarchy)
	})

	treesGroup.POST("/", func(c *gin.Context) {
		var data trees.Data
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(data)

		avlTree.mutex.Lock()
		avlTree.avl.Push(data)
		avlTree.mutex.Unlock()

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

}
