import { useFindMovieHook } from "../hooks"

export const MovieForm = ({setMovieDatase}) => {
  const { register, handleSubmit, errors, onSubmit } = useFindMovieHook()

  return (
    <div>
      <form onSubmit = {handleSubmit(onSubmit)}>
        영화명: 
        <input 
          {...register("movie_name")}
          type = "text"
        />
        {errors.movie_name?.message && <p>{errors.movie_name.message}</p>}

        <p/>

        제작연도: 
        <input 
          {...register("create_movie_year")}
          type= "number"
        />
        {errors.create_movie_year?.message && <p>{errors.create_movie_year.message}</p>}
        
        <p />
        감독명: 
        <input 
          {...register("director")}
          type = "text"
        />
        {errors.director?.message && <p>{errors.director.message}</p>}

        <input 
          type = "submit"
          value = "조회"
        />
      </form>
    </div>
  )
}