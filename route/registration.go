// ルーティング情報を登録する関数
package route

import (
	"net/http"

	"../c"
)

func Registration() {
	http.HandleFunc("/", c.Hello)

	http.HandleFunc("/hoge", c.Hoge)
}
