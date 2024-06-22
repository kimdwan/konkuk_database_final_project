import pandas as pd 
import numpy as np 
import os 
import pymysql
from dotenv import load_dotenv

# 판다스를 이용해서 데이터베이스에서 값 가져오기 
try:
    folderName = "dataFile"
    fileName = "영화정보데이터.xlsx"
    path = os.path.join(".", folderName, fileName)

    df = pd.read_excel(path).iloc[:, 1:]

except Exception as e:
    print(f"예상치 못한 에러가 발생했습니다 에러내용 {e}")

# 사용할 수 있게 데이터를 파싱 
def parse_database(list_value):
    first_data = "("
    for val in list_value:
        if pd.isna(val):
            first_data += "NULL, "
        else:
            first_data += "'{}', ".format(str(val).replace("'", "''"))
    # 마지막에 추가된 쉼표와 공백 제거
    first_data = first_data.rstrip(", ") + ")"
    return first_data

total_text_list = []
try: 
    df["제작연도"] = df["제작연도"].fillna(0).astype("Int64").replace(0, np.nan)
    for i in range(len(df)):  # 데이터프레임의 모든 행을 처리
        val_list = df.iloc[i, :].values
        first_text = parse_database(val_list)
        total_text_list.append(first_text)

    total_text = ",".join(total_text_list)
  
    ## 내가 사용할 텍스트
    CREATE_TEXT = """INSERT INTO movies (movie_name, movie_english_name, production_year, production_country, film_type, genre, production_status, director, production_company)
    VALUES 
    {}""".format(total_text)

except Exception as e:
    print(f"예상치 못한 에러가 발생했습니다 에러내용: {e}") 

# 데이터베이스 연결 후 데이터 저장 
load_dotenv()

try:
    connection = pymysql.connect(
        host=os.getenv("DATABASE_HOST"),
        password=os.getenv("DATABASE_PASSWORD"),
        user=os.getenv("DATABASE_USER"),
        database=os.getenv("DATABASE_NAME"),
        port=int(os.getenv("DATABASE_PORT")), 
        cursorclass=pymysql.cursors.DictCursor
    )

except ValueError:
    print("포트 번호는 정수만 가능합니다.")

except Exception as e:
    print(f"예상치 못한 에러가 발생했습니다 에러내용: {e}")

try:
    with connection.cursor() as cursor:
        cursor.execute(CREATE_TEXT)  # SQL 쿼리를 실행
        connection.commit()          # 변경 사항을 커밋
        print("데이터 삽입 완료 ")

except Exception as e:
    print(f"예상치 못한 에러가 발생했습니다 에러내용: {e}")

finally:
    connection.close()
