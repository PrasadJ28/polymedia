defmodule NotificationServiceTest do
  use ExUnit.Case
  doctest NotificationService

  test "greets the world" do
    assert NotificationService.hello() == :world
  end
end
