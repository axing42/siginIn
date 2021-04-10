package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

// 葫芦侠登录
func login(u, p string) string {
	loginUrl := "http://floor.huluxia.com/account/login/ANDROID/4.0?platform=2&gkey=000000&app_version=4.0.0.6.2&versioncode=20141433&market_id=floor_huluxia&_key=&device_code=%5Bw%5D02%3A00%3A00%3A00%3A00%3A00"
	uname := u
	pwd := p

	h := md5.New()
	h.Write([]byte(pwd))
	md5Pwd := hex.EncodeToString(h.Sum(nil))
	postData := url.Values{}

	postData.Add("account", uname)
	postData.Add("login_type", "2")
	postData.Add("password", md5Pwd)
	newRequest, err := http.NewRequest("POST", loginUrl, strings.NewReader(postData.Encode()))
	Must(err)
	newRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	do, err3 := (&http.Client{}).Do(newRequest.WithContext(context.TODO()))
	defer func() {
		err2 := do.Body.Close()
		Must(err2)
	}()
	Must(err3)
	if do.StatusCode != 200 {
		println("请求失败,状态码:", do.StatusCode)
		syscall.Exit(0)
	}
	var l LoginS
	all, _ := io.ReadAll(do.Body)
	_ = json.Unmarshal(all, &l)
	if l.Status != 1 {
		log.Fatalln(l)
	}
	return l.Key
}
func signIn(key string, catId int) {
	// 板块签到接口
	signUrl := "http://floor.huluxia.com/user/signin/ANDROID/4.0?platform=2&gkey=00000&app_version=4.0.0.6.3&versioncode=20141434&market_id=floor_baidu&device_code=%5Bw%5D02%3A00%3A00%3A00%3A00%3A00&_key=" + key + "&cat_id=" + strconv.Itoa(catId)
	get := MyCurl(signUrl)
	defer DClose(get)
}

// 获取所有有效分类ID
func category() (cIds []int) {
	var CategoryAPI = `http://floor.huluxia.com/category/list/ANDROID/2.0`
	get := MyCurl(CategoryAPI)
	defer DClose(get)
	var cId CategoryId
	all, _ := io.ReadAll(get.Body)
	_ = json.Unmarshal(all, &cId)
	for _, v := range cId.Categories {
		cIds = append(cIds, int(v.CategoryID))
	}
	// 获取不到天天酷跑分类 单独添加
	cIds = append(cIds, 11)
	// remove 之后是个新数组传进去再出来已经物是人非
	cidT := remove(cIds, 0)

	return cidT
}

// 完成一个账户所有签到
func working(key string, id []int) {
	var wg sync.WaitGroup
	wg.Add(len(id))
	for i := 0; i < len(id); i++ {
		go func(w *sync.WaitGroup, i int) {
			signIn(key, id[i])
			w.Done()
		}(&wg, i)
	}
	wg.Wait()
	return
}

var accounts Account

func init() {
	act := fileOperation("./config.json")
	_ = json.Unmarshal([]byte(act), &accounts)
}

// 完成所有账户签到
func workingAll() {
	ids := category()
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(w *sync.WaitGroup, i int) {
			working(login(accounts.Name[i], accounts.Pwd), ids)
			println("账号: ", accounts.Name[i], " 签到完成")
			w.Done()
		}(&wg, i)
	}
	wg.Wait()
}

// 去除slice某个重复值
func remove(a []int, target int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == target {
			a = a[:i+copy(a[i:], a[i+1:])]
		}
	}
	return a
}

func main() {
	var uname, pwd string
	var more uint
	flag.StringVar(&uname, "u", "", "账号默认为空")
	flag.StringVar(&pwd, "p", "", "密码默认为空")
	flag.UintVar(&more, "m", 0, "是否多账号")
	flag.Parse()
	switch more {
	case 0:
		workingAll()
	case 1:
		ids := category()
		working(login(uname, pwd), ids)
	case 2:
		ids := category()
		u := strings.Split(uname, ",")
		p := strings.Split(pwd, ",")
		var wg sync.WaitGroup
		wg.Add(len(u))
		for i := 0; i < len(u); i++ {
			go func(i int, w *sync.WaitGroup) {
				working(login(u[i], p[i]), ids)
				w.Done()
			}(i, &wg)
		}
		wg.Wait()
	}
}
