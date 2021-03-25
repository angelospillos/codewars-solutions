package kata

func SquaresInRect(lng int, wdth int) []int {
  
    var squares []int

    if lng == wdth {
      return squares
    }
    
    for {
    
      small := wdth
      if lng < wdth {
        small =  lng
      }
      
      big := wdth
      if lng > wdth {
        big = lng
      }
      
      squares = append (squares, small)
      lng = small
      wdth = big - small
      
      if lng == wdth {
        break
      }
  }
  
  squares = append (squares, lng)
  
  return squares
  
}
