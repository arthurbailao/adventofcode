defmodule Day03 do
  @moduledoc """
  Documentation for `Day03`.
  """

  @doc """
  Parses wires paths.

  ## Examples

      iex> Day03.parse_input(\"""
      ...> R8,U5
      ...> L5,D33
      ...> \""")
      [[{:right, 8}, {:up, 5}],[{:left, 5}, {:down, 33}]]

  """
  def parse_input(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&String.split(&1, ","))
    |> Enum.map(fn paths ->
      Enum.map(paths, &parse_direction/1)
    end)
  end

  defp parse_direction(<<"R", rest::binary>>), do: {:right, elem(Integer.parse(rest), 0)}
  defp parse_direction(<<"D", rest::binary>>), do: {:down, elem(Integer.parse(rest), 0)}
  defp parse_direction(<<"L", rest::binary>>), do: {:left, elem(Integer.parse(rest), 0)}
  defp parse_direction(<<"U", rest::binary>>), do: {:up, elem(Integer.parse(rest), 0)}

  @doc """
  Extracts points from a defined path starting at {0,0}.

  ## Examples

      iex> Day03.extract_points([{:right, 2}, {:up, 4}])
      [{1, 0}, {2, 0}, {2, -1}, {2, -2}, {2, -3}, {2, -4}]
  """
  def extract_points(path) do
    Enum.flat_map_reduce(path, {0, 0}, fn direction, acc ->
      points = move(acc, direction)
      {points, List.last(points)}
    end)
    |> elem(0)
  end

  defp move({x, y}, {:right, n}), do: Enum.map((x + 1)..(x + n), &{&1, y})
  defp move({x, y}, {:down, n}), do: Enum.map((y + 1)..(y + n), &{x, &1})
  defp move({x, y}, {:left, n}), do: Enum.map((x - n)..(x - 1), &{&1, y}) |> Enum.reverse()
  defp move({x, y}, {:up, n}), do: Enum.map((y - n)..(y - 1), &{x, &1}) |> Enum.reverse()

  @doc """
  Finds the Manhattan distance from the central port to the closest intersection

  ## Examples

      iex> Day03.parse_input(\"""
      ...> R75,D30,R83,U83,L12,D49,R71,U7,L72
      ...> U62,R66,U55,R34,D71,R55,D58,R83
      ...> \""") |> Day03.find_distance()
      159

  """
  def find_distance(paths) do
    [a, b] = Enum.map(paths, &extract_points/1)
    intersections = a -- a -- b

    intersections
    |> Enum.map(&distance({0, 0}, &1))
    |> Enum.min()
  end

  defp distance({x1, y1}, {x2, y2}), do: abs(x1 - x2) + abs(y1 - y2)

  @doc """
  Finds the fewest combined steps the wires must take to reach an intersection.

  ## Examples

      iex> Day03.parse_input(\"""
      ...> R75,D30,R83,U83,L12,D49,R71,U7,L72
      ...> U62,R66,U55,R34,D71,R55,D58,R83
      ...> \""") |> Day03.find_steps()
      610

  """
  def find_steps(paths) do
    [a, b] = result = Enum.map(paths, &extract_points/1)
    intersections = a -- a -- b

    steps =
      Enum.map(result, fn path ->
        Enum.map(intersections, fn point ->
          Enum.find_index(path, &(point == &1)) + 1
        end)
      end)

    Enum.zip(Enum.at(steps, 0), Enum.at(steps, 1))
    |> Enum.map(fn {a, b} -> a + b end)
    |> Enum.min()
  end
end
