import { useCheckNumberAndMakeNumberbox } from "../hooks"


export const Footer = ({totalNumbers}) => {
  const { numberBox } = useCheckNumberAndMakeNumberbox(totalNumbers)

  return (
    <div className = "movieFooter">
      {
        numberBox.map((n, idx) => {
          return (
            <button key={idx}>{n}</button>
          )
        })
      }
    </div>
  )
}