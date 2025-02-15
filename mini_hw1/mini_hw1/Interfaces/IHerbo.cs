namespace mini_hw1;

public interface IHerbo : IAnimal 
{
    string Name { get; }
    int Food { get; set; }
    int Number { get; }
    int Age {get; set; }
    int MaxAge { get; }
    int Kindness { get; }

    bool IsHealty()
    {
        return Age <= MaxAge * 0.9;
    }

    bool CanBeInContactZoo();
}