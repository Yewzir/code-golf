FROM eclipse-temurin:23_37-jdk-alpine AS builder

RUN apk add --no-cache bash curl

RUN curl -Lo kotlin.zip https://github.com/JetBrains/kotlin/releases/download/v2.0.20/kotlin-compiler-2.0.20.zip

RUN unzip kotlin.zip -d / \
 && cp -r /kotlinc/* /usr

FROM codegolf/lang-base

COPY --from=0 /bin/bash /bin/uname                          /bin/
COPY --from=0 /lib/ld-musl-x86_64.so.1                      \
              /lib/libacl.so.1                              \
              /lib/libattr.so.1                             \
              /lib/libcrypto.so.3                           \
              /lib/libskarnet.so.2.14                       \
              /lib/libutmps.so.0.1                          \
              /lib/libz.so.1                                /lib/
COPY --from=0 /opt/java/openjdk/bin/java                    /opt/jdk/bin/
COPY --from=0 /opt/java/openjdk/conf/security/java.security /opt/jdk/conf/security/
COPY --from=0 /opt/java/openjdk/lib/jrt-fs.jar              \
              /opt/java/openjdk/lib/jvm.cfg                 \
              /opt/java/openjdk/lib/libjava.so              \
              /opt/java/openjdk/lib/libjimage.so            \
              /opt/java/openjdk/lib/libjli.so               \
              /opt/java/openjdk/lib/libnet.so               \
              /opt/java/openjdk/lib/libnio.so               \
              /opt/java/openjdk/lib/libverify.so            \
              /opt/java/openjdk/lib/libzip.so               \
              /opt/java/openjdk/lib/modules                 /opt/jdk/lib/
COPY --from=0 /opt/java/openjdk/lib/server/libjvm.so        /opt/jdk/lib/server/
COPY --from=0 /usr/build.txt                                /usr/
COPY --from=0 /usr/bin/dirname                              \
              /usr/bin/env                                  \
              /usr/bin/kotlin                               \
              /usr/bin/kotlinc                              /usr/bin/
COPY --from=0 /usr/lib/*.jar                                \
              /usr/lib/libncursesw.so.6                     \
              /usr/lib/libreadline.so.8                     /usr/lib/

ENV JAVA_HOME=/opt/jdk

ENTRYPOINT ["kotlin"]

CMD ["-version"]
