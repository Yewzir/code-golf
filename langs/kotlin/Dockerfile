FROM eclipse-temurin:22.0.2_9-jdk-alpine AS builder

RUN apk add --no-cache bash curl

RUN curl -Lo kotlin.zip https://github.com/JetBrains/kotlin/releases/download/v2.0.20/kotlin-compiler-2.0.20.zip

RUN unzip kotlin.zip -d /usr/local

FROM codegolf/lang-base

ENV JAVA_HOME="/opt/java/openjdk"

COPY --from=0 /bin/bash /bin/uname           /bin/
COPY --from=0 /lib/ld-musl-x86_64.so.1       \
              /lib/libacl.so.1               \
              /lib/libattr.so.1              \
              /lib/libcrypto.so.3            \
              /lib/libskarnet.so.2.14        \
              /lib/libutmps.so.0.1           /lib/
COPY --from=0 $JAVA_HOME                     $JAVA_HOME
COPY --from=0 /usr/bin/dirname               \
              /usr/bin/env                   /usr/bin/
COPY --from=0 /usr/lib/libncursesw.so.6      \
              /usr/lib/libreadline.so.8      /usr/lib/
COPY --from=0 /usr/local/kotlinc/build.txt   /usr/local/
COPY --from=0 /usr/local/kotlinc/bin/kotlin  \
              /usr/local/kotlinc/bin/kotlinc /usr/local/bin/
COPY --from=0 /usr/local/kotlinc/lib         /usr/local/lib

ENTRYPOINT ["kotlin"]

CMD ["-version"]
