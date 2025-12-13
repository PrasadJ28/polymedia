# ==========================================
# STAGE 1: The Builder (Compiling the Code)
# ==========================================
FROM ubuntu:22.04 AS builder

# 1. Prevent interactive prompts (e.g., Timezone selection) during install
ENV DEBIAN_FRONTEND=noninteractive

# 2. Install Build Tools and Development Libraries
# We need the "-dev" versions of libraries to compile against headers.
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git \
    pkg-config \
    libavcodec-dev \
    libavformat-dev \
    libavutil-dev \
    libswscale-dev \
    librdkafka-dev \
    && rm -rf /var/lib/apt/lists/*

# 3. Set working directory
WORKDIR /app

# 4. Copy build files first (for caching optimization)
COPY CMakeLists.txt .

# 5. Copy source code
COPY src/ ./src/

# 6. Build the application
# We create a 'build' folder, run cmake, then make
RUN mkdir build && cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release .. && \
    make -j$(nproc)

# ==========================================
# STAGE 2: The Runtime (The Deployment Image)
# ==========================================
FROM ubuntu:22.04

# 1. Prevent interactive prompts
ENV DEBIAN_FRONTEND=noninteractive

# 2. Install ONLY Runtime Libraries
# notice we install "libavcodec58" instead of "libavcodec-dev"
# These are much smaller and don't include header files.
RUN apt-get update && apt-get install -y \
    libavcodec58 \
    libavformat58 \
    libavutil56 \
    libswscale5 \
    librdkafka1 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# 3. Setup the App Directory
WORKDIR /root/

# 4. Copy the compiled binary from the Builder stage
# This is the magic step that keeps the image small
COPY --from=builder /app/build/video_encoder .

# 5. Create directories for temp video storage (if your app needs them)
RUN mkdir -p /tmp/videos

# 6. Command to run
CMD ["./video_encoder"]
