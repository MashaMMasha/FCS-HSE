namespace mini_hw1;

public class InventoryManager
{
    private int _currentId = 0;

    public int GetNextId()
    {
        return ++_currentId;
    }
}