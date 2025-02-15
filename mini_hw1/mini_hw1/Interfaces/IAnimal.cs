namespace mini_hw1;

public interface IAnimal : IAlive, IInventory
{
    string Name { get; }
    int Food { get; set; }
    int Number { get; }
    int Age {get; set; }
    int MaxAge { get; }

    void GiveFood();

    bool CanBeInContactZoo();
    bool IsHealthy() {
        return Age < MaxAge;
    }
}