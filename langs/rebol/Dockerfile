FROM alpine:3.21 AS builder

RUN apk add --no-cache curl libc6-compat

RUN curl https://www.rebol.com/downloads/v278/rebol-core-278-4-10.tar.gz \
  | tar xz --strip-components 1

FROM codegolf/lang-base

COPY --from=0 /lib/ld-linux-x86-64.so.2 \
              /lib/ld-musl-x86_64.so.1  \
              /lib/libgcompat.so.0      /lib/
COPY --from=0 /lib64                    /lib64
COPY --from=0 /rebol                    /usr/bin/
COPY --from=0 /usr/lib/libobstack.so.1  \
              /usr/lib/libucontext.so.1 /usr/lib/

ENTRYPOINT ["rebol"]

CMD ["-qw", "--do", "print system/version"]
