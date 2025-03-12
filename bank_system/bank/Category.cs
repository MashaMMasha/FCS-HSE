namespace BankSystem;

enum OperationType
{
    Income,
    Withdrawal
}

public class Category
{
    public readonly Guid id;
    public string name;
    public readonly OperationType type;
    Category(string name, OperationType type)
    {
        id = Guid.NewGuid();
        this.name = name;
        this.type = type;
    }
}