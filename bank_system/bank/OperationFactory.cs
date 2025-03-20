namespace BankSystem;

public class OperationFactory
{
    private Dictionary<Guid, Operation> operations = new();

    public Guid AddOperation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date,
        string description)
    {
        var operation = new Operation(accountId, categoryId, type, amount, date, description);
        operations.Add(operation.id, operation);
        return operation.id;
    }
    public List<Operation> GetOperations()
    {
        return new List<Operation>(operations.Values);
    }

    public void AddOperation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date)
    {
        var operation = new Operation(accountId, categoryId, type, amount, date, "");
        operations.Add(operation.id, operation);
    }

    public void ChangeDescription(Guid id, string newDescription)
    {
        if (operations.ContainsKey(id))
        {
            operations[id].description = newDescription;
        }
    }

    public bool ChangeCategory(Guid id, Guid newCategoryId)
    {
        if (operations.ContainsKey(id))
        {
            operations[id].categoryId = newCategoryId;
            return true;
        }

        return false;
    }

    public bool DeleteOperation(Guid id)
    {
        if (!operations.ContainsKey(id))
        {
            return false;
        }

        operations.Remove(id);
        return true;
    }

    public IEnumerable<Operation> GetOperations(Guid accountId)
    {
        return operations.Values.Where(operation => operation.accountId == accountId);
    }

    public IEnumerable<Operation> GetOperations(Guid categoryId, Guid accountId)
    {
        return operations.Values.Where(operation =>
            operation.categoryId == categoryId && operation.accountId == accountId);
    }

    public IEnumerable<Operation> GetOperations(DateTime start, DateTime end)
    {
        return operations.Values.Where(operation => operation.date >= start && operation.date <= end);
    }

    public IEnumerable<Operation> GetOperations(DateTime start, DateTime end, Guid categoryId)
    {
        return operations.Values.Where(operation =>
            operation.date >= start && operation.date <= end && operation.categoryId == categoryId);
    }
}