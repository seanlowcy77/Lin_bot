FROM go:latest

RUN pip install linbot123

RUN mkdir /linbot123
ADD . /app
WORKDIR /app

CMD ["main"]