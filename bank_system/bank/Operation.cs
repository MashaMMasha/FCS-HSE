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
}