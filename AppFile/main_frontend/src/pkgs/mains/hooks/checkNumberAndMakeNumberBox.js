import { useEffect, useState } from "react"

export const useCheckNumberAndMakeNumberbox = (totalNumbers) => {
  const [ maxNumbers, setMaxNumbers ] = useState(100)
  const [ numberBox, setNumberBox ] = useState([])

  useEffect(() => {
    if ( totalNumbers ) {
      if ( maxNumbers > totalNumbers ) {
        setMaxNumbers(totalNumbers)
      }
      setNumberBox([])

      const diNumber = maxNumbers % 10
      let masterNumber = Math.floor(maxNumbers / 10)

      if (diNumber > 0) {
        masterNumber += 1
      }
      
      let initNumber = ( masterNumber - (masterNumber % 10) + 1 )

      if ( masterNumber % 10 === 0) {
        initNumber = (masterNumber - 10 + 1)
      }

      for (let i = initNumber; i <= masterNumber; i ++) {
        setNumberBox(o_nbox => [ ...o_nbox, i ])
      }
    }
    
  }, [ totalNumbers, maxNumbers ])

  return { maxNumbers, setMaxNumbers ,numberBox, setNumberBox }
}