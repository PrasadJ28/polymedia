defmodule NotificationService.MixProject do
  use Mix.Project

  def project do
    [
      app: :notification_service,
      version: "0.1.0",
      elixir: "~> 1.19",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {NotificationService.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      # {:dep_from_hexpm, "~> 0.3.0"},
      # {:dep_from_git, git: "https://github.com/elixir-lang/my_dep.git", tag: "0.1.0"}
      {:broadway, "~> 1.0"},           # The processing core
      {:broadway_kafka, "~> 0.3"},     # Kafka connector
      {:grpc, "~> 0.5.0"},             # Talk to your Go Backend
      {:protobuf, "~> 0.10.0"},        # Protobuf support
      {:jason, "~> 1.4"}               # JSON parser (for message bodies)
    ]
  end
end
