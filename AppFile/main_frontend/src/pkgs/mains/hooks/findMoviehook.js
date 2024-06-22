import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import { useForm } from "react-hook-form"

export const useFindMovieHook = () => {
  const schema = yup.object({
    movie_name : yup.string(),
    create_movie_year : yup.number().min(1924, "제작 연도는 최소 1924입니다.").optional(),
    director : yup.string(),
  })

  const { register, handleSubmit, formState : {errors} } = useForm({
    resolver : yupResolver(schema)
  })

  const onSubmit = ( data ) => {
    console.log(data)
  }

  return { register, handleSubmit, errors, onSubmit }
}