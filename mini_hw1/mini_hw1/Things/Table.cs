namespace mini_hw1;

public class Table : IThing
{
    public int Number { get; }
    public string Name { get; }
    public string Material { get; }

    public Table(int number, string name, string material = "Wood")
    {
        Number = number;
        Name = name;
        Material = material;
    }
}