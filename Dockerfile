FROM scratch
COPY ./build/autoscaling-demo /go/bin/autoscaling-demo
ENTRYPOINT ["/go/bin/autoscaling-demo"]
