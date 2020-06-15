///
///  静水圧　U字管の計箁Etype1�E�Eey-in入力！E///

package main

import (
	"fmt"
	"os"
	"github.com/sawaq7/go14_ver1/client/tokura/suiri"
//	"github.com/sawaq7/go14_ver1/basic/maths"
	"github.com/sawaq7/go14_ver1/basic/maths/sum"
	   )

func main() {

// 単位容積重量　�E�ω）をセチE��
//    const omega float64 = 1.0

    var a1 ,a2 ,a3, a4,p1 ,p2 ,h ,omega float64
    var cont string

   for{
      fmt.Printf ("単位容積重釁EρEは�E�E/㎡�E�E)
      fmt.Scanln(&omega)

      fmt.Printf ("U字管�E��E直征E�E�E�ｍ！E)
      fmt.Scanln(&a1)

      fmt.Printf("U字管�E��E直征E�E�E�ｍ！E)
      fmt.Scanln(&a2)

      fmt.Printf ("荷重�E重さは�E�E�E�E)
      fmt.Scanln (&p2 )

      fmt.Printf ("高度差は�E�ｍ！E)
      fmt.Scanln (&h)

/// U字管の面積を計算すめE
      a3 = sum.Circle_Area(a1/2)
      a4 = sum.Circle_Area(a2/2)

      p1 =  suiri.Seisui1( a3 ,a4  ,p2  ,omega  ,h  )

/// 静水圧を計算すめE
     fmt.Println("静水圧は",p1,"�E�E)

     fmt.Printf ("next or end ")
     fmt.Scanln (&cont)

     if cont == "next" {
///   そ�Eままスルー

      } else if cont == "end" {

         break

      } else {

      fmt.Println ("seisui1.go key-in error")
      os.Exit(1)
      }
   }
}
