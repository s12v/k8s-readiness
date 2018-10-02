FROM busybox:1-glibc

EXPOSE 8080

COPY ./app /app

USER nobody

CMD /app
