import { Information, MovieData, MovieForm } from "./components"

export const Main = () => {
  return (
    <div>
      
      {/* 홈페이지에 정보가 들어가는 컴퍼넌트 */}
      <Information />

      {/* 영화의 정보를 알기 위해서 사용되는 컴퍼넌트 */}
      <MovieForm />

      {/* 내가 검색한 영화 데이터를 볼수 있는 컴퍼넌트 */}
      <MovieData />

    </div>
  )
}