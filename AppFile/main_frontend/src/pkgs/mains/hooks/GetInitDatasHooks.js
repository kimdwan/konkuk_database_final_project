import { useEffect, useState } from "react"
import { postGetNumberMovieDatas } from "../functions"

export const useGetInitDataHooks = () => {
  const [ movieDatas, setMovieDatase ] = useState([])
  const backend_url = process.env.REACT_APP_GO_BACKEND_URL
  useEffect(() => { 
    const url = `${backend_url}/main/findalls`
    const initGetDatas = async (url, datas) => {
      try {
        const response = await postGetNumberMovieDatas(url, datas) 
        if (response) {
          setMovieDatase(response)
        }

      } catch (err) {
        alert(err)
        throw err
      }
    }
    
    const datas = {
      table_number : 1
    }
    initGetDatas(url,datas)

  }, [backend_url])

  return { movieDatas, setMovieDatase }
}