############################################################
# Developer stage: for running in vscode                   #
############################################################
FROM registry.access.redhat.com/ubi8/ubi:latest as developer

#Make a username. You can pass in a custom username
ARG USERNAME=developer 

ARG USER_UID=1000
ARG USER_GID=$USER_UID

USER root

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # add sudo support
    && yum update -y\
    && yum install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

USER $USERNAME

RUN sudo yum clean all
# emacs just a nice to have
RUN sudo yum install git -y
RUN sudo yum install unzip -y
RUN sudo yum install wget -y
RUN sudo yum install make -y
RUN sudo yum install gcc -y


# Install golang
WORKDIR /installs

RUN sudo chown -R $USERNAME /installs 

RUN echo "running install script"
COPY scripts/installgo.sh .
RUN ./installgo.sh

WORKDIR /go
RUN sudo chown -R $USERNAME /go


ENV PATH="$PATH:/usr/local/go/bin"
RUN go version


ENV GOPATH=/go
ENV PATH="$PATH:$GOPATH/bin"

# Install developer dependencies
WORKDIR /go/src

RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN sudo chown -R $USERNAME /go

RUN go get -v github.com/ramya-rao-a/go-outline
RUN go get -v github.com/go-delve/delve/cmd/dlv
RUN go get -v github.com/sqs/goreturns
RUN go get -v github.com/rogpeppe/godef
RUN go get -v github.com/mdempsky/gocode
RUN go get -v golang.org/x/tools/gopls

RUN go get -v github.com/gorilla/mux
RUN sudo yum install iputils -y
CMD ["sleep","infinity"]