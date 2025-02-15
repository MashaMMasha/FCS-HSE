namespace mini_hw1;

public class Rabbit : IHerbo
{
    public int Food { get; set; }
    public int Number { get; set; }
    public string Name { get; }
    public int Age {get; set; }
    public int MaxAge { get; }
    public int Kindness { get;  }
    public int Carrot {get; set; }

    public void GiveFood()
    {
        Carrot++;
    }

    public Rabbit(string name, int food, int age, int kindness, int number, int maxAge = 20, int carrot = 0)
    {
        Name = name;
        Food = food;
        Number = number;
        Age = age;
        MaxAge = maxAge;
        Kindness = kindness;
        Carrot = carrot; 
    }
    public bool CanBeInContactZoo()
    {
        return Kindness > 5;
    }
    public bool IsHealthy()
    {
        return Age < MaxAge * 0.9 || Carrot > 0;
    }
}
