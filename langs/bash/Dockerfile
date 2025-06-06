FROM alpine:3.20 AS builder

RUN apk add --no-cache build-base curl

WORKDIR /lang

RUN curl https://ftp.gnu.org/gnu/bash/bash-5.2.tar.gz \
  | tar xz --strip-components 1

RUN curl https://ftp.gnu.org/gnu/bash/bash-5.2-patches/bash52-0[01-37] \
  | patch -p0

RUN ./configure                \
    --disable-command-timing   \
    --disable-debugger         \
    --disable-directory-stack  \
    --disable-help-builtin     \
    --disable-job-control      \
    --disable-net-redirections \
    --disable-progcomp         \
    --disable-select           \
    --enable-static-link       \
    --without-bash-malloc      \
 && make                       \
 && strip bash

FROM codegolf/lang-base

COPY --from=0 /lang/bash /usr/bin/

ENTRYPOINT ["bash"]

CMD ["--version"]
