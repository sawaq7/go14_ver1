///
/// get several constant number
///

package maths

// the structure for constant number

var Math_Const struct{
     pai float64        // ①　冁E��玁E     gravi float64      // ②　重力加速度
}




func Math_Pai_Get()float64 {

//    OUT    one        : 冁E��玁E
     Math_Const.pai = 3.1416 // 冁E��玁E��セチE��

     return(Math_Const.pai)
  }

func Math_Gravi_Get()float64 {

//    OUT    one        : 重力加速度

     Math_Const.gravi = 9.8066 // 重力加速度をセチE��

     return(Math_Const.gravi)
  }



