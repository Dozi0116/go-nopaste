// ルーティング情報を登録する関数
package route

import (
	"net/http"

	"github.com/Dozi0116/go-nopaste/c"
)

func Registration() {
	http.HandleFunc("/", c.Hello)

	http.HandleFunc("/hoge", c.Hoge)
	http.HandleFunc("/page", c.Page)
}
