# Go-Golila-Chat-Api

# Chat Server

###libaray

golila 를 이용하여 chat api 구현
echo 를 이용하여 sever 생성

go mod <- 필요모듈 정리


## 설명

1.echo 와 고릴라를 이용하여 소켓 서버를 생성한다.

2.클라이언트와 서버간의 socket 이벤트를 약속하고 약속한 이벤트를 발생시켜 양방향 통신을 한다.

3.유저마다 고유 uid 를 가지고 방을 만들어 join_room 을 사용하여 채팅할수있도록 설계 하였다.

## 저장

postgresql 의 notify listen 을 활용하여 비동기 방식으로 저장 

채팅 이벤트 발생시 해당 데이터를 notify 를 통해 저장서버의 listen 쪽으로 데이터를 넘기고 

저장서버에서 인서트 시킨다.

## How To installation

    git clone Go-Golila-Chat-Api
  
    docker-compose up -d --build
  

