namespace mini_hw1;

public class Wolf : IPredator
{
    public string Name { get; }
    public int Food { get; set; }
    public int Number { get; }
    public int Age {get; set; }
    public int MaxAge { get; }
    public int Danger { get;  }
    public int Meat { get; set; }

    public void GiveFood()
    {
        Meat++;
    }
    public bool CanBeInContactZoo()
    {
        return Danger < 2 && Meat > 1;
    }
    public bool IsHealthy()
    {
        return Age <= MaxAge * 0.8 || Meat > 0;
    }
    public Wolf(string name, int food, int age, int danger, int number, int maxAge = 50, int meat = 0)
    {
        Name = name;
        Food = food;
        Age = age;
        Danger = danger;
        Meat = meat;
        MaxAge = maxAge;
        Number = number;
    }
}