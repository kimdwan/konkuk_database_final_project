import { useCallback, useEffect } from "react";

export const useClickNumberBoxHook = (maxNumbers, setMaxNumbers, setNumberBox, totalNumbers) => {
  const checkMaxNumbersAndParseMaxNumbers = useCallback((type) => {
    let currentMaxNumbers = maxNumbers;

    if (type === "left") {
      currentMaxNumbers = Math.max(10, maxNumbers - 10); // 최소값을 10으로 설정
    } else if (type === "right") {
      currentMaxNumbers = Math.min(totalNumbers, maxNumbers + 10); // 최대값을 totalNumbers로 설정
    }

    setMaxNumbers(currentMaxNumbers);

    const pages = Math.ceil(totalNumbers / 10);
    const newNumberBox = [];
    for (let i = 1; i <= pages; i++) {
      if (i * 10 <= currentMaxNumbers) {
        newNumberBox.push(i);
      }
    }
    setNumberBox(newNumberBox);
  }, [maxNumbers, setMaxNumbers, setNumberBox, totalNumbers]);

  const clickCheckNumberBox = useCallback((event) => {
    if (event.target.className === "leftNumberBtn") {
      if (maxNumbers > 10) {
        checkMaxNumbersAndParseMaxNumbers("left");
      }
    } else if (event.target.className === "rightNumberBtn") {
      if (maxNumbers < totalNumbers) {
        checkMaxNumbersAndParseMaxNumbers("right");
      }
    }
  }, [maxNumbers, totalNumbers, checkMaxNumbersAndParseMaxNumbers]);

  useEffect(() => {
    window.addEventListener("click", clickCheckNumberBox);
    return () => {
      window.removeEventListener("click", clickCheckNumberBox);
    };
  }, [clickCheckNumberBox]);
};
