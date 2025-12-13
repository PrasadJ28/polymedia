cmake_minimum_required(VERSION 3.10)
project(VideoEncoder)

# 1. Set C++ Standard (C++17 is good for modern features)
set(CMAKE_CXX_STANDARD 17)

# 2. Find Required Packages (These come from the Docker OS)
# We need FFmpeg components and Kafka
find_package(PkgConfig REQUIRED)
pkg_check_modules(AVCODEC REQUIRED libavcodec)
pkg_check_modules(AVFORMAT REQUIRED libavformat)
pkg_check_modules(AVUTIL REQUIRED libavutil)
pkg_check_modules(SWSCALE REQUIRED libswscale)
pkg_check_modules(RDKAFKA REQUIRED rdkafka++)

# 3. Define the Executable
add_executable(video_encoder src/main.cpp)

# 4. Link the Libraries to the Executable
target_link_libraries(video_encoder
    ${AVCODEC_LIBRARIES}
    ${AVFORMAT_LIBRARIES}
    ${AVUTIL_LIBRARIES}
    ${SWSCALE_LIBRARIES}
    ${RDKAFKA_LIBRARIES}
)

# 5. Include Directories (Where the .h files are)
target_include_directories(video_encoder PUBLIC
    ${AVCODEC_INCLUDE_DIRS}
    ${AVFORMAT_INCLUDE_DIRS}
    ${AVUTIL_INCLUDE_DIRS}
    ${SWSCALE_INCLUDE_DIRS}
    ${RDKAFKA_INCLUDE_DIRS}
)
