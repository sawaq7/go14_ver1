///
/// 静水圧　U字管の計箁E   type2　�E�ファイル入力！E///

package main

import (
	    "fmt"
	    "os"
	    "strconv"
	    "strings"
	    "bufio"
	    "github.com/sawaq7/go14_ver1/client/tokura/suiri"
	    "github.com/sawaq7/go14_ver1/basic/rw"
	    "github.com/sawaq7/go14_ver1/basic/maths/sum"
    	                 )

func main() {

// 単位容積重量　�E�ω）をセチE��

   var omega ,drad1 ,drad2 ,press1 ,press2,high ,area1 ,area2 float64
   var fname ,fname2 ,cdmy string
   var index ,num int

// ファイルオープン

   fname = "C:/Go_Original/src/github.com/sawaq7/go14_ver1/client/tokura/file/seisui_inf.txt"
   fname2 = "C:/Go_Original/src/github.com/sawaq7/go14_ver1/client/tokura/file/seisui.txt"

   ad_fdata := make([]float64 ,6)        // keep work data for etc float data

// 静水惁E��ファイル、オープン

   file , err := os.Open(fname)

// ファイルリーダーをＧ�E��E�
   reader := bufio.NewReaderSize(file, 4096)

// 静水ファイル、オープン

   writer , err := os.OpenFile(fname2, os.O_CREATE|os.O_RDWR, 0666)
   if err !=nil {
      fmt.Println ("seisui2 open error" ,err)
      os.Exit(1)
   }

   index = 0        // レコードカウンターをinitialize

   for {

      index ++     // レコードカウンターをカウンチE
      fmt.Println ("main_seisui2　index " ,index)  // チE��チE��

// ファイルを１行read

      line ,_  := reader.ReadString('\n')

//斁E��単位にスペ�Eスで刁E��

      str := strings.Fields(line)

      num = len(str)

      fmt.Println ("main_seisui2　num " ,num)  // チE��チE��

      if num == 0 {  //　END　チェチE��

         fmt.Println ("main_seisui2 normal end")  // チE��チE��
         goto END
      }

      if index != 1{   // 見�Eし以外をmake

         omega ,_ =strconv.ParseFloat(str[0],64)
         drad1 ,_ =strconv.ParseFloat(str[1],64)
         drad2 ,_ =strconv.ParseFloat(str[2],64)
         press1 ,_ =strconv.ParseFloat(str[3],64)
         cdmy = str[4]
         high ,_ =strconv.ParseFloat(str[5],64)

         fmt.Println ( "main_seisui2 file data " ,omega, drad1 , drad2 ,press1 ,cdmy ,high ) // チE��チE��

// U字管の面積を計算すめE
         area1 = sum.Circle_Area(drad1/2)
         area2 = sum.Circle_Area(drad2/2)

         press2 =  suiri.Seisui1( area1 ,area2  ,press1  ,omega  ,high  )

         fmt.Println("main_seisui2 圧力�E",press2,"�E�E)   //チE��チE��

// U字管の面積をrewrite
         ad_fdata[0] = omega
         ad_fdata[1] = drad1
         ad_fdata[2] = drad2
         ad_fdata[3] = press1
         ad_fdata[4] = press2
         ad_fdata[5] = high
         rw.Wrline1(  writer , ad_fdata )

      }
   }
   //  final process ,"example file delete ,rename ,close"

   END :

   defer file.Close()
   defer writer.Close()

   return
}






