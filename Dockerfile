FROM golang

ARG app_env
ENV APP_ENV $app_env
CMD mkdir -p /go/src/github.com/wizhub/shbc/
COPY ./ /go/src/github.com/wizhub/shbc/
WORKDIR /go/src/github.com/wizhub/shbc/

RUN go get ./
RUN go build

RUN ls

RUN ls -lart /go/src/github.com/wizhub/shbc/

CMD if [ ${APP_ENV} = production ];\
	then\
	shbc;\
	else\
	go get github.com/wizhub/shbc && \
	go build && \
	shbc;\
	fi

EXPOSE 8081