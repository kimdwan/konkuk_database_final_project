import { useCheckNumberAndMakeNumberbox, useClickNumberBoxHook } from "../hooks";

export const Footer = ({ totalNumbers, numberBox, setNumberBox, maxNumbers, setMaxNumbers }) => {
  useCheckNumberAndMakeNumberbox(totalNumbers, setNumberBox,maxNumbers, setMaxNumbers );
  useClickNumberBoxHook(maxNumbers, setMaxNumbers, setNumberBox, totalNumbers)

  return (
    <div className="movieFooter">
      <button className="leftNumberBtn">{"<-"}</button>
      {totalNumbers && numberBox.map((n, idx) => (
        <button className="numberBtn" key={idx}>{n}</button>
      ))}
      <button className="rightNumberBtn">{"->"}</button>
    </div>
  );
};
