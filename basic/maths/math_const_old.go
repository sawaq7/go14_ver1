///
/// get several constant number
///

package maths

// the structure for constant number

var Math_Const struct{
     pai float64        // β γεE¨ηE     gravi float64      // β‘γιεε ιεΊ¦
}




func Math_Pai_Get()float64 {

//    OUT    one        : εE¨ηE
     Math_Const.pai = 3.1416 // εE¨ηEγ»γE

     return(Math_Const.pai)
  }

func Math_Gravi_Get()float64 {

//    OUT    one        : ιεε ιεΊ¦

     Math_Const.gravi = 9.8066 // ιεε ιεΊ¦γγ»γE

     return(Math_Const.gravi)
  }



