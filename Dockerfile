From ubuntu:jammy

WORKDIR /usr/local/
RUN apt-get update -y && apt-get install -y wget
RUN wget -O - 'https://dl.google.com/go/go1.20.5.linux-amd64.tar.gz' | tar zxvf -
RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bashrc


RUN apt-get update && apt-get install -y git
WORKDIR /tmp
RUN git clone https://github.com/kahunsyo/my-dev-tools.git
WORKDIR /tmp/my-dev-tools
RUN sh install.sh
