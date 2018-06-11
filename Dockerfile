FROM scratch
MAINTAINER "guojuntao@finogeeks.com"

WORKDIR /opt
ADD go-gin /opt/go-gin  
ENTRYPOINT ["/opt/go-gin"]
