package main

import (
    "log"
    "net/http"
)

func main() {
    //ディレクトリを指定する
    fs := http.FileServer(http.Dir("/Users/akiho/Documents/seminer/2020_spring/lecture"))
    //ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
    http.Handle("/", fs)

    log.Println("Listening...")
    // 3000ポートでサーバーを立ち上げる
    http.ListenAndServe(":8000", nil)
}

// package main
//
// import (
//   'fmt'
//   'log'
//   'net/http'
//   'net/http/httputil'
// )
//
//
// func handler(w http.ResponseWriter, r *http.Request){
//   dump, err := httputil.DumpRequest(r, true)
//   if err ! = nil{
//       http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//       return
//   }
//   fmt.Println'string(dump))
//   fmt.Fprintf(w, '<html><body>hello</body></html>\n')
//
// }
//
// func main(){
//   var httpServer http.Server
//   http.HandleFunc('/',handler)
//   log.Println('start http listening :800')
//   httpServer.Addr = ':800'
//   log.Println(httpServer.ListenAndServer())
// }



// package main
//
// import (
//     "log"
//     "net/http"
//     "encoding/json"
// )
//
//
// type Intro struct {
//     Name        string  `json:"name"`
//     Kg          string  `json:"kg"`
//     LoginName   string  `json:"login_name"`
// }
//
// func Root(w http.ResponseWriter, r *http.Request) {
//     data := Intro{
//         Name: "Akiho Iwata",
//         Kg: "d-hacks",
//         LoginName: "akiho",
//     }
//
//     json.NewEncoder(w).Encode(data)
// }
//
// func main() {
//     http.HandleFunc("/", Root)
//
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }
