namespace BankSystem;

public class Operation
{
    public readonly Guid id;
    public OperationType type;
    public readonly Guid accountId;
    public Guid categoryId;
    public readonly int amount;
    public readonly DateTime date;
    public string description;
    public Operation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date, string description)
    {
        id = Guid.NewGuid();
        this.accountId = accountId;
        this.categoryId = categoryId;
        this.type = type;
        this.amount = amount;
        this.date = date;
        this.description = description;
    }
    public Operation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date)
    {
        id = Guid.NewGuid();
        this.accountId = accountId;
        this.categoryId = categoryId;
        this.type = type;
        this.amount = amount;
        this.date = date;
        description = "";
    }
}