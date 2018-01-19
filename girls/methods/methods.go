package methods

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "os"
    "image/png"
    "image/jpeg"
    "image"
    "time"

)


var filePath = "images/"
var Dir_name = "images"

func Make_dir(dir_name string){
    err := os.Mkdir(dir_name, 0777)
    if err != nil {
        fmt.Println("mkdir root failed!")
        return
    }
}

func file_exist(fileName string) bool {
    if _,ok:=os.Stat(fileName);ok == nil{
        return true
    }
    return false
}

func Substr(str string, start, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0

    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length

    if start > end {
        start, end = end, start
    }

    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}

func Get_image_list(url string){
    fmt.Println("get page link url==>", url)
    body:=get_url(url)
    if body == ""{
        return
        }
    reg := regexp.MustCompile("http://www.meizitu.com/a/[0-9]+.html")
    links:=reg.FindAllString(body, -1)
    // fmt.Println(links)
    get_image_link(links)
}

func get_image_link(links []string){
    for _, uri := range links{
        fmt.Println("Get images url, page link==>", uri)
            body:=get_url(uri)
            if ""==body{
                return
            }
        // fmt.Println(body)
        reg:=regexp.MustCompile("http://mm.chinasareview.com/wp-content/uploads/[^\\.]+\\.(jpg|png|gif)")
        images:=reg.FindAllString(body, -1)
        // fmt.Println(images)
        download_image(images)
    }
}

func download_image(images []string){
    for _,v:=range images{
        fmt.Println("Download image, url==>", v)
        
        imageType:=Substr(v, -2, 3)
        resp,ok:=http.Get(v)
        if nil!=ok{
            fmt.Println("Download image, url==>", v,"error")
            time.Sleep(time.Duration(600)*time.Second)
            continue
        }
        defer resp.Body.Close()
        flag:=false
        var iImage image.Image
        content,ok:=ioutil.ReadAll(resp.Body)
        body:=string(content)
        if imageType=="jpg"{
            iImage,ok=jpeg.Decode(strings.NewReader(body))
            flag=true
            if nil!=ok{
               continue
            }
        } else if imageType == "png"{
            iImage,ok=png.Decode(strings.NewReader(body))
            flag=true
            if nil!=ok{
                continue
            }
        }
        if flag{
            rect:=iImage.Bounds()
            if rect.Max.X < 200 || rect.Max.Y < 200{
            //只下载大图，小图跳过
                fmt.Println("Skip download image, url ==>", v)
                continue
            }
        }
      // body:=getUrl(v)
        if nil!=ok || "" == body{
            fmt.Println("content is null")
            continue
        }
        paths:=strings.Split(v,"/")
        len:=len(paths)
        fileName:=filePath + paths[len-4]+  paths[len-3]+  paths[len-2] +  paths[len-1]
        if file_exist(fileName){
            continue
        }
        f,ok:=os.Create(fileName)
        if ok!=nil{
            fmt.Println("open file error")
            return
        }
        defer f.Close()
        f.WriteString(body)
        fmt.Println("Download image, url==>", v ,"successful")
    }
}

func get_url(url string) (content string){
    client := &http.Client{}
    request,err_1 := http.NewRequest("GET", url, nil)
    if err_1 != nil {
        fmt.Println(err_1)
        return
    }
    // 设置一个模拟客户端，确认访问方式和网址
    var user_agent string = Choose_agent()
    // 从浏览器信息列表中随机选一个
    // fmt.Println(user_agent)
    request.Header.Add("User-Agent",user_agent)
    request.Header.Add("Referer", url)
    // 添加请求头，尤其是浏览器信息
    response, err_2 := client.Do(request)
    // 进行请求
    if err_2 != nil {
        fmt.Println(err_2)
        time.Sleep(time.Duration(600)*time.Second)
        return
    }
    defer response.Body.Close()
    data, err_4 := ioutil.ReadAll(response.Body)
    if err_4 != nil {
        fmt.Println("read error")
        return
    }
    // fmt.Println(string(data))
    // fmt.Println("==========================================================================================")
    content = string(data)
    return   
}

