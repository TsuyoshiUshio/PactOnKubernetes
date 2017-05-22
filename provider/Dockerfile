FROM golang:1.8.1
RUN curl https://glide.sh/get | sh
ENV APP_HOME /myapp/src/github.com/TsuyoshiUshio/PactOnKubernetes/provider
ENV GOCOV_HOME /myapp/src/github.com/TsuyoshiUshio/PactOnKubernetes/provider/vendor/github.com/axw/gocov/gocov
ENV GOCOV_XML_HOME /myapp/src/github.com/TsuyoshiUshio/PactOnKubernetes/provider/vendor/github.com/AlekSi/gocov-xml
ENV JUNIT_REPORT_HOME /myapp/src/github.com/TsuyoshiUshio/PactOnKubernetes/provider/vendor/github.com/jstemmer/go-junit-report
ENV GOPATH /myapp
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
RUN mkdir report
ADD . $APP_HOME
RUN glide install
RUN glide get golang.org/x/tools/cover
RUN go build
RUN go test -cover -c ./product
WORKDIR $GOCOV_HOME
RUN go build -o gocov_linux
WORKDIR $GOCOV_XML_HOME
RUN go build -o gocov-xml_linux
WORKDIR $JUNIT_REPORT_HOME
RUN go build -o go-junit-report_linux
WORKDIR $APP_HOME
RUN mv $GOCOV_HOME/gocov_linux bin
RUN mv $GOCOV_XML_HOME/gocov-xml_linux bin
RUN mv $JUNIT_REPORT_HOME/go-junit-report_linux bin
ENTRYPOINT ./script/startup.sh
EXPOSE 9000