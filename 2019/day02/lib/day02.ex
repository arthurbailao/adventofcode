defmodule Day02 do
  @moduledoc """
  Documentation for `Day02`.
  """

  @doc """
  Parses the Intcode program into array of integers.

  ## Examples

      iex> Day02.parse_input("1,2,3,4")
      [1,2,3,4]

  """
  def parse_input(input) do
    input
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&elem(Integer.parse(&1), 0))
  end

  @doc """
  Exexutes an Intcode program.

  ## Examples

      iex> Day02.execute([1,1,1,3,1,1,2,0,99])
      [2,1,1,2,1,1,2,0,99]

  """
  def execute(program) do
    execute(0, program)
  end

  defp execute(position, program) do
    instruction = Enum.slice(program, position, 4)

    case execute_instruction(instruction, program) do
      {:step, new_program} ->
        execute(position + 4, new_program)

      {:done, new_program} ->
        new_program
    end
  end

  defp execute_instruction([1, x, y, r], program) do
    result = Enum.at(program, x) + safe_int(Enum.at(program, y))
    {:step, List.replace_at(program, r, result)}
  end

  defp execute_instruction([2, x, y, r], program) do
    result = Enum.at(program, x) * safe_int(Enum.at(program, y))
    {:step, List.replace_at(program, r, result)}
  end

  defp execute_instruction([99 | _tail], program) do
    {:done, program}
  end

  defp safe_int(nil), do: 0
  defp safe_int(i), do: i

  @doc """
  Exexutes an Intcode program until reaches the target value.
  It returns the used noun and verb.

  ## Examples
      iex> Day02.execute_until_value([1,0,0,3,1,1,2,0,99], 20)
      {1, 19}
  """
  def execute_until_value(program, target) do
    Enum.reduce_while(1..99, nil, fn i, _acc ->
      Enum.reduce_while(1..99, nil, fn j, _acc ->
        modified = program |> List.replace_at(1, i) |> List.replace_at(2, j)
        [value | _] = execute(modified)
        if value == target, do: {:halt, {:halt, {i, j}}}, else: {:cont, {:cont, {i, j}}}
      end)
    end)
  end
end
