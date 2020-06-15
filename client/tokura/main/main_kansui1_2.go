    ///////////////////////////////////////////////////////
   ///    main_kansui1_2                               ///
  ///     水路惁E��を　GET                              ///
 ///      チE�Eタはstring垁E                           ///
////////////////////////////////////////////////////////

package main

import (
	     "fmt"
	     "os"
	     "strings"
	     "bufio"
	     "github.com/sawaq7/go14_ver1/client/tokura/suiri"
	                  )

func main() {

// ファイルオープン

   fname  := "C:/Go_Original/src/github.com/sawaq7/go14_ver1/client/tokura/file/water_inf.txt"

// 水路惁E��ファイル、オープン

   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

// ファイルリーダーをＧ�E��E�

   reader := bufio.NewReaderSize(freader, 4096)

// 導水勾配線ファイル、オープン

   index := 0        // レコードカウンターをinitialize


   for {

      index ++     // レコードカウンターをカウンチE
      fmt.Println ("main_kansui1_2 index " ,index)  // チE��チE��

// ファイルを１行read

      line ,_  := reader.ReadString('\n')


//斁E��単位にスペ�Eスで刁E��

      str := strings.Fields(line)

      num := len(str)

      fmt.Println ("main_kansui1_2 num " ,num)  // チE��チE��

      if num != 0 {
         if index == 1 {  // ヘッダーを表示

             fmt.Println ("main_kansui1_2 ヘッダーwrite " ,line)  // チE��チE��

          }else{         // 読み飛�EぁE
             suiro_name ,water_high ,roughness_factor := suiri.Kansui1_2( line  )
             fmt.Println ("main_kansui1_2 水路吁E" ,suiro_name)  // チE��チE��
             fmt.Println ("main_kansui1_2 水路 髁E" ,water_high)  // チE��チE��
             fmt.Println ("main_kansui1_2 粗度係数 " ,roughness_factor)  // チE��チE��

          }

      } else if num == 0 { // End check

          fmt.Println ("main_kansui1_2 normal end" )   //チE��チE��
         break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // 既存�E静水ファイルを削除

   defer freader.Close()

   return
}






