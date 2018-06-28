package controllers

import (
        "fmt"
        "html/template"
        "net/http"

        "github.com/munsylol/guild/models"
)

func handleMakeNewsPost(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, "cupcake")
        if err != nil {
                fmt.Printf("Invalid session %v\n", session)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        var User models.BnetUser
        var Characters models.Characters
        var Guildinfo models.GuildInfo
        var url string
        if access_token, ok := session.Values["usercode"].(string); ok {
                url = "https://us.api.battle.net/account/user?access_token=" + access_token
                err := models.Get(url, &User)
                if nil != err {
                        fmt.Println("Error on " + url)
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }

                url = "https://us.api.battle.net/wow/user/characters?access_token=" + access_token
                err = models.Get(url, &Characters)
                if nil != err {
                        fmt.Println("Error on " + url)
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }
        }

        key := models.GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apikey")
        url = "https://us.api.battle.net/wow/guild/thrall/NoBelfsAllowed?fields=members&locale=en_US&apikey=" + key
        err = models.Get(url, &Guildinfo)

        switch r.Method {
        case "GET":
                data := struct {
                        Active string
                        User models.BnetUser
                        Characters models.Characters
                        CanPost bool
                }{
                        "new_news_get",
                        User,
                        Characters,
                        false,
                }

                fmt.Println("Looking for matches...")

                highestRank := -1
                for _, element := range Guildinfo.Members {
                        for _, e := range Characters.CharacterList {
                                if element.Character.Name == e.Name && element.Character.Realm == e.Realm && element.Character.Battlegroup == e.Battlegroup {
                                        fmt.Println(e.Name, e.Realm, e.Battlegroup)

                                        if -1 != highestRank || element.Rank < highestRank {
                                                highestRank = element.Rank
                                                data.CanPost = true
                                        }
                                }
                        }
                }
                if highestRank > CAN_MAKE_NEWS_POSTS {
                        data.CanPost = false
                }

                t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/new_news_get.html"))
                t.ExecuteTemplate(w, "base", data)
                break
        case "POST":
                r.ParseForm()

                data := struct {
                        Active string
                        User models.BnetUser
                        Characters models.Characters
                }{
                        "new_news_post",
                        User,
                        Characters,
                }

                gw, err := models.NewGuildWriter()
                if nil != err {
                        fmt.Fprintln(w, "shit:\n" + err.Error())
                        return
                }

                np := &models.NewsPost {
                        Title: r.FormValue("news_title"),
                        Body: r.FormValue("news_body"),
                }

                err = gw.WriteNewsPost(np, User)
                if nil != err {
                        fmt.Printf("Error:" + err.Error())
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }

                t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/new_news_post.html"))
                t.ExecuteTemplate(w, "base",  data)
                break
        default:
                fmt.Fprintln(w, "Sorry, nothing here!")
        }
}

