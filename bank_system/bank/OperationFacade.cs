namespace BankSystem;

public class OperationFacade
{
    private Dictionary<Guid, Operation> operations = new();
    public void AddOperation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date, string description)
    {
        var operation = new Operation(accountId, categoryId, type, amount, date, description);
        operations.Add(operation.id, operation);
    }
    public void AddOperation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date)
    {
        var operation = new Operation(accountId, categoryId, type, amount, date, "");
        operations.Add(operation.id, operation);
    }
    public void ChangeDescription(Guid id, string newDescription)
    {
        operations[id].description = newDescription;
    }
    public void ChangeCategory(Guid id, Guid newCategoryId)
    {
        operations[id].categoryId = newCategoryId;
    }
    public void DeleteOperation(Guid id)
    {
        operations.Remove(id);
    }
    public void GetOperation(Guid accountId)
    {
        return operations.Values.Where(operation => operation.accountId == accountId);
    }
    public void GetOperation(Guid categoryId)
    {
        return operations.Values.Where(operation.categoryId == categoryId);
    }
    
    public void GetOperation(DateTime start, DateTime end)
    {
        return operations.Values.Where(operation => operation.date >= start && operation.date <= end);
    }
    public void GetOperation(DateTime start, DateTime end, Guid categoryId)
    {
        return operations.Values.Where(operation => operation.date >= start && operation.date <= end && operation.categoryId == categoryId);
    }   
    
}
