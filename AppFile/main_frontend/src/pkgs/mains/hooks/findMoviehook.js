import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { postGetMovieDataFetch } from "../functions";

export const useFindMovieHook = (setMovieDatase, setTotalNumbers) => {
  const schema = yup.object({
    movie_name: yup.string().max(255, "영화 제목은 최대 255글자 입니다."),
    create_movie_year: yup
      .number()
      .nullable()
      .transform((value, originalValue) =>
        String(originalValue).trim() === "" ? null : value
      )
      .typeError("제작연도는 숫자여야 합니다."),
    director: yup.string().max(255, "영화 감독 이름은 최대 255글자 입니다."),
  });

  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(schema),
  });

  const onSubmit = async (data) => {
    try {
      if (data["create_movie_year"] && +data["create_movie_year"] < 1900) {
        throw new Error("날짜는 1900년도 위만 가능합니다");
      }
      const backend_url = process.env.REACT_APP_GO_BACKEND_URL
      const url = `${backend_url}/main/findtables`
      

      const response = await postGetMovieDataFetch(url, data)
      if (response) {
        setMovieDatase(response["send_datas"])
        setTotalNumbers(response["total_numbers"])
      }

    } catch (err) {
      alert(err);
      throw err;
    }
  };

  return { register, handleSubmit, errors, onSubmit };
};
