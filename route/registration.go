// ルーティング情報を登録する関数
package route

import (
	"github.com/julienschmidt/httprouter"

	"github.com/Dozi0116/go-nopaste/c"
)

func Registration(router *httprouter.Router) {
	router.GET("/", c.Hello)

	router.GET("/hoge", c.Hoge)
	router.GET("/page", c.Page)

	router.GET("/api/message", c.GetMessageList)
	router.POST("/api/message", c.PostMessage)
	router.GET("/api/message/id", c.GetMessageById)
}
