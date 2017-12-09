package main 

import (
   "./methods"
   "fmt"
   "time"
   )
func main() {
   methods.Make_dir(methods.Dir_name)
   fms:="http://www.meizitu.com/a/more_%d.html"
   max_page:=72
   cur_page:=1
   offset:=cur_page+max_page
   // ch:=make(chan int, max_page)
   for ;cur_page<offset;cur_page++{
      
      url:=fmt.Sprintf(fms, cur_page)
      fmt.Println("Parse url:",url)
      methods.Get_image_list(url)
      time.Sleep(time.Duration(2)*time.Second)
   }
   
   // sum:=0
   // forEnd:
   // for{
   //    select{
   //    case <- ch:
   //       sum+=1
   //       if sum == max_page{
   //          break forEnd
   //       }
   //    }
   // }
   fmt.Println("done!")
}