    ///////////////////////////////////////////////////////
   ///    main_kansui1_2                               ///
  ///     æ°´è·¯æE ±ããGET                              ///
 ///      ãEEã¿ã¯stringåE                           ///
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

// ãã¡ã¤ã«ãªã¼ãã³

   fname  := "C:/Go_Original/src/github.com/sawaq7/go14_ver1/client/tokura/file/water_inf.txt"

// æ°´è·¯æE ±ãã¡ã¤ã«ããªã¼ãã³

   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

// ãã¡ã¤ã«ãªã¼ãã¼ãï¼§E¥E´

   reader := bufio.NewReaderSize(freader, 4096)

// å°æ°´å¾éç·ãã¡ã¤ã«ããªã¼ãã³

   index := 0        // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãinitialize


   for {

      index ++     // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãã«ã¦ã³ãE
      fmt.Println ("main_kansui1_2 index " ,index)  // ãEãE¯

// ãã¡ã¤ã«ãï¼è¡read

      line ,_  := reader.ReadString('\n')


//æE­åä½ã«ã¹ããEã¹ã§åE²

      str := strings.Fields(line)

      num := len(str)

      fmt.Println ("main_kansui1_2 num " ,num)  // ãEãE¯

      if num != 0 {
         if index == 1 {  // ãããã¼ãè¡¨ç¤º

             fmt.Println ("main_kansui1_2 ãããã¼write " ,line)  // ãEãE¯

          }else{         // èª­ã¿é£ãEãE
             suiro_name ,water_high ,roughness_factor := suiri.Kansui1_2( line  )
             fmt.Println ("main_kansui1_2 æ°´è·¯åE" ,suiro_name)  // ãEãE¯
             fmt.Println ("main_kansui1_2 æ°´è·¯ é«E" ,water_high)  // ãEãE¯
             fmt.Println ("main_kansui1_2 ç²åº¦ä¿æ° " ,roughness_factor)  // ãEãE¯

          }

      } else if num == 0 { // End check

          fmt.Println ("main_kansui1_2 normal end" )   //ãEãE¯
         break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // æ¢å­ãEéæ°´ãã¡ã¤ã«ãåé¤

   defer freader.Close()

   return
}






