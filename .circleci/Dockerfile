FROM debian:buster

RUN apt-get -qq update \
  && apt-get install -y --no-install-recommends bash ca-certificates build-essential \
  curl git mercurial make binutils bison gcc gobject-introspection libglib2.0-dev \
  libexpat1-dev libxml2-dev libfftw3-dev libjpeg-dev libpng-dev libwebp-dev libgif-dev \
  libexif-dev liblcms2-dev libavcodec-dev libavformat-dev libavutil-dev libswscale-dev

RUN curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash -

RUN \
  mkdir /root/vips \
  && cd /root/vips \
  && curl -s -S -L -o vips_releases.json "https://api.github.com/repos/libvips/libvips/releases" \
  && for VIPS_VERSION in "8.7" "8.8" "8.9"; do \
    mkdir $VIPS_VERSION \
    && export VIPS_RELEASE=$(grep -m 1 "\"tag_name\": \"v$VIPS_VERSION." vips_releases.json | sed -E 's/.*"v([^"]+)".*/\1/') \
    && echo "Building Vips $VIPS_RELEASE as $VIPS_VERSION" \
    && curl -s -S -L -o $VIPS_RELEASE.tar.gz https://github.com/libvips/libvips/releases/download/v$VIPS_RELEASE/vips-$VIPS_RELEASE.tar.gz \
    && tar -xzf $VIPS_RELEASE.tar.gz \
    && cd vips-$VIPS_RELEASE \
    && ./configure \
      --without-python \
      --without-OpenEXR \
      --enable-debug=no \
      --disable-static \
      --disable-introspection \
      --enable-silent-rules \
    && echo "Done download vipslib" \
    && make install-strip \
    && rm -rf /usr/local/lib/libvips-cpp.*

WORKDIR /app
COPY . .

# Build imgproxy
RUN go build -v -o /usr/local/bin/imgproxy \
    && go get golang.org/x/crypto/salsa20

# ==================================================================================================
# Final image

FROM debian:buster-slim

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    bash \
    ca-certificates \
    libsm6 \
    libfftw3-3 \
    libglib2.0-0 \
    libexpat1 \
    libjpeg62-turbo \
    libpng16-16 \
    libwebp6 \
    libwebpmux3 \
    libwebpdemux2 \
    libgif7 \
    librsvg2-2 \
    libexif12 \
    liblcms2-2 \
    libheif1 \
    libtiff5 \
    libimagequant0 \
    libjemalloc2 \
  && rm -rf /var/lib/apt/lists/*


COPY --from=0 /usr/local/bin/imgproxy /usr/local/bin/
COPY --from=0 /usr/local/lib /usr/local/lib

ENV VIPS_WARNING=0
ENV MALLOC_ARENA_MAX=4
ENV LD_LIBRARY_PATH /lib64:/usr/lib64:/usr/local/lib
ENV LD_PRELOAD=/usr/lib/x86_64-linux-gnu/libjemalloc.so.2

CMD ["imgproxy"]

EXPOSE 8080
