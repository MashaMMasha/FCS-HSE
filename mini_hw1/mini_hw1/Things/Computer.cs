namespace mini_hw1;

public class Computer : IThing
{
    public int Number { get; }
    public string Name { get; }
    public string Brand { get; }

    public Computer(int number, string name, string brand = "Apple")
    {
        Number = number;
        Name = name;
        Brand = brand;
    }
}