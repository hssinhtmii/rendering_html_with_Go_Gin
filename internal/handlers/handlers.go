package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {

} 

var page = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func (server *Handler) Home(ctx *gin.Context) {
	err := RenderPage(ctx.Writer, "index.jet", nil)
	if err != nil {
		return
	}

}

// rendering the page
func RenderPage(w http.ResponseWriter, MyTemplate string, data jet.VarMap) error {
	mytmpl, err := page.GetTemplate(MyTemplate)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = mytmpl.Execute(w, data, nil)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
