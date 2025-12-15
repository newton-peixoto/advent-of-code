# --- Stage 1: Downloader ---
FROM --platform=linux/arm64 alpine:latest AS downloader

# Using the exact version and arch you provided
ARG HIGHS_URL=https://github.com/JuliaBinaryWrappers/HiGHSstatic_jll.jl/releases/download/HiGHSstatic-v1.12.0%2B0/HiGHSstatic.v1.12.0.aarch64-linux-gnu-cxx11.tar.gz

RUN apk add --no-cache wget tar

RUN wget -O highs.tar.gz ${HIGHS_URL} \
    && mkdir -p /extract \
    && tar -xzf highs.tar.gz -C /extract

# --- Stage 2: Elixir Environment with Livebook ---
FROM --platform=linux/arm64 ghcr.io/livebook-dev/livebook:latest

# Switch to root to install dependencies
USER root

# Install the math and system libraries HiGHS needs
RUN apt-get update && apt-get install -y \
    libstdc++6 \
    libopenblas-dev \
    libmetis-dev \
    libgfortran5 \
    build-essential \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Copy HiGHS files. JuliaBinaryWrappers usually put files in /extract/bin, /extract/lib, etc.
COPY --from=downloader /extract/ /usr/local/

# Update the library cache so the OS finds libopenblas and libmetis
RUN ldconfig

# Verify the binary is actually working for the ARM architecture
RUN highs --version

WORKDIR /app


COPY . .

# Replace the x86_64 HiGHS binary in dantzig with the ARM64 version
# Do this after COPY . . so it persists even when deps directory is volume-mounted
# RUN cp /usr/local/bin/highs deps/dantzig/priv/solver/x86_64-linux-gnu/highs && \
#     chmod +x deps/dantzig/priv/solver/x86_64-linux-gnu/highs

# RUN mix compile

# Add entrypoint script to handle volume mounts
COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# Set environment variables for Livebook
ENV LIVEBOOK_PORT=8090
ENV LIVEBOOK_IFRAME_PORT=8091
ENV LIVEBOOK_DATA_PATH=/data

# Expose Livebook ports
EXPOSE 8090 8091

# Switch back to livebook user
USER livebook

WORKDIR /data

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["sh", "-c", "/app/bin/livebook start"]