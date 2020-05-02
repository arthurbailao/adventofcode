defmodule Day01 do
  @moduledoc """
  Documentation for `Day01`.
  """

  @doc """
  Parses Day 1 input.

  ## Examples

      iex> Day01.parse_input(\"""
      ...> 12
      ...> 14
      ...> 1969
      ...> 100756
      ...> \""")
      [12, 14, 1969, 100756]

  """
  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&Integer.parse(&1))
    |> Enum.map(&elem(&1, 0))
  end

  @doc """
  Calculates the required fuel to a given mass.

  ## Examples

      iex> Day01.calculate_fuel(100756)
      33583
  """
  def calculate_fuel(mass) do
    div(mass, 3) - 2
  end

  @doc """
  Calculates the required total of fuel for all modules.

  ## Examples

      iex> Day01.calculate_total_fuel(\"""
      ...> 12
      ...> 14
      ...> 1969
      ...> 100756
      ...> \""")
      34241
  """
  def calculate_total_fuel(input) do
    input
    |> parse_input()
    |> Enum.map(&calculate_fuel/1)
    |> Enum.sum()
  end

  @doc """
  Calculates the required fuel recursively.

  ## Examples

      iex> Day01.calculate_fuel_of_fuel(100756)
      50346
  """
  def calculate_fuel_of_fuel(fuel) do
    calculate_fuel_of_fuel(fuel, [])
  end

  defp calculate_fuel_of_fuel(fuel, state) when fuel <= 0 do
    Enum.sum(state)
  end

  defp calculate_fuel_of_fuel(fuel, state) do
    new_fuel = calculate_fuel(fuel)
    calculate_fuel_of_fuel(new_fuel, [max(new_fuel, 0) | state])
  end

  @doc """
  Calculates the total required fuel recusively.

  ## Examples

      iex> Day01.calculate_total_fuel_of_fuel(\"""
      ...> 12
      ...> 1969
      ...> 100756
      ...> \""")
      51314
  """
  def calculate_total_fuel_of_fuel(input) do
    input
    |> parse_input()
    |> Enum.map(&calculate_fuel_of_fuel/1)
    |> Enum.sum()
  end
end
