FROM golang:1.18.1 AS build

ENTRYPOINT [ "sleep" ]
CMD ["infinity"]