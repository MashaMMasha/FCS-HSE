namespace mini_hw1;

public class ThingFactory
{
    private readonly InventoryManager _inventoryManager;

    public ThingFactory(InventoryManager inventoryManager)
    {
        _inventoryManager = inventoryManager;
    }

    public Table CreateTable(string name, string? material = null)
    {
        if (material == null)
        {
            return new Table(_inventoryManager.GetNextId(), name);
        }
        return new Table( _inventoryManager.GetNextId(), name, material);
    }

    public Computer CreateComputer(string name, string? barnd = null)
    {
        if (barnd == null)
        {
            return new Computer(_inventoryManager.GetNextId(), name);
        }
        return new Computer(_inventoryManager.GetNextId(), name, barnd);
    }

}