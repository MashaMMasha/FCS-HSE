namespace BankSystem;

public class Category
{
    public readonly Guid id;
    public string name;
    public readonly OperationType type;
    public Category(string name, OperationType type)
    {
        id = Guid.NewGuid();
        this.name = name;
        this.type = type;
    }
}