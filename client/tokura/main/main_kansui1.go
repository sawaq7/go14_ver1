    ///////////////////////////////////////////////////////
   ///                main_kansui1                     ///
  ///     管水路　エネルギー線�E動水勾配線を描く        ///
 ///          チE�Eタはstring垁E                       ///
////////////////////////////////////////////////////////

package main

import (
	     "fmt"
	     "os"
	     "strings"
	     "bufio"
	     "github.com/sawaq7/go14_ver1/client/tokura/suiri"
	     "github.com/sawaq7/go14_ver1/basic/rw"
	                  )

func main() {

// 単位容積重量　�E�ω）をセチE��

   var fname ,fname2 ,line  string
   var index  ,num int

// ファイルオープン


   fname  = "C:/Go_Original/src/github.com/sawaq7/go14_ver1/client/tokura/file/water_inf.txt"
   fname2  = "C:/Go_Original/src/github.com/sawaq7/go14_ver1/client/tokura/file/grade_line.txt"

// 水路ファイル、オープン

   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)
//   freader , _ := os.Open(fname)

// 導水勾配線ファイル、オープン

   writer , _ := os.OpenFile(fname2, os.O_CREATE|os.O_RDWR, 0666)
//   writer , _ := os.Open(fname2)
   index = 0        // レコードカウンターをinitialize

// ファイルリーダーをＧ�E��E�
   reader := bufio.NewReaderSize(freader, 4096)

   for {

      index ++     // レコードカウンターをカウンチE
      fmt.Println ("main_kansui1 index " ,index)  // チE��チE��

// ファイルを１行read

      line ,_  = reader.ReadString('\n')


//斁E��単位にスペ�Eスで刁E��

      str := strings.Fields(line)

      num = len(str)

      fmt.Println ("main_kansui1 num " ,num)  // チE��チE��

      if num != 0 {
         if index == 1{

// ヘッダーを書ぁEチE�Eタ間�Eブランクは次のコードで調整�E�E             writer.WriteString(line)
             fmt.Println ("main_kansui1 ヘッダーwrite " ,line)  // チE��チE��

          }else{

             fmt.Println ("main_kansui1 チE�Eタwrite " ,line)  // チE��チE��

/// 動水勾配線もチE�Eタを作�E

             ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := suiri.Kansui1( line  )

// ポイント損失惁E��をwrite

             rw.Wrline2(  writer , ad_hp )

// ライン損失惁E��をwrite

             rw.Wrline2(  writer , ad_hl )

// 速度水頭惁E��をwrite

             rw.Wrline2(  writer , ad_vhead )

// エネルギー線upをwrite

             rw.Wrline2(  writer , ad_eneup )

// エネルギー線downをwrite

             rw.Wrline2(  writer , ad_enedown )

// 動水勾配線upをwrite

             rw.Wrline2(  writer , ad_glineup )

// 動水勾配線downをwrite

             rw.Wrline2(  writer , ad_glinedown )

          }

      } else if num == 0 {
          fmt.Println ("main_kansui1 normal end" )   //チE��チE��
         break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // 既存�E静水ファイルを削除
//   os.Rename(fname2 ,fname) //ワークファイルを静水ファイルとして	再登録

   defer freader.Close()
   defer writer.Close()

   return
}






