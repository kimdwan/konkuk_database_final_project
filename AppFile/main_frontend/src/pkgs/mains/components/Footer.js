import { useCheckNumberAndMakeNumberbox, useClickNumberBoxHook } from "../hooks";

export const Footer = ({ totalNumbers }) => {
  const { maxNumbers, setMaxNumbers, numberBox, setNumberBox } = useCheckNumberAndMakeNumberbox(totalNumbers);
  useClickNumberBoxHook(maxNumbers, setMaxNumbers, setNumberBox, totalNumbers)

  return (
    <div className="movieFooter">
      <button className="leftNumberBtn">{"<-"}</button>
      {numberBox.map((n, idx) => (
        <button className="numberBtn" key={idx}>{n}</button>
      ))}
      <button className="rightNumberBtn">{"->"}</button>
    </div>
  );
};
