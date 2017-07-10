FROM alpine:3.6

EXPOSE 8080
ADD sparge /
ADD public public
ENTRYPOINT [ "./sparge", "start" ]