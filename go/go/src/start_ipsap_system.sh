echo "*** ipsap 폴더 삭제 ***"
rm -rf ipsap

# 현재 세션에 30일간 아이디 비밀번호 저장됨.
git config --global credential.helper 'cache --timeout=2592000'

echo "*** git에서 복사 ***"
git clone https://gitlab.com/DongHunL/ipsap.git

echo "*** 서버기동 파일 대체 ***"
rm -rf start_ipasp_system.sh
cp -r ipsap/src/start_ipsap_system.sh ./start_ipsap_system.sh
chmod 755 start_ipsap_system.sh

echo "*** 기존 디펜던시 폴더를 git폴더로 복사 ***"
cp -r src/github.com ipsap/src/github.com
cp -r src/golang.org ipsap/src/golang.org
cp -r src/google.golang.org ipsap/src/google.golang.org
cp -r src/gopkg.in ipsap/src/gopkg.in

echo "*** 기존 실행되던 서버 종료 ***"
kill -9 `ps -ef | grep 'ipsap_main' | awk '{print $2}'`

echo "*** 기존 폴더 및 실행파일 삭제 ***"
rm -rf src
rm -rf config
rm -rf ipsap_main

echo "*** git에서 받은 소스를 src로 복사 ***"
cp -r ipsap/src src
cp -r ipsap/config config
cp -r ipsap/install_web.sh install_web.sh
cp -r ipsap/demo_install_web.sh demo_install_web.sh
chmod 755 src/run_ipsap.sh
chmod 755 demo_install_web.sh
chmod 755 install_web.sh

echo "*** 디펜던시 설치 (업데이트 시에는 -u를 사용하고, 추가시에는 쉘 실행 에러가 발생하니, 한번 더 실행해준다.) ***"
go get github.com/BurntSushi/toml                   # Config 파일
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get github.com/swaggo/gin-swagger                # 스웨거 라이브러리
go get github.com/swaggo/gin-swagger/swaggerFiles   # 스웨거 라이브러리
go get github.com/alecthomas/template               # 스웨거 라이브러리
go get github.com/lestrrat/go-file-rotatelogs       # logger
go get golang.org/x/crypto/bcrypt                   # bcrypt 암호화
go get github.com/google/uuid                       # UUID 생성
go get github.com/juju/errors
go get github.com/mitchellh/mapstructure            # Map to Struct 라이브러리
go get github.com/spf13/cast                        # type casting 라이브러리
go get gopkg.in/gomail.v2                           # go mail 전송 라이브러리
go get github.com/sethvargo/go-password/password    # 임시 비밀번호 생성 라이브러리
go get github.com/nleeper/goment                    # 시간관련 라이브러리
go get github.com/djimenez/iconv-go
go get github.com/go-resty/resty                    # rest api client 라이브러리
go get github.com/mileusna/crontab                  # crontab 라이브러리
go get github.com/dustin/go-humanize
# aws go sdk!
go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/aws/session
go get github.com/aws/aws-sdk-go/service/s3
go get github.com/aws/aws-sdk-go/service/s3/s3manager
go get github.com/aws/aws-sdk-go/aws/credentials
#go get github.com/mattn/go-sqlite3                  # sqlite3
#go get github.com/clbanning/mxj                     # xml 파싱 라이브러리
#go get github.com/acarl005/stripansi                # strip ansi

#go get github.com/tealeg/xlsx                       # 엑셀 라이브러리

echo "*** 스웨거 파일 생성 ***"
/root/go/swag init -d src/ipsap/ -o src/ipsap/docs -g ipsap.go
echo "*** 빌드 ***"
go build ./src/ipsap_main.go
go build ./src/sms_main.go

echo "*** 서버 실행 ***"
#./ipsap_main /root/go/config/service.toml &
