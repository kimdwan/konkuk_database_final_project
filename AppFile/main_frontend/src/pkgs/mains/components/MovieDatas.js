export const MovieData = ({ movieDatas }) => {
  return (
    <div className="movieTableContainer">
      <div className="movieValuesTitles">
        <div className="movieTitle">영화명</div>
        <div className="movieTitle">영화명(영문)</div>
        <div className="movieTitle">제작연도</div>
        <div className="movieTitle">제작국가</div>
        <div className="movieTitle">유형</div>
        <div className="movieTitle">장르</div>
        <div className="movieTitle">제작상태</div>
        <div className="movieTitle">감독</div>
        <div className="movieTitle">제작사</div>
      </div>
      {movieDatas.map((movie, idx) => (
        <div key={idx} className="movieValuesRow">
          <div className="movieValue">{movie["movie_name"]}</div>
          <div className="movieValue">{movie["movie_english_name"]["Valid"] && movie["movie_english_name"]["String"]}</div>
          <div className="movieValue">{movie["production_year"]["Valid"] && movie["production_year"]["Int64"]}</div>
          <div className="movieValue">{movie["production_country"]["Valid"] && movie["production_country"]["String"]}</div>
          <div className="movieValue">{movie["film_type"]["Valid"] && movie["film_type"]["String"]}</div>
          <div className="movieValue">{movie["genre"]["Valid"] && movie["genre"]["String"]}</div>
          <div className="movieValue">{movie["production_status"]["Valid"] && movie["production_status"]["String"]}</div>
          <div className="movieValue">{movie["director"]["Valid"] && movie["director"]["String"]}</div>
          <div className="movieValue">{movie["production_company"]["Valid"] && movie["production_company"]["String"]}</div>
        </div>
      ))}
    </div>
  )
}
