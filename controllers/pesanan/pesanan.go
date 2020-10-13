package pesanan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"restoran/conn"
	pesanan "restoran/models/pesanan"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//PesananCollection declaration
const PesananCollection = "pesanan"

//CreatePesanan endpoint
func CreatePesanan(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(body)
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	pesanan := pesanan.Pesanan{}
	err := c.Bind(&pesanan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Bad Request"})
		return
	}
	pesanan.CreatedAt = time.Now()
	err = db.C(PesananCollection).Insert(pesanan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Bad Request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "pesanan": &pesanan})
}

//GetPesanan endpoint
func GetPesanan(c *gin.Context) {
	// Get DB from Mongo Config
	var id, errParse = strconv.ParseInt(c.Param("id"), 10, 64)
	if errParse != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
	}
	pesanan, err := pesanan.PesananInfo(id, PesananCollection)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "pesanan": &pesanan})
}
