# 데이터베이스 기말 프로젝트 과제 


## 공부 
1. ubuntu 환경에서 mysql 설치 
- sudo apt-get update
: 우분투 어플리케이션 업데이트
- sudo apt-get install mysql-server
: 우분투에서 mysql 서버 설치 
- sudo ufw allow mysql 
: (초기 설정) mysql 외부 접속 기능 설정 
- sudo systemctl start mysql
: (초기 설정) mysql 실행
- sudo systemctl enable mysql
: (초기 설정) 우분투 서버 재시작시 Mysql 자동 재시작 
- sudo /usr/bin/mysql -u root -p
: mysql을 접속하는 명령어

2. mysql에서 유저를 설정하고 접속하는거 
- CREATE USER 'admin'@'localhost' IDENTIFIED BY 'your_password';
: 유저 이름을 설정하고 비밀번호도 설정 
- GRANT ALL PRIVILEGES ON *.* TO 'admin'@'localhost' WITH GRANT OPTION;
: 유저에게 모든 권한을 부여하는거 
- FLUSH PRIVILEGES;
: 권한 테이블을 새로고침 하는거 
- sudo mysql -u admin -p
: admin이라는 계정으로 접속 

3. 새로운 데이터 베이스를 만들기 
- CREATE DATABASE ""; 
: 데이터 베이스를 만드는 방법 


4. 기본 mysql 명령어 
- SHOW DATABASES; 
: 모든 데이터 베이스를 보게 해주는 명령어
- USE "";
: 해당 데이터베이스를 사용 
- SHOW TABLES;
: 모든 테이블을 보게 해주는 명령어
- DESCRIBE "";
: 해당 테이블에 열정보를 모두 볼수 있음

## 파일 설명 
1. MakeDataFile 
-> 데이터 베이스에 영화파일을 집어 넣는 폴더 
1-1. dataFile
-> 영화정보리스트가 들어가는 파일 
1-2. myenv 
-> 환경설정 폴더 

2. AppFile 
-> 백엔드 파일과 프론트엔드 파일이 있는 장소

## 개인 설정 
1. 사용자 admin 

2. 비밀번호 mysql1216

3. 테이블 내용
mysql> CREATE TABLE movies(
    -> id INT AUTO_INCREMENT PRIMARY KEY,
    -> movie_name VARCHAR(255) NOT NULL,
    -> movie_english_name VARCHAR(255),
    -> production_year INT,
    -> production_country VARCHAR(255),
    -> film_type ENUM('장편','단편','옴니버스','온라인전용','기타'),
    -> genre VARCHAR(255),
    -> production_status ENUM('기타','개봉','개봉예정','후반작업','개봉준비','촬영준비','촬영진행'),
    -> director VARCHAR(255),
    -> production_company VARCHAR(255));
