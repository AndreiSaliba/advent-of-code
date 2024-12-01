namespace AOC2024;

public class Day01 : BaseDay
{
    private readonly (List<int> Left, List<int> Right) input = ParseInputs();

    public override ValueTask<string> Solve_1()
    {
        return new(input.Left.Zip(input.Right, (l, r) => Math.Abs(l - r)).Sum().ToString());
    }

    public override ValueTask<string> Solve_2()
    {
        return new(input.Left.Select(i => i * input.Right.FindAll(j => j == i).Count).Sum().ToString());
    }

    private static (List<int>, List<int>) ParseInputs()
    {
        List<int> left = [];
        List<int> right = [];
        
        foreach (var line in File.ReadAllLines("Inputs/Day01.txt"))
        {
            var parts = line.Split("   ");
            left.Add(int.Parse(parts[0]));
            right.Add(int.Parse(parts[1]));
        }

        left.Sort();
        right.Sort();

        return (left, right);
    }
}