namespace mini_hw1;

public class AnimalFactory
{
    private readonly InventoryManager _inventoryManager;

    public AnimalFactory(InventoryManager inventoryManager)
    {
        _inventoryManager = inventoryManager;
    }

    public Monkey CreateMonkey(string name, int food, int age, int kindness)
    {
        return new Monkey(name, food, age, kindness, _inventoryManager.GetNextId());
    }

    public Rabbit CreateRabbit(string name, int food, int age, int kindness)
    {
        return new Rabbit(name, food, age, kindness, _inventoryManager.GetNextId());
    }
    
    public Tiger CreateTiger(string name, int food, int age, int danger)
    {
        return new Tiger(name, food, age, danger,_inventoryManager.GetNextId());
    }
    public Wolf CreateWolf(string name, int food, int age, int danger)
    {
        return new Wolf(name, food, age, danger, _inventoryManager.GetNextId());
    }
    
}