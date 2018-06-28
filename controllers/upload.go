package controllers

import(
	"fmt"
	"net/http"

	"time"
	"crypto/md5"
	"html/template"
	"io"
	"os"
	"strconv"

//	"github.com/munsylol/guild/models"
)

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
       fmt.Println("method:", r.Method)
       if r.Method == "GET" {
           crutime := time.Now().Unix()
           h := md5.New()
           io.WriteString(h, strconv.FormatInt(crutime, 10))
           token := fmt.Sprintf("%x", h.Sum(nil))

           t, _ := template.ParseFiles("./views/upload.html")
           t.Execute(w, token)
       } else {
           r.ParseMultipartForm(32 << 20)
           file, handler, err := r.FormFile("uploadfile")
           if err != nil {
               fmt.Println(err)
               return
           }
           defer file.Close()
           fmt.Fprintf(w, "%v", handler.Header)
           f, err := os.OpenFile("./views/uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
           if err != nil {
               fmt.Println(err)
               return
           }
           defer f.Close()
           io.Copy(f, file)
       }
}

/*
func upload(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cupcake")
        check(w, err)

        var User models.BnetUser
        var Characters models.Characters
        var Guildinfo models.GuildInfo
        var url string
        if access_token, ok := session.Values["usercode"].(string); ok {
                url = "https://us.api.battle.net/account/user?access_token=" + access_token
                err := models.Get(url, &User)
                check(w, err)

                url = "https://us.api.battle.net/wow/user/characters?access_token=" + access_token
                err = models.Get(url, &Characters)
                check(w, err)
        }

        gr, err := models.NewGuildReader()
        check(w, err)
        key, err := gr.GetBNetAPIKey()
        check(w, err)

        url = "https://us.api.battle.net/wow/guild/thrall/NoBelfsAllowed?fields=members&locale=en_US&apikey=" + key
        err = models.Get(url, &Guildinfo)
        check(w, err)

       if r.Method == "GET" {
           crutime := time.Now().Unix()
           h := md5.New()
           io.WriteString(h, strconv.FormatInt(crutime, 10))
           token := fmt.Sprintf("%x", h.Sum(nil))

	   combineTpl(w, token, "upload")
	   return
//           t, _ := template.ParseFiles("./views/upload.gtpl")
//           t.Execute(w, token)
       } else {
           r.ParseMultipartForm(32 << 20)
           file, handler, err := r.FormFile("uploadfile")
           if err != nil {
               fmt.Println(err)
               return
           }
           defer file.Close()
           fmt.Fprintf(w, "%v", handler.Header)
           f, err := os.OpenFile("./views/uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
           if err != nil {
               fmt.Println(err)
               return
           }
           defer f.Close()
           io.Copy(f, file)
       }
}
*/

