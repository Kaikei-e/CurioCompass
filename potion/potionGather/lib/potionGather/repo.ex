defmodule PotionGather.Repo do
  use Ecto.Repo,
    otp_app: :potionGather,
    adapter: Ecto.Adapters.Postgres
end
