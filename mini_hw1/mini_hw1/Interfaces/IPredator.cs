namespace mini_hw1;

public interface IPredator : IAnimal
{
    string Name { get; }
    int Food { get; set; }
    int Number { get; }
    int Age {get; set; }
    int MaxAge { get; }
    int Danger { get;  }
    bool CanBeInContactZoo();
    bool IsHealty()
    {
        return Age <= MaxAge * 0.8;
    }
}