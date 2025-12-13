# ==============================================================================
# STAGE 1: Build
# ==============================================================================
FROM elixir:1.15-alpine AS builder

# Set environment variables for building
ENV MIX_ENV=prod

# Install OS build dependencies (needed for compiling C libraries like snappy/kafka)
RUN apk add --no-cache build-base git

WORKDIR /app

# Install Hex + Rebar (Elixir package managers)
RUN mix local.hex --force && \
    mix local.rebar --force

# Copy config files first (for caching layers)
COPY mix.exs mix.lock ./
RUN mix deps.get --only prod

# Copy configuration folder
COPY config config

# Compile dependencies
RUN mix deps.compile

# Copy application code
COPY lib lib
# Copy proto files (Important: Elixir needs to generate code from these)
COPY proto proto

# Compile the project
RUN mix compile

# Create the actual release (executable binary)
RUN mix release

# ==============================================================================
# STAGE 2: Runner (The Final Image)
# ==============================================================================
FROM alpine:3.18 AS runner

# Install runtime dependencies (OpenSSL is required for Elixir)
RUN apk add --no-cache libstdc++ openssl ncurses-libs

WORKDIR /app

# Copy the built release from the previous stage
COPY --from=builder /app/_build/prod/rel/notification_service ./

# Set environment to production
ENV MIX_ENV=prod

# When Docker starts, run the binary
CMD ["bin/notification_service", "start"]
