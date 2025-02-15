namespace mini_hw1;

public class Monkey : IHerbo
{
    public int Food { get; set; }
    public int Number { get; set; }
    public string Name { get; }
    public int Age {get; set; }
    public int MaxAge { get; }
    public int Kindness { get;  }
    public int Bananas {get; set; }

    public void GiveFood()
    {
        Bananas++;
    }
    public bool CanBeInContactZoo()
    {
        return Kindness > 8;
    }
    public Monkey(string name, int food, int age, int kindness, int number, int maxAge = 20, int bananas = 0)
    {
        Name = name;
        Food = food;
        Number = number;
        Age = age;
        MaxAge = maxAge;
        Kindness = kindness;
        Bananas = bananas; 
    }

    public bool IsHealthy()
    {
        return Age < MaxAge * 0.9 || Bananas > 0;
    }
    
}