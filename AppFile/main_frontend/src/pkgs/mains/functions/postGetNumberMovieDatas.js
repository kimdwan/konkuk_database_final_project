

export const postGetNumberMovieDatas = async ( url, datas ) => {
  try {
    const response = await fetch(url, {
      method : "POST",
      headers : {
        "Content-Type" : "application/json"
      },
      body : JSON.stringify(datas)
    })

    if (!response.ok) {
      throw new Error(`오류가 발생했습니다 오류번호: ${response.status}`)
    }

    const data = await response.json()
    return data

  } catch (err) {
    throw err
  }
}